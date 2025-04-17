module github.com/icco/unifi

go 1.21

toolchain go1.24.2

require (
	github.com/icco/gutil v0.0.0-20220221170217-9aa326c389ec
	github.com/unpoller/unifi v0.4.3
)

require (
	github.com/benbjohnson/clock v1.3.0 // indirect
	github.com/brianvoe/gofakeit/v6 v6.28.0 // indirect
	github.com/go-chi/chi/v5 v5.0.7 // indirect
	github.com/icco/zapdriver v1.4.0 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	go.uber.org/atomic v1.9.0 // indirect
	go.uber.org/multierr v1.7.0 // indirect
	go.uber.org/zap v1.21.0 // indirect
	golang.org/x/net v0.24.0 // indirect
)

// Add explicit replace directives for OpenTelemetry packages
replace (
	go.opentelemetry.io/otel => go.opentelemetry.io/otel v1.4.1
	go.opentelemetry.io/otel/sdk => go.opentelemetry.io/otel/sdk v1.4.1
	go.opentelemetry.io/otel/sdk/export/metric => go.opentelemetry.io/otel/sdk/export/metric v0.26.0
	go.opentelemetry.io/otel/sdk/metric => go.opentelemetry.io/otel/sdk/metric v0.26.0
	go.opentelemetry.io/otel/trace => go.opentelemetry.io/otel/trace v1.4.1
)
