package signal

import (
	"fmt"
	"os"
)

// SignalError is an error with os.Signal included.
type SignalError struct {
	os.Signal
}

// Error returns the error message.
func (err SignalError) Error() string {
	return fmt.Sprintf("signal: %s", err.Signal.String())
}
