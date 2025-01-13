package main

import (
	"fmt"
	"os"
)

func main() {
	e, err := InitializeEvent("this is the message!")
	if err != nil {
		fmt.Printf("failed to create event: %s\n", err)
		os.Exit(2)
	}
	e.Start()
}
