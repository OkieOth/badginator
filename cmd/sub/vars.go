package sub

import "github.com/spf13/cobra"

var databaseFile string

func initDefaultVars(cmd *cobra.Command) {
	cmd.Flags().StringVarP(&databaseFile, "db_file", "d", "", "Path to the db file to use (required)")
	cmd.MarkFlagFilename("db_file")
	cmd.MarkFlagRequired("db_file")
}
