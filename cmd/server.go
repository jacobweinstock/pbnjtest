package cmd

import (
	"context"
	"fmt"
	"os"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	grpc_validator "github.com/grpc-ecosystem/go-grpc-middleware/validator"
	"github.com/jacobweinstock/pbnjtest/cmd/zaplog"
	"github.com/jacobweinstock/pbnjtest/server/grpcsvr"
	"github.com/packethost/pkg/log/logr"
	"github.com/spf13/cobra"
	"goa.design/goa/grpc/middleware"
	"google.golang.org/grpc"
)

const (
	requestIDKey    = "x-request-id"
	requestIDLogKey = "requestID"
)

var (
	port string
	// serverCmd represents the server command
	serverCmd = &cobra.Command{
		Use:   "server",
		Short: "Run PBnJ server",
		Long:  `Run PBnJ server for interacting with BMCs.`,
		Run: func(cmd *cobra.Command, args []string) {
			ctx := context.Background()
			ctx, cancel := context.WithCancel(ctx)
			defer cancel()

			logger, zlog, err := logr.NewPacketLogr(
				logr.WithServiceName("github.com/jacobweinstock/pbnjtest"),
				logr.WithLogLevel(logLevel),
			)
			if err != nil {
				fmt.Fprintf(os.Stderr, "%v\n", err)
				os.Exit(1)
			}
			defer zlog.Sync() // nolint

			// Make sure that log statements internal to gRPC library are logged using the zapLogger as well.
			grpc_zap.ReplaceGrpcLoggerV2(zlog)

			grpcServer := grpc.NewServer(
				grpc_middleware.WithUnaryServerChain(
					grpc_ctxtags.UnaryServerInterceptor(grpc_ctxtags.WithFieldExtractor(grpc_ctxtags.CodeGenRequestFieldExtractor)),
					middleware.UnaryRequestID(middleware.UseXRequestIDMetadataOption(true), middleware.XRequestMetadataLimitOption(512)),
					zaplog.UnaryLogRequestID(zlog, requestIDKey, requestIDLogKey),
					grpc_zap.UnaryServerInterceptor(zlog),
					grpc_validator.UnaryServerInterceptor(),
				),
			)

			if err := grpcsvr.RunServer(ctx, zaplog.RegisterLogger(logger), grpcServer, port); err != nil {
				logger.Error(err, "error running server")
				os.Exit(1)
			}
		},
	}
)

func init() {
	serverCmd.PersistentFlags().StringVar(&port, "port", "9090", "server port (default is 9090")
	rootCmd.AddCommand(serverCmd)
}
