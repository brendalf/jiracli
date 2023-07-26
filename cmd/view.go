package cmd

import (
	"fmt"
	"jira/api"

	"github.com/MakeNowJust/heredoc"
	"github.com/spf13/cobra"
)

func CreateViewCommand() *cobra.Command {
    var cmd = &cobra.Command{
        Use: "view <ticket_id>",
        Short: "View a jira ticket",
        Long: heredoc.Doc(`
            Display the title, description, person assigned and further information of a jira ticket.

            With '--web', open the ticket in a web browser instead.
        `),
        Args: cobra.ExactArgs(1),
        RunE: func(cmd *cobra.Command, args []string) error {
            var ticket_id = args[0]

            fmt.Printf("Looking for ticket %v\n", ticket_id)

            api.ViewTicket(ticket_id)

            return nil
        },
    }

    return cmd
}
