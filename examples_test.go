// Provides examples on how to use this test library.
package log_test

import (
	"github.com/guiln/starter-log/logger"
	"github.com/guiln/starter-log/messages"
)

func ExampleNewBuilder_Logger_WithoutCorrlationId() {
	loggr := logger.NewBuilder(). // creates a new builder with defaults
					WithLogFlags(0). // removes log flags to assert output
					Build()          // builds the logger

	loggr.Debug(messages.New().WithCorrelationId("2323424").WithMessage("this is a log").Build())

	// Output:
	// [DEBUG]{"correlation_id":"2323424","message":"this is a log"}
}

func ExampleNewBuilder_Logger_WithCorrelationId() {
	loggr := logger.NewBuilder(). // creates a new builder with defaults
					WithLogFlags(0).              // removes log flags to assert output
					WithCorrelationId("2323424"). // setting correlation id so we expect the same result as Example above
					Build()                       // builds the logger.

	loggr.Debug(messages.New().WithMessage("this is a log").Build())

	// Output:
	// [DEBUG]{"correlation_id":"2323424","message":"this is a log"}
}
