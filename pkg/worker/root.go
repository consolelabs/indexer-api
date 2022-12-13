package worker

import (
	"github.com/spf13/cobra"

	"github.com/consolelabs/indexer-api/pkg/logger"
)

type rootCmd struct {
	log logger.Logger

	rootCmd *cobra.Command
}

var c = &rootCmd{}

func Init(log logger.Logger) {
	c.log = log
	c.rootCmd = &cobra.Command{Use: "worker"}

	// add version command
	c.rootCmd.AddCommand(&cobra.Command{
		Use:     "version",
		Aliases: []string{"v"},
		Short:   "Print the version number of indexer",
		Run: func(cmd *cobra.Command, args []string) {
			c.log.Infof("indexer version: %s", "0.0.1")
		},
	})
}

func AddCommand(command ICommand) {
	c.rootCmd.AddCommand(&cobra.Command{
		Use:     command.Name(),
		Aliases: command.Aliases(),
		Short:   command.Usage(),
		RunE: func(cmd *cobra.Command, args []string) error {
			c.log.Info("Executing command: " + command.Name())
			if err := command.Run(); err != nil {
				return err
			}
			return nil
		},
	})
}

func Execute() error {
	return c.rootCmd.Execute()
}
