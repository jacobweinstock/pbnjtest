package taskrunner

import (
	"context"
	"testing"
	"time"

	"github.com/grpc-ecosystem/go-grpc-middleware/logging/zap/ctxzap"
	"github.com/jacobweinstock/pbnjtest/cmd/zaplog"
	"github.com/jacobweinstock/pbnjtest/pkg/repository"
	"github.com/jacobweinstock/pbnjtest/server/grpcsvr/persistence"
	"github.com/packethost/pkg/log/logr"
	"github.com/philippgille/gokv"
	"github.com/philippgille/gokv/freecache"
)

func TestRoundTrip(t *testing.T) {
	description := "test task"
	defaultError := &repository.Error{
		Code:    0,
		Message: "",
		Details: nil,
	}
	ctx := context.Background()
	f := freecache.NewStore(freecache.DefaultOptions)
	s := gokv.Store(f)
	defer s.Close()
	repo := &persistence.GoKV{Store: s, Ctx: ctx}
	l, zapLogger, _ := logr.NewPacketLogr()
	logger := zaplog.RegisterLogger(l)
	ctx = ctxzap.ToContext(ctx, zapLogger)
	runner := Runner{
		Repository: repo,
		Ctx:        ctx,
		Log:        logger,
	}

	id, err := runner.Execute(ctx, description, func(s chan string) (string, error) {
		return "didnt do anything", defaultError //nolint
	})
	if err != nil {
		t.Fatal(err)
	}
	if len(id) != 20 {
		t.Fatalf("expected id of length 20,  got: %v (%v)", len(id), id)
	}

	// must be min of 3 because we sleep 2 seconds in worker function to allow final status messages to be written
	time.Sleep(500 * time.Millisecond)
	record, err := runner.Status(ctx, id)
	if err != nil {
		t.Fatal(err)
	}
	if record.Complete != true {
		t.Fatalf("expected task to be complete, got: %+v", record)
	}
}
