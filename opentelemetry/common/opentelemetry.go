package common

import (
	"context"
	"time"

	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/exporters/stdout/stdoutmetric"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
)

func NewTraceExporterOTLP(ctx context.Context, endpoint string) (*otlptrace.Exporter, error) {
	insecureOpt := otlptracegrpc.WithInsecure()
	endpointOpt := otlptracegrpc.WithEndpoint(endpoint)
	return otlptracegrpc.New(ctx, insecureOpt, endpointOpt)
}

func NewTraceExporterConsole() (trace.SpanExporter, error) {
	return stdouttrace.New()
}

func NewMetricExporterConsole() (metric.Exporter, error) {
	return stdoutmetric.New()
}

func NewTracerProvider(exp trace.SpanExporter, res *resource.Resource) *trace.TracerProvider {
	return trace.NewTracerProvider(
		trace.WithBatcher(exp),
		trace.WithResource(res),
	)
}

func NewMeterProvider(exp metric.Exporter, res *resource.Resource) *metric.MeterProvider {
	return metric.NewMeterProvider(
		metric.WithResource(res),
		metric.WithReader(metric.NewPeriodicReader(exp, metric.WithInterval(3*time.Second))),
	)
}
