package cmd

import (
    "fmt"
    "runtime/debug"

	"github.com/MakeNowJust/heredoc"
	"github.com/spf13/cobra"
)

func CreateVersionCommand() *cobra.Command {
    var cmd = &cobra.Command{
        Use: "version",
        Short: "Print the version number of Jira CLI",
        Long: heredoc.Doc(`Print the version number of Jira CLI`),
        RunE: func(cmd *cobra.Command, args []string) error {
            if info, ok := debug.ReadBuildInfo(); ok && info.Main.Version != "(devel)" {
                buildVersion := info.Main.Version
                fmt.Printf("jira-cli version %s\n", buildVersion)
            }
            return nil
        },
    }

    return cmd
}
