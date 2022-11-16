package logger

import "github.com/guiln/starter-log/messages"

type CrashReporterShutdownFunction = func() bool

// CompanyCrashReporter
//
// Provides a standard interface for crash reporting.
// Initially used as a `sentry wrapper`.
type CompanyCrashReporter interface {
	// CaptureMessage
	CaptureMessage(*messages.LogMessage)
	// CaptureException
	CaptureException(exc error)
	// InitCrashReporter
	//
	// Initializes crasher reporter and returns a
	// function to hadle the crash reporter shutdown.
	// Use the function as with a `defer` statement.
	InitCrashReporter() (CrashReporterShutdownFunction, error)
}
