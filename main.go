package main

import (
	"fmt"
	"time"

	"github.com/getsentry/sentry-go"
	"github.com/mfbmina/enxame/cmd"
	"github.com/spf13/viper"
)

func main() {
	defer panicRecover()
	viper.AutomaticEnv()
	viper.SetDefault("SENTRY_DSN", "https://2ed32e0bcb336bf19065411bdbe0887f@o4506350446575616.ingest.sentry.io/4506350457847808")

	sentry.Init(sentry.ClientOptions{
		Dsn:              viper.Get("SENTRY_DSN").(string),
		TracesSampleRate: 1.0,
	})
	err := cmd.Execute()
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
