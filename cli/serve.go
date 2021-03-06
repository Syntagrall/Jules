package cli

import (
	"github.com/bbuck/dragon-mud/logger"
	"github.com/bbuck/dragon-mud/random"
	"github.com/spf13/cobra"
)

var (
	serveCmd = &cobra.Command{
		Use:   "serve",
		Short: "Start the DragonMUD server.",
		Long: `Starts the DragonMUD Game server to listen for new player connections.
All lifecycle scripts will be notified during boot and the configuration
information will be processed.`,
		Run: func(cmd *cobra.Command, args []string) {
			logger.Log().Infof("A %s dragon arrives to serve you today.", getDragonColor())
		},
	}

	dragonColors = []string{
		"{black+h:white+h}black{reset}",
		"{red+h}red{reset}",
		"{green+h}green{reset}",
		"{yellow+h}gold{reset}",
		"{blue+h}blue{reset}",
		"{white+h}white{reset}",
		"{white+u}silver{reset}",
	}
)

func init() {
	RootCmd.AddCommand(serveCmd)
}

func getDragonColor() string {
	index := random.Int(len(dragonColors))

	return dragonColors[index]
}
