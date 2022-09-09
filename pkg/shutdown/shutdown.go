package shutdown

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/rs/zerolog/log"
)

// GracefulStop handles signal and graceful shutdown by executing callback function
// when signal is received callback is called followed after by os.Exit(0), it is responsibility of callback to handle timeout
// if second signal is received will terminate process by a call to os.Exit(1)
func GracefulStop(stop func()) {
	signalChan := make(chan os.Signal, 1)
	signal.Notify(
		signalChan,
		syscall.SIGHUP,  // kill -SIGHUP XXXX
		syscall.SIGINT,  // kill -SIGINT XXXX or Ctrl+c
		syscall.SIGTERM, // kill -SIGTERM XXXX
		syscall.SIGQUIT, // kill -SIGQUIT XXXX
	)

	<-signalChan
	log.Info().Msg("os.Interrupt - shutting down...")

	// terminate after second signal before callback is done
	go func() {
		<-signalChan
		log.Fatal().Msg("os.Interrupt - shutting down...")
	}()

	stop()

	os.Exit(0)
}
