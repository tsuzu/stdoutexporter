package stdoutexporter

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/exporter"
	"go.opentelemetry.io/collector/exporter/exporterhelper"
)

// NewFactory creates a new exporter factory for the stdoutexporter.
func NewFactory() exporter.Factory {
	return exporter.NewFactory(
		"stdoutexporter",
		createDefaultConfig,
		exporter.WithTraces(createTracesExporter, component.StabilityLevelDevelopment),
		exporter.WithMetrics(createMetricsExporter, component.StabilityLevelDevelopment),
		exporter.WithLogs(createLogsExporter, component.StabilityLevelDevelopment),
	)
}

func createDefaultConfig() component.Config {
	return &Config{}
}

func createTracesExporter(
	ctx context.Context,
	set exporter.CreateSettings,
	cfg component.Config,
) (exporter.Traces, error) {
	exp := newStdoutExporter(cfg.(*Config), set)
	return exporterhelper.NewTracesExporter(ctx, set, cfg, exp.ConsumeTraces, exporterhelper.WithShutdown(exp.Shutdown))
}

func createMetricsExporter(
	ctx context.Context,
	set exporter.CreateSettings,
	cfg component.Config,
) (exporter.Metrics, error) {
	exp := newStdoutExporter(cfg.(*Config), set)
	return exporterhelper.NewMetricsExporter(ctx, set, cfg, exp.ConsumeMetrics, exporterhelper.WithShutdown(exp.Shutdown))
}

func createLogsExporter(
	ctx context.Context,
	set exporter.CreateSettings,
	cfg component.Config,
) (exporter.Logs, error) {
	exp := newStdoutExporter(cfg.(*Config), set)
	return exporterhelper.NewLogsExporter(ctx, set, cfg, exp.ConsumeLogs, exporterhelper.WithShutdown(exp.Shutdown))
}
