package logger

import "github.com/guiln/starter-log/internal"

// CreateSentryCrashReporter
//
// If config is passed as nil then will default to production configuration.
func CreateSentryCrashReporter(config *internal.SentryCrashReporterConfiguration) CompanyCrashReporter {
	return internal.NewSentryCrashReporter(config)
}
