package metric

import (
	"context"

	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc"
	"go.opentelemetry.io/otel/exporters/stdout/stdoutmetric"
	"go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/resource"
	"google.golang.org/grpc"
)

func NewExporterConsole() (metric.Exporter, error) {
	return stdoutmetric.New()
}

func NewExporterOTLP(ctx context.Context, conn *grpc.ClientConn) (*otlpmetricgrpc.Exporter, error) {
	return otlpmetricgrpc.New(ctx, otlpmetricgrpc.WithGRPCConn(conn))
}

func NewProvider(res *resource.Resource, exp metric.Exporter) *metric.MeterProvider {
	return metric.NewMeterProvider(
		metric.WithResource(res),
		metric.WithReader(metric.NewPeriodicReader(exp)),
	)
}
