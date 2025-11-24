package config_test

import (
	"fmt"
	"testing"

	"github.com/okieoth/badginator/pkg/config"
	"github.com/stretchr/testify/require"
)

func TestGetConfigFromFile(t *testing.T) {
	tests := []struct {
		testFile string
		expected config.AppConfig
	}{
		{
			testFile: "../../resources/tests/testConfig.json",
			expected: config.AppConfig{
				Server: config.Server{
					Port: 8080,
				},
				Storage: config.Storage{
					Sqlite: config.Sqlite{
						File: "temp/testdb_01.db",
					},
				},
			},
		},
	}
	for i, test := range tests {
		require.FileExists(t, test.testFile, fmt.Sprintf("input file doesn't exist for test: %d", i))
		appCfg, err := config.GetConfigFromFile(test.testFile)
		require.NoError(t, err, fmt.Sprintf("cant's retrieve config for test: %d", i))
		require.Equal(t, test.expected, appCfg, fmt.Sprintf("got unexpected result for test: %d", i))
	}
}
