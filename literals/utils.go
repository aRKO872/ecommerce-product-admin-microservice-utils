package literals

var (
	MegaByte                   = 1024 * 1024

	// context keys
	ContextKeyCorrelationID    = "correlation_id"

	// Printing Templates
	LogTemplate                = "[AppID: %s] [Level: %s] [Message: %s] [Timestamp: %s]"

	// logging levels
	LogLevelInfo               = "INFO"
	LogLevelWarn               = "WARN"
	LogLevelError              = "ERROR"
)