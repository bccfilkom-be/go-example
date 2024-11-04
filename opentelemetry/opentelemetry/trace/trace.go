package trace

import (
	"context"

	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	"google.golang.org/grpc"
)

func NewExporterConsole() (trace.SpanExporter, error) {
	return stdouttrace.New()
}

func NewExporterOTLP(ctx context.Context, conn *grpc.ClientConn) (*otlptrace.Exporter, error) {
	return otlptracegrpc.New(ctx, otlptracegrpc.WithGRPCConn(conn))
}

func NewProvider(res *resource.Resource, exp trace.SpanExporter) *trace.TracerProvider {
	return trace.NewTracerProvider(trace.WithResource(res), trace.WithBatcher(exp))
}
