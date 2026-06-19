package loggingframework

import "sync"

type Logger struct {
	Config LogConfiguration
	mu     sync.Mutex
}

var (
	instance *Logger
	once     sync.Once
)

func NewLogger(logConfiguration LogConfiguration) *Logger {
	once.Do(func() {
		instance = &Logger{
			Config: logConfiguration,
		}
	})
	return instance
}

func GetLogger() *Logger {
	return instance
}

func (l *Logger) log(logLevel LogLevel, message string) error {
	l.mu.Lock()
	defer l.mu.Unlock()

	config := l.Config

	if config.Level > logLevel {
		return nil
	}

	err := config.Appender.Append(NewMessage(logLevel, message))
	if err != nil {
		return err
	}
	return nil
}

func (l *Logger) Debug(message string) error {
	return l.log(DEBUG, message)
}

func (l *Logger) Info(message string) error {
	return l.log(INFO, message)
}

func (l *Logger) Warning(message string) error {
	return l.log(WARN, message)
}

func (l *Logger) Error(message string) error {
	return l.log(ERROR, message)
}

func (l *Logger) Fatal(message string) error {
	return l.log(FATAL, message)
}
