package main

import (
	"fmt"
	"os"
)

func main() {
	app := NewCliApp()
	err := app.Run(os.Args)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		os.Exit(1)
	}
}
