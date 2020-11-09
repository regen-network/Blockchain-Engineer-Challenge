package main

import (
	"os"

	"github.com/amaurymartiny/bec/app/regen/cmd"
)

func main() {
	rootCmd, _ := cmd.NewRootCmd()
	if err := cmd.Execute(rootCmd); err != nil {
		os.Exit(1)
	}
}
