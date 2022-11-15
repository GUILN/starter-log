package messages

// LogMessage
// provides basic structure for company log message.
type LogMessage struct {
	CorrelationId string `json:"correlation_id"`
	Message       string `json:"message,omitempty"`
}

// LogMessageBuilder
//
// To provide a fluent interface for building company log message.
type LogMessageBuilder struct {
	companyLog *LogMessage
}

// New
//
// # Returns a new LogMessageBuilder
//
// @param logMessage can be changed later using `WithMessage` method.
func New(logMessage string) *LogMessageBuilder {
	return &LogMessageBuilder{
		companyLog: &LogMessage{
			Message: logMessage,
		},
	}
}

func (clb *LogMessageBuilder) WithCorrelationId(correlationId string) *LogMessageBuilder {
	clb.companyLog.CorrelationId = correlationId
	return clb
}

func (clb *LogMessageBuilder) WithMessage(message string) *LogMessageBuilder {
	clb.companyLog.Message = message
	return clb
}

// Build
//
// Builds LogMessage required by company logger.
func (clb *LogMessageBuilder) Message() *LogMessage {
	return clb.companyLog
}
