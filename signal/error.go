package signal

import (
	"fmt"
	"os"
)

type signalError struct {
	os.Signal
}

func (err signalError) Error() string {
	return fmt.Sprintf("signal: %s", err.Signal.String())
}
