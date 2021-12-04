package main

import (
	"github.com/marcelobarreto/banking/app"
	"github.com/marcelobarreto/banking/logger"
)

func main() {
	logger.Info("Starting application")
	app.Start()
}
