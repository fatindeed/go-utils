// Package signal provides graceful shutdown for your apps.
package signal

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

var (
	dispatcher *signalDispatcher
	mutex      sync.Mutex
)

type signalDispatcher struct {
	sigs    chan os.Signal
	stop    chan bool
	lastSig os.Signal
}

func (s *signalDispatcher) Listen() {
	for {
		select {
		case <-s.stop:
			signal.Stop(s.sigs)
			close(s.stop)
			close(s.sigs)
			return
		case sig := <-s.sigs:
			fmt.Println("\r")
			s.lastSig = sig
		}
	}
}

func initDispatcher(sig ...os.Signal) {
	mutex.Lock()
	defer mutex.Unlock()

	if dispatcher != nil {
		return
	}
	dispatcher = new(signalDispatcher)
	dispatcher.sigs = make(chan os.Signal, 1)
	dispatcher.stop = make(chan bool, 1)
	if len(sig) == 0 {
		signal.Notify(dispatcher.sigs, syscall.SIGINT, syscall.SIGTERM)
	} else {
		signal.Notify(dispatcher.sigs, sig...)
	}
	go dispatcher.Listen()
}

// Listen can customize signals to be notified.
func Listen(sig ...os.Signal) {
	if dispatcher == nil {
		initDispatcher(sig...)
		return
	}
	signal.Stop(dispatcher.sigs)
	signal.Notify(dispatcher.sigs, sig...)
}

// Dispatch will check if new signal received, and return a SignalError if yes.
func Dispatch() error {
	initDispatcher()
	if dispatcher.lastSig == nil {
		return nil
	}
	return SignalError{Signal: dispatcher.lastSig}
}

// Stop singal dispatcher if exists.
func Stop() {
	if dispatcher != nil {
		dispatcher.stop <- true
	}
}
