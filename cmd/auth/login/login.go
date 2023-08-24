package login

import (
	"github.com/spf13/cobra"
	"github.com/MakeNowJust/heredoc"
)

func CreateLoginCommand() *cobra.Command {
    var hostname string

    var cmd = &cobra.Command{
        Use: "login",
        Short: "Authenticate the cli with Jira",
        Long: heredoc.Doc(`
            Authenticate the cli with Jira.

            This command will prompt you for your Jira username and token.
            The token can be generated at https://id.atlassian.com/manage/api-tokens.
        `),
        Example: heredoc.Doc(`
            # start interactive login setup
            $ jira auth login

            # authenticate with a specific Jira instance
            $ jira auth login --hostname enterprise.atlassian.net
        `),
        RunE: func(cmd *cobra.Command, args []string) error {
            return nil
        },
    }

    cmd.Flags().StringVarP(&hostname, "hostname", "h", "", "Jira hostname to authenticate with")

    return cmd
}
