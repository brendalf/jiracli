/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/MakeNowJust/heredoc"
	"github.com/spf13/cobra"
)

func CreateRootCommand() (*cobra.Command, error) {
    var rootCmd = &cobra.Command{
        Use: "jiracli <command> <subcommand> [flags]",
        Short: "Work seamlessly with Jira from the command line.",
        Long: heredoc.Doc(`
            Jira CLI aims to help developers quick update and list Jira tickets.
        `),
        Example: heredoc.Doc(`
            $ jiracli list
            $ jiracli view TICKET_ID
        `),
    }

    rootCmd.PersistentFlags().Bool("help", false, "Show help for command")

    rootCmd.AddCommand(CreateListCommand())
    rootCmd.AddCommand(CreateViewCommand())

    return rootCmd, nil
}


