package main

import (
	"os"

	"capital-gains/cmd/cli/internal"
)

func main() {
	if err := internal.Run(os.Stdin, os.Stdout); err != nil {
		panic(err)
	}
}
