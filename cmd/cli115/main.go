package main

import (
	"go.dead.blue/cli115/internal/app"
)

func main() {
	err := app.Run()
	if err != nil {
		panic(err)
	}
}
