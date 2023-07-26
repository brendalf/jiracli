package cmd

import (
	"jira/api"

	"github.com/MakeNowJust/heredoc"
	"github.com/spf13/cobra"
)

func CreateListCommand() *cobra.Command {
    var cmd = &cobra.Command{
        Use: "list",
        Short: "List all jira tickets",
        Long: heredoc.Doc(`
            List jira tickets ...
        `),
        RunE: func(cmd *cobra.Command, args []string) error {
            api.ListTickets()
            return nil
        },
    }

    return cmd
}
