package monitoring

import (
	"log"
	"os"

	"github.com/newrelic/go-agent/v3/newrelic"
)

var NrApp *newrelic.Application

func SetupNewRelic(serviceName string, licenseKey string) {
	app, nrErr := newrelic.NewApplication(
		newrelic.ConfigAppName(serviceName),
		newrelic.ConfigLicense(licenseKey),
		newrelic.ConfigDebugLogger(os.Stdout),
	)
	if nrErr != nil {
		log.Println("Unable to create New Relic Monitoring agent. Reason:", nrErr)
	}
	NrApp = app
}
