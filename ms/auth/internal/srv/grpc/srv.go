// Package grpc implements gRPC method handlers.
package grpc

import (
	"context"
	"crypto/tls"

	"google.golang.org/grpc"

	api "github.com/powerman/go-monolith-example/api/proto/powerman/example/auth"
	"github.com/powerman/go-monolith-example/ms/auth/internal/app"
	"github.com/powerman/go-monolith-example/pkg/grpcx"
)

// Ctx is a synonym for convenience.
type Ctx = context.Context

type Config struct {
	Cert *tls.Certificate
}

type server struct {
	api.UnimplementedNoAuthSvcServer
	api.UnimplementedAuthSvcServer
	api.UnimplementedAuthIntSvcServer
	appl app.Appl
}

// NewServer creates and returns gRPC server.
func NewServer(appl app.Appl, cfg Config) *grpc.Server {
	srv := &server{
		appl: appl,
	}
	server := grpcx.NewServer(app.ServiceName, app.Metric, metric.server, cfg.Cert, srv.authn)
	api.RegisterNoAuthSvcServer(server, srv)
	api.RegisterAuthSvcServer(server, srv)
	metric.server.InitializeMetrics(server)
	return server
}

// NewServerInt creates and returns gRPC server.
func NewServerInt(appl app.Appl, cfg Config) *grpc.Server {
	srv := &server{
		appl: appl,
	}
	server := grpcx.NewServer(app.ServiceName, app.Metric, metric.server, cfg.Cert, srv.authn)
	api.RegisterAuthIntSvcServer(server, srv)
	metric.server.InitializeMetrics(server)
	return server
}