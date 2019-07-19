package process

import (
	"os"
	"os/signal"
)

// BlockUntilSignal blocks until
func BlockUntilSignal(signals ...os.Signal) {
	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)

	signal.Notify(sigs, signals...)

	go func() {
		<-sigs
		done <- true
	}()

	<-done
}
