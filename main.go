package main

import (
	"github.com/akhidrb/toggl-cards/cmd"
	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{}
	rootCmd.AddCommand(
		cmd.API(),
	)
	cobra.CheckErr(rootCmd.Execute())
}
