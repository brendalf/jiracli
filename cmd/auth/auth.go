package cmd

import (
	"github.com/spf13/cobra"
    authLoginCmd "jira/cmd/auth/login"
)

func CreateAuthCommand() *cobra.Command {
    var cmd = &cobra.Command{
        Use: "auth <command>",
        Short: "Authenticate the cli with Jira",
    }

    cmd.AddCommand(authLoginCmd.CreateLoginCommand())

    return cmd
}
