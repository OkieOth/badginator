package sub

import (
	"github.com/okieoth/badginator/pkg/dummy"
	"github.com/spf13/cobra"
)

var DummyCmd = &cobra.Command{
	Use:   "dummy",
	Short: "Only to test if the sqlite handling works",
	Long:  "Opens an sqlite file and prints the contained tables",
	Run: func(cmd *cobra.Command, args []string) {
		dummy.GetTables(databaseFile)
	},
}

func init() {
	initDefaultVars(DummyCmd)
}
