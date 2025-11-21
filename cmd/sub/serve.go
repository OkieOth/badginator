package sub

import "github.com/spf13/cobra"

var ServeCmd = &cobra.Command{
	Use:   "serve",
	Short: "starts the http server to deliver badges",
	Long:  "start an http server on the configured port and takes GET requests to deliver SVG badges",
	Run: func(cmd *cobra.Command, args []string) {
		// TODO
	},
}
