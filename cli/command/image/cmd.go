package image

import (
	"github.com/spf13/cobra"

	"github.com/docker/docker/cli"
	"github.com/docker/docker/cli/command"
)

// NewImageCommand returns a cobra command for `image` subcommands
func NewImageCommand(dockerCli *command.DockerCli) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "image",
		Short: "Manage images",
		Args:  cli.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			cmd.SetOutput(dockerCli.Err())
			cmd.HelpFunc()(cmd, args)
		},
	}
	cmd.AddCommand(
		NewBuildCommand(dockerCli),
		NewHistoryCommand(dockerCli),
		NewImportCommand(dockerCli),
		NewLoadCommand(dockerCli),
		NewPullCommand(dockerCli),
		NewPushCommand(dockerCli),
		NewSaveCommand(dockerCli),
		NewTagCommand(dockerCli),
		newListCommand(dockerCli),
		newRemoveCommand(dockerCli),
		newInspectCommand(dockerCli),
		NewPruneCommand(dockerCli),
	)
	return cmd
}
