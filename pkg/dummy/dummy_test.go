package dummy_test

import (
	"testing"

	"github.com/okieoth/badginator/pkg/dummy"
	"github.com/stretchr/testify/require"
)

func TestGetTables(t *testing.T) {
	dbFile := "../../temp/test.db"
	err := dummy.GetTables(dbFile)
	require.NoError(t, err)
}
