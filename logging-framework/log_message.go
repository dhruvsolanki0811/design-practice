package loggingframework

import (
	"fmt"
	"time"
)

type LogMessage struct {
	LogLevel  LogLevel
	Message   string
	Timestamp int
}

func NewMessage(logLevel LogLevel, message string) LogMessage {
	return LogMessage{
		Message:   message,
		LogLevel:  logLevel,
		Timestamp: int(time.Now().UnixMilli()),
	}
}

func (l LogMessage) getLogString() string {
	return fmt.Sprintf("[%s] %d - %s", l.LogLevel, l.Timestamp, l.Message)
}
