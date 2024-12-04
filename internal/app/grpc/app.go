package grpcapp

import (
	"fmt"
	"github.com/op/go-logging"
	"google.golang.org/grpc"
	"net"
	authgRPC "sso/internal/grpc/auth"
)

var log = logging.MustGetLogger("app")

type App struct {
	gRPCServer *grpc.Server
	port       int
}

func New(port int) *App {
	gRPC := grpc.NewServer()

	authgRPC.Register(gRPC)

	return &App{
		gRPCServer: gRPC,
		port:       port,
	}
}

func (a *App) MustRun() {
	if err := a.Run(); err != nil {
		panic(err)
	}
}

func (a *App) Run() error {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", a.port))
	if err != nil {
		return fmt.Errorf("Error starting gRPC server: %s", err)
	}

	log.Infof("grpc server is running address: %s", listener.Addr())
	fmt.Println("grpc server is running address: %s", listener.Addr())
	if err := a.gRPCServer.Serve(listener); err != nil {
		return fmt.Errorf("Error starting gRPC server: %s", err)
	}

	return nil
}

func (a *App) Stop() {
	log.Infof("stopping grpc server")
	a.gRPCServer.GracefulStop()
}
