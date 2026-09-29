module go.icco.me/unifi

go 1.26.0

require (
	github.com/unpoller/unifi v0.4.3
	go.icco.me/gutil v1.0.27-0.20260929105600-f3e11752b837
)

require (
	github.com/brianvoe/gofakeit/v6 v6.28.0 // indirect
	github.com/go-chi/chi/v5 v5.3.2 // indirect
	go.opentelemetry.io/otel v1.46.0 // indirect
	go.opentelemetry.io/otel/trace v1.46.0 // indirect
	go.uber.org/multierr v1.11.0 // indirect
	go.uber.org/zap v1.28.0 // indirect
	golang.org/x/net v0.59.0 // indirect
)

// Add explicit replace directives for OpenTelemetry packages
replace (
	go.opentelemetry.io/otel => go.opentelemetry.io/otel v1.4.1
	go.opentelemetry.io/otel/sdk => go.opentelemetry.io/otel/sdk v1.4.1
	go.opentelemetry.io/otel/sdk/export/metric => go.opentelemetry.io/otel/sdk/export/metric v0.26.0
	go.opentelemetry.io/otel/sdk/metric => go.opentelemetry.io/otel/sdk/metric v0.26.0
	go.opentelemetry.io/otel/trace => go.opentelemetry.io/otel/trace v1.4.1
)
