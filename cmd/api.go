package cmd

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func API() *cobra.Command {
	return &cobra.Command{
		Use: "api",
		Run: func(cmd *cobra.Command, args []string) {
			log.Info("Starting toggl-cards api")
			runAPI()
		},
	}
}

func runAPI() {
}
