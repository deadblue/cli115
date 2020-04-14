package main

import (
	"go.dead.blue/cli115"
)

func main() {
	err := cli115.Run()
	if err != nil {
		panic(err)
	}
}
