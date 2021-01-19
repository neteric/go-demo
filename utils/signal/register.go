package signal_demo

import (
	"golang.org/x/sys/unix"
	"os"
	"os/signal"
	"syscall"
)

var onlyOnceHandler = make(chan struct{})

var shutdownSignals = []os.Signal{os.Interrupt, syscall.SIGTERM, unix.SIGTERM}

func SetupSignalHandler() (stopChan chan struct{}) {

	// panic when call SetupSignalHandler twice
	close(onlyOnceHandler)

	var stop = make(chan struct{})

	var c = make(chan os.Signal, 2)
	signal.Notify(c, shutdownSignals...)

	go func() {
		<-c
		close(stop)
		<-c
		os.Exit(1)

	}()

	return stop

}
