package app

import (
	grpcapp "sso/internal/app/grpc"
	"time"
)

type App struct {
	Server *grpcapp.App
}

func New(gRPCport int, storagePath string, tokenTTL time.Duration) *App {
	gRPCapp := grpcapp.New(gRPCport)

	return &App{
		Server: gRPCapp,
	}
}
