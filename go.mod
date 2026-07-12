module github.com/icco/unifi

go 1.25.0

require (
	github.com/icco/gutil v0.0.0-20250215014032-7b1b73930901
	github.com/unpoller/unifi v0.4.3
)

require (
	github.com/brianvoe/gofakeit/v6 v6.28.0 // indirect
	github.com/go-chi/chi/v5 v5.2.4 // indirect
	github.com/icco/zapdriver v1.4.0 // indirect
	go.uber.org/multierr v1.11.0 // indirect
	go.uber.org/zap v1.27.0 // indirect
	golang.org/x/net v0.55.0 // indirect
)

// Add explicit replace directives for OpenTelemetry packages
replace (
	go.opentelemetry.io/otel => go.opentelemetry.io/otel v1.4.1
	go.opentelemetry.io/otel/sdk => go.opentelemetry.io/otel/sdk v1.4.1
	go.opentelemetry.io/otel/sdk/export/metric => go.opentelemetry.io/otel/sdk/export/metric v0.26.0
	go.opentelemetry.io/otel/sdk/metric => go.opentelemetry.io/otel/sdk/metric v0.26.0
	go.opentelemetry.io/otel/trace => go.opentelemetry.io/otel/trace v1.4.1
)
