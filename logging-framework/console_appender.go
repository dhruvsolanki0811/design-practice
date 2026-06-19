package loggingframework

import (
	"fmt"
	"sync"
)

type ConsoleAppender struct {
	mu sync.Mutex
}

func NewConsoleAppender() *ConsoleAppender {
	return &ConsoleAppender{}
}

func (c *ConsoleAppender) Append(logMessage LogMessage) error {
	line := logMessage.getLogString()

	c.mu.Lock()
	fmt.Println(line)
	c.mu.Unlock()

	return nil
}
