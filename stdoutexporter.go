package stdoutexporter

import (
	"context"
	"fmt"
	"os"

	"go.opentelemetry.io/collector/exporter"
	"go.opentelemetry.io/collector/pdata/plog"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/pdata/ptrace"
)

// Config defines configuration for the stdoutexporter.
type Config struct {
}

// stdoutExporter implements a traces exporter that writes telemetry data in JSON format to stdout.
type stdoutExporter struct {
	config   *Config
	settings exporter.Settings
}

// newStdoutExporter creates a new stdoutExporter with the provided configuration and settings.
func newStdoutExporter(cfg *Config, settings exporter.Settings) *stdoutExporter {
	return &stdoutExporter{
		config:   cfg,
		settings: settings,
	}
}

// ConsumeTraces marshals the incoming traces into JSON and writes them to stdout.
func (exp *stdoutExporter) ConsumeTraces(ctx context.Context, td ptrace.Traces) error {
	marshaler := ptrace.JSONMarshaler{}

	b, err := marshaler.MarshalTraces(td)
	if err != nil {
		return fmt.Errorf("failed to marshal traces: %w", err)
	}

	fmt.Fprintln(os.Stdout, string(b))

	return nil
}

// ConsumeMetrics marshals the incoming metrics into JSON and writes them to stdout.
func (exp *stdoutExporter) ConsumeMetrics(ctx context.Context, md pmetric.Metrics) error {
	marshaler := pmetric.JSONMarshaler{}

	b, err := marshaler.MarshalMetrics(md)
	if err != nil {
		return fmt.Errorf("failed to marshal metrics: %w", err)
	}

	fmt.Fprintln(os.Stdout, string(b))

	return nil
}

// ConsumeLogs marshals the incoming logs into JSON and writes them to stdout.
func (exp *stdoutExporter) ConsumeLogs(ctx context.Context, ld plog.Logs) error {
	marshaler := plog.JSONMarshaler{}

	b, err := marshaler.MarshalLogs(ld)
	if err != nil {
		return fmt.Errorf("failed to marshal logs: %w", err)
	}

	fmt.Fprintln(os.Stdout, string(b))

	return nil
}

// Shutdown performs any cleanup needed by the exporter.
func (exp *stdoutExporter) Shutdown(ctx context.Context) error {
	return nil
}
