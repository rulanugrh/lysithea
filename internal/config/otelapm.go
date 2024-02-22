package config

import (
	"log"

	apmotel "go.elastic.co/apm/module/apmotel/v2"
	"go.elastic.co/apm/v2"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/trace"
)

func NewTracerOtel() *trace.TracerProvider {

	provider, err := apmotel.NewTracerProvider()
	if err != nil {
		log.Fatal(err)
	}

	otel.SetTracerProvider(provider)

	return &provider
}

func NewMeterOtel() *metric.MeterProvider {
	provider, err := apmotel.NewGatherer()
	if err != nil {
		log.Fatal(err)
	}

	metric := metric.NewMeterProvider(metric.WithReader(provider))
	otel.SetMeterProvider(metric)

	apm.DefaultTracer().RegisterMetricsGatherer(provider)
	return metric
}

func GetTracer(tr trace.TracerProvider) trace.Tracer {
	trace := tr.Tracer("lysithea-app")
	return trace
}
