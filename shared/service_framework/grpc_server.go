package service_framework

import (
	"context"
	"fmt"
	"net"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// StartGRPCServer starts a grpc server in the background
func (m *Manager) StartGRPCServer(grpcServer *grpc.Server, port int) {
	go func() {
		listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
		if err != nil {
			m.Logger().Errorf("Failed to open listener on port %d: %v", port, err)
			m.cancelFunc() // Shutdown the service if this fails
		}

		reflection.Register(grpcServer)

		m.Logger().Infof("serving on port %d", port)
		if err := grpcServer.Serve(listener); err != nil {
			m.Logger().Errorf("error serving on port %d: %v", port, err)
			m.cancelFunc() // Shutdown the service if this fails
		}
	}()
}

// NewGRPCServer creates a new GRPC server with the coherent addons
func (m *Manager) NewGRPCServer() *grpc.Server {
	opts := []grpc.ServerOption{grpc.UnaryInterceptor(m.metricsInterceptor)}
	grpcServer := grpc.NewServer(opts...)
	reflection.Register(grpcServer)
	return grpcServer
}

func (m *Manager) metricsInterceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (interface{}, error) {
	start := time.Now()

	// Calls the handler
	ret, err := handler(ctx, req)

	m.Metrics().Gauge(fmt.Sprintf("grpc.%s.time", info.FullMethod), float64(time.Since(start)), []string{}, 1.0)

	return ret, err
}
