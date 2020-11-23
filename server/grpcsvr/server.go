package grpcsvr

import (
	"context"
	"net"
	"os"
	"os/signal"

	v1 "github.com/jacobweinstock/pbnjtest/api/v1"
	"github.com/jacobweinstock/pbnjtest/pkg/logging"
	"github.com/jacobweinstock/pbnjtest/pkg/repository"
	"github.com/jacobweinstock/pbnjtest/server/grpcsvr/persistence"
	"github.com/jacobweinstock/pbnjtest/server/grpcsvr/rpc"
	"github.com/jacobweinstock/pbnjtest/server/grpcsvr/taskrunner"
	"github.com/philippgille/gokv"
	"github.com/philippgille/gokv/freecache"
	"google.golang.org/grpc"
)

// Server options
type Server struct {
	repository.Actions
}

// ServerOption for setting optional values
type ServerOption func(*Server)

// WithPersistence sets the log level
func WithPersistence(repo repository.Actions) ServerOption {
	return func(args *Server) { args.Actions = repo }
}

// RunServer registers all services and runs the server
func RunServer(ctx context.Context, log logging.Logger, grpcServer *grpc.Server, port string, opts ...ServerOption) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	defaultStore := gokv.Store(freecache.NewStore(freecache.DefaultOptions))

	// instantiate a Repository for task persistence
	repo := &persistence.GoKV{
		Store: defaultStore,
		Ctx:   ctx,
	}

	defaultServer := &Server{
		Actions: repo,
	}

	for _, opt := range opts {
		opt(defaultServer)
	}

	taskRunner := &taskrunner.Runner{
		Repository: defaultServer.Actions,
		Ctx:        ctx,
		Log:        log,
	}

	ms := rpc.MachineService{
		Log:        log,
		TaskRunner: taskRunner,
	}
	v1.RegisterMachineServer(grpcServer, &ms)

	bs := rpc.BmcService{
		Log:        log,
		TaskRunner: taskRunner,
	}
	v1.RegisterBMCServer(grpcServer, &bs)

	ts := rpc.TaskService{
		Log:        log,
		TaskRunner: taskRunner,
	}
	v1.RegisterTaskServer(grpcServer, &ts)

	listen, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return err
	}

	// graceful shutdowns
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)
	go func() {
		for range sigChan {
			log.V(0).Info("sig received, shutting down PBnJ")
			grpcServer.GracefulStop()
			<-ctx.Done()
		}
	}()

	go func() {
		<-ctx.Done()
		log.V(0).Info("ctx cancelled, shutting down PBnJ")
		grpcServer.GracefulStop()
	}()

	log.V(0).Info("starting PBnJ gRPC server")
	return grpcServer.Serve(listen)
}
