package log

import (
	"context"

	"go.opentelemetry.io/otel/exporters/otlp/otlplog/otlploggrpc"
	"go.opentelemetry.io/otel/sdk/log"
	"go.opentelemetry.io/otel/sdk/resource"
	"google.golang.org/grpc"
)

func NewExporterOTLP(ctx context.Context, conn *grpc.ClientConn) (*otlploggrpc.Exporter, error) {
	return otlploggrpc.New(ctx, otlploggrpc.WithGRPCConn(conn))
}

func NewProvider(res *resource.Resource, exp log.Exporter) *log.LoggerProvider {
	return log.NewLoggerProvider(
		log.WithResource(res),
		log.WithProcessor(log.NewBatchProcessor(exp)),
	)
}
