package main

import (
	"fmt"
	"time"

	"github.com/getsentry/sentry-go"
	"github.com/mfbmina/enxame/cmd"
)

func main() {
	defer panicRecover()

	err := sentry.Init(sentry.ClientOptions{
		Dsn:              "",
		TracesSampleRate: 1.0,
	})
	err = cmd.Execute()
	if err != nil {
		sentry.CaptureException(err)
	}
}

func panicRecover() {
	err := recover()
	if err != nil {
		fmt.Println("Something weird happened. Please report this to the developers.")
		sentry.CurrentHub().Recover(err)
	}

	sentry.Flush(time.Second * 5)
}
