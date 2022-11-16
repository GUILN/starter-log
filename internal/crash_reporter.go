// Setup crasher reporter centralizer, such as sentry.
package internal

import (
	"time"

	"github.com/getsentry/sentry-go"
	"github.com/guiln/starter-log/messages"
)

type SentryCrashReporterConfiguration struct {
	Dsn             string
	TraceSampleRate float64
}

type SentryCrashReporter struct {
	config *SentryCrashReporterConfiguration
}

func NewSentryCrashReporter(config *SentryCrashReporterConfiguration) *SentryCrashReporter {
	return &SentryCrashReporter{
		config: config,
	}
}

func (scr *SentryCrashReporter) CaptureMessage(message *messages.LogMessage) {
	sentry.CaptureMessage(message.AsJson())
}

func (scr *SentryCrashReporter) CaptureException(exc error) {
	sentry.CaptureException(exc)
}

// InitCrashReporter
//
// If config is passed as nil then will default to production configuration.
// @param *CrasherReporterConfiguration
//
// Returns
// (func() bool, error) - function that returns is meant to be called in the defer manner.
// So it can flush before application shutdown.
func (scr *SentryCrashReporter) InitCrashReporter() (func() bool, error) {
	config := scr.config
	// Set defaults
	if config == nil {
		config = &SentryCrashReporterConfiguration{
			Dsn:             "https://1f4807f97cdd40059b73be6f24e6d25c@o4504168838070272.ingest.sentry.io/4504170441146368",
			TraceSampleRate: 1,
		}
	}

	if err := sentry.Init(sentry.ClientOptions{
		Dsn:              scr.config.Dsn,
		TracesSampleRate: config.TraceSampleRate,
	}); err != nil {
		return nil, err
	}

	// Function to be called at the end of the execution of the main function.
	// With a `defer` statement.
	flushFunc := func() bool {
		return sentry.Flush(2 * time.Second)
	}
	return flushFunc, nil
}
