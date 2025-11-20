package sub

import "github.com/spf13/cobra"

var InspectCmd = &cobra.Command{
	Use:   "inspect",
	Short: "Allows you to query the database of the service",
	Long:  "Provides a set of functions to get insights about the badge database",
	Run: func(cmd *cobra.Command, args []string) {
		// TODO
	},
}
