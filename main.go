package main

import (
	"context"
	"os"

	"github.com/sirupsen/logrus"

	"github.com/jfk9w-go/flu"

	"github.com/jfk9w-go/homebot/ext/tinkoff/sync"

	"github.com/jfk9w-go/homebot/app"
	"github.com/jfk9w-go/homebot/ext/tinkoff"
)

var GitCommit = "dev"

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	app, err := app.Create(GitCommit, flu.DefaultClock, flu.File(os.Args[1]))
	if err != nil {
		logrus.Fatal(err)
	}

	defer flu.CloseQuietly(app)
	if err := app.ConfigureLogging(); err != nil {
		logrus.Fatal(err)
	}

	app.ApplyPlugins(tinkoff.Extension{sync.Accounts, sync.TradingOperations, sync.PurchasedSecurities})
	if err := app.Run(ctx); err != nil {
		logrus.Fatal(err)
	}

	flu.AwaitSignal()
}
