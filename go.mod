module github.com/open-telemetry/opentelemetry-collector-contrib

go 1.12

require (
	github.com/client9/misspell v0.3.4
	github.com/google/addlicense v0.0.0-20190907113143-be125746c2c4
	github.com/open-telemetry/opentelemetry-collector v0.2.1-0.20191016224815-dfabfb0c1d1e
	github.com/open-telemetry/opentelemetry-collector-contrib/exporter/signalfxexporter v0.0.0-00010101000000-000000000000
	github.com/open-telemetry/opentelemetry-collector-contrib/exporter/stackdriverexporter v0.0.0-20191021165924-bb954188ac10
	github.com/open-telemetry/opentelemetry-collector-contrib/receiver/collectdreceiver v0.0.0-00010101000000-000000000000
	github.com/open-telemetry/opentelemetry-collector-contrib/receiver/zipkinscribereceiver v0.0.0-20191021165924-bb954188ac10
	golang.org/x/lint v0.0.0-20190930215403-16217165b5de
	golang.org/x/tools v0.0.0-20191029190741-b9c20aec41a5
	honnef.co/go/tools v0.0.1-2019.2.3
)

replace github.com/open-telemetry/opentelemetry-collector-contrib/receiver/collectdreceiver => ./receiver/collectdreceiver

replace github.com/open-telemetry/opentelemetry-collector-contrib/exporter/signalfxexporter => /home/owais/gohack/github.com/pjanotti/opentelemetry-service-contrib/exporter/signalfxexporter
