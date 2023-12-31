package application

import (
	"fmt"

	"github.com/gustavo-m-franco/qd-common/pkg/grpcserver"
	"github.com/gustavo-m-franco/qd-common/pkg/log"

	"qd-authentication-api/internal/config"
	grpcFactory "qd-authentication-api/internal/grpcserver"
	"qd-authentication-api/internal/service"
)

// Applicationer provides the main functions to start the application
type Applicationer interface {
	StartServer()
	Close()
	GetGRPCServerAddress() string
}

// Application is the main application
type Application struct {
	grpcServiceServer grpcserver.GRPCServicer
	grpcServerAddress string
	service           service.Servicer
	logger            log.Loggerer
}

// NewApplication creates a new application
func NewApplication(config *config.Config) Applicationer {
	logFactory := log.NewLogFactory(config.Environment)
	logger := logFactory.NewLogger()

	service, err := (&service.Factory{}).CreateService(config)
	if err != nil {
		logger.Error(err, "Failed to create authentication service")
	}

	grpcServerAddress := fmt.Sprintf("%s:%s", config.GRPC.Host, config.GRPC.Port)
	grpcServiceServer, err := (&grpcFactory.Factory{}).Create(
		grpcServerAddress,
		service.GetAuthenticationService(),
		logFactory,
	)

	if err != nil {
		logger.Error(err, "Failed to create grpc server")
	}

	return New(grpcServiceServer, grpcServerAddress, service, logger)
}

// New creates a new application with raw parameters
func New(grpcServiceServer grpcserver.GRPCServicer, grpcServerAddress string, service service.Servicer, logger log.Loggerer) Applicationer {
	return &Application{
		grpcServiceServer: grpcServiceServer,
		grpcServerAddress: grpcServerAddress,
		service:           service,
		logger:            logger,
	}
}

// StartServer starts the gRPC server
func (application *Application) StartServer() {
	application.logger.Info(fmt.Sprintf("Starting gRPC server on %s:...", application.grpcServerAddress))
	err := application.grpcServiceServer.Serve()
	if err != nil {
		application.logger.Error(err, "Failed to serve grpc server")
	}
}

// Close closes the gRPC server and services used by the application
func (application *Application) Close() {
	switch {
	case application.service == nil:
		application.logger.Error(nil, "Service is not created")
		return
	case application.grpcServiceServer == nil:
		application.logger.Error(nil, "gRPC server is not created")
		return
	}
	application.service.Close()
	application.grpcServiceServer.Close()
	application.logger.Info("Application closed")
}

// GetGRPCServerAddress returns the gRPC server address
func (application *Application) GetGRPCServerAddress() string {
	return application.grpcServerAddress
}
