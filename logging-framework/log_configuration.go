package loggingframework

type LogConfiguration struct {
	Level    LogLevel
	Appender LogAppender
}

func NewLogConfiguration(level string, LogAppender LogAppender) LogConfiguration {
	return LogConfiguration{
		Level:    ToLogLevel(level),
		Appender: LogAppender,
	}
}
