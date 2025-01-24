package hermine

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"path/filepath"
	"regexp"
	"testing"
)

const belMngrEmptySqLiteDatabaseFileName = "BelegManager_empty.db4"

func Test_startBelegManagerSQLiteDb(t *testing.T) {
	belegManagerSqLiteDBAbsoluteFilePath := filepath.Join("testdata", belMngrEmptySqLiteDatabaseFileName)

	db := StartBelegManagerSQLiteDB(belegManagerSqLiteDBAbsoluteFilePath)
	require.NotNil(t, db)

	CloseDB(db)
	pingErr := db.Ping()
	assert.Error(t, pingErr)
}

func Test_newBmDocUuid(t *testing.T) {
	bmDocUUID := newBmDocUUID()

	match := regexp.MustCompile(`^\{[0-9a-fA-F-]{36}}$`).MatchString(bmDocUUID)
	require.True(t, match)
}
