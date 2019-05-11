package main

import (
	"github.com/discipleRaceServer/cmd"
	"os"
	"fmt"
)

func main() {
	// Execute root cmd
	if err := cmd.RootCmd.Execute(); err != nil {
		fmt.Println(fmt.Sprintf("Error: %s", err))
		os.Exit(-1)
	}
}
