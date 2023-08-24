package main

import (
	"fmt"
	"jira/cmd"
	"os"
)

func main() {
	var cmd, error = cmd.CreateRootCommand()

	if error != nil {
		fmt.Println("Failed to create root command")
		os.Exit(1)
	}

	cmd.Execute()
}
