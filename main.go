/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package main

import (
    "os"
    "fmt"
    "jira/cmd"
)

func main() {
    var cmd, error = cmd.CreateRootCommand()

    if error != nil {
        fmt.Println("Failed to create root command")
        os.Exit(1)
    }

    cmd.Execute()
}
