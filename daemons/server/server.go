package server

import (
	"net"

	gometrics "github.com/hashicorp/go-metrics"
	"github.com/tellor-io/layer/daemons/constants"
	"github.com/tellor-io/layer/daemons/server/types"
	daemontypes "github.com/tellor-io/layer/daemons/types"
	"github.com/tellor-io/layer/lib/metrics"

	// "syscall"
	"cosmossdk.io/log"

	"github.com/cosmos/cosmos-sdk/telemetry"
)

// Server struct defines the shared gRPC server for all daemons.
// The struct contains fields related to gRPC server that are common to all daemons.
// In addition, the struct contains fields that are specific to various daemon services.
// needed for various services.
// to check queryServer implements ServiceServer

type Server struct {
	logger        log.Logger
	gsrv          daemontypes.GrpcServer
	fileHandler   daemontypes.FileHandler
	socketAddress string
	PriceFeedServer
	TokenBridgeFeedServer
}

// NewServer creates a single gRPC server that's shared across multiple daemons for communication.
func NewServer(
	logger log.Logger,
	grpcServer daemontypes.GrpcServer,
	fileHandler daemontypes.FileHandler,
	socketAddress string,
) *Server {
	return &Server{
		logger:        logger,
		gsrv:          grpcServer,
		fileHandler:   fileHandler,
		socketAddress: socketAddress,
	}
}

// Stop stops the daemon server's gRPC service.
func (server *Server) Stop() {
	server.gsrv.Stop()
}

// reportValidResponse reports a valid request/response from a daemon service for metrics collection purposes.
func (server *Server) reportValidResponse(daemonKey string) {
	telemetry.IncrCounterWithLabels(
		[]string{
			metrics.DaemonServer,
			metrics.ValidResponse,
		},
		1,
		[]gometrics.Label{
			metrics.GetLabelForStringValue(metrics.Daemon, daemonKey),
		},
	)
}

// Start clears the current socket and establishes a new socket connection
// on the local filesystem.
// See URL for more information: https://eli.thegreenplace.net/2019/unix-domain-sockets-in-go/
func (server *Server) Start() {
	if err := server.fileHandler.RemoveAll(server.socketAddress); err != nil {
		server.logger.Error("Failed to clear socket for daemon gRPC server", "error", err)
		panic(err)
	}

	// Restrict so that only user can read or write to socket generated by `net.Listen`.
	// oldValue := syscall.Umask(constants.UmaskUserReadWriteOnly)

	ln, err := net.Listen(constants.UnixProtocol, server.socketAddress)
	// Restore umask bits back to previous value so that the entire process is not restricted to `UmaskUserReadWriteOnly`.
	// syscall.Umask(oldValue)
	if err != nil {
		server.logger.Error("Failed to listen to daemon gRPC server", "error", err)
		panic(err)
	}

	server.logger.Info("Daemon gRPC server is listening", "address", ln.Addr())

	// Register gRPC services needed by the daemons. This is required before invoking `Serve`.
	// https://pkg.go.dev/google.golang.org/grpc#Server.RegisterService

	// Register Server to ingest gRPC requests from price feed daemon and update market prices.
	types.RegisterPriceFeedServiceServer(server.gsrv, server)
	types.RegisterTokenBridgeServiceServer(server.gsrv, server)

	if err := server.gsrv.Serve(ln); err != nil {
		server.logger.Error("daemon gRPC server stopped with an error", "error", err)
		panic(err)
	}
}
