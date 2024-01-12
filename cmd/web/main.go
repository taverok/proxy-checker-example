package main

import (
	"github.com/taverok/proxy-checker-example/service/checker"
)

func main() {
	app, err := checker.NewApp()
	if err != nil {
		panic(err)
	}

	app.Listen()
}
