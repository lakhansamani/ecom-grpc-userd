package service

import (
	"go.opentelemetry.io/otel/sdk/trace"
	otrace "go.opentelemetry.io/otel/trace"

	user "github.com/lakhansamani/ecom-grpc-apis/user/v1"
	"github.com/prometheus/client_golang/prometheus"

	"github.com/lakhansamani/ecom-grpc-userd/db"
)

type Config struct {
	// Add configuration here
	JWTSecret string
}

type Dependencies struct {
	// Add dependencies here
	DBProvider    db.Provider
	TraceProvider *trace.TracerProvider
}

// Service implements the User service.
type Service interface {
	user.UserServiceServer
}

type service struct {
	Config
	Dependencies
	trace otrace.Tracer
}

// New creates a new User service.
func New(cfg Config, deps Dependencies) Service {
	trace := deps.TraceProvider.Tracer("service")
	prometheus.MustRegister(loginMetrics)
	return &service{
		Config:       cfg,
		Dependencies: deps,
		trace:        trace,
	}
}
