package sub

import (
	"log"

	"github.com/okieoth/badginator/pkg/config"
	"github.com/okieoth/badginator/pkg/serve"
	"github.com/spf13/cobra"
)

var ServeCmd = &cobra.Command{
	Use:   "serve",
	Short: "starts the http server to deliver badges",
	Long:  "start an http server on the configured port and takes GET requests to deliver SVG badges",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := config.GetConfigFromFile(configFile)
		if err != nil {
			log.Fatalf("Cancel execution, because config file couldn't be loaded: %v", err)
		}
		serve.Start(cfg)
	},
}

var configFile string

func init() {
	ServeCmd.Flags().StringVarP(&configFile, "config", "c", "", "config file to adjust service parameters")
	ServeCmd.MarkFlagFilename("config")
	ServeCmd.MarkFlagRequired("config")
}
