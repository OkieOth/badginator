package main

import (
	"fmt"

	"github.com/okieoth/badginator/cmd/sub"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "badginator",
	Short: "Tool to deliver badges",
	Long:  `Tool to provide a badge delivery service and some functions around their maintenance.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(":-)")
		// if cmdToCall, err := ic0bra.RunInteractiveWithHistory(cmd, "badginator"); err == nil {
		// 	if cmdToCall != nil {
		// 		cmdToCall.Run(cmdToCall, args)
		// 	}
		// } else {
		// 	fmt.Println("error while running in interactive mode:", err)
		// }
	},
}

func init() {
	rootCmd.AddCommand(sub.VersionCmd)
	rootCmd.AddCommand(sub.DummyCmd)
}

func main() {
	rootCmd.Execute()
}
