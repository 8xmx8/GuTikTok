package prom

import (
	"GuTikTok/config"
	"context"
	"github.com/prometheus/client_golang/prometheus"
	"go.opentelemetry.io/otel/trace"
)

func ExtractContext(ctx context.Context) prometheus.Labels {
	if span := trace.SpanContextFromContext(ctx); span.IsSampled() {
		return prometheus.Labels{
			"traceID": span.TraceID().String(),
			"spanID":  span.SpanID().String(),
			"podId":   config.Conf.Server.Address,
		}
	}
	return nil
}
