package wrappers_test

import (
	"context"
	"database/sql"
	"fmt"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	"github.com/stretchr/testify/require"

	"github.com/komandakycto/integro/wrappers"
)

func TestNew_MySQL(t *testing.T) {
	image := "mysql:8"

	// Create a new test container with MySQL image
	container, err := wrappers.New(image)
	require.NoError(t, err)

	// Apply migration from file
	err = container.Migrate(fmt.Sprintf("file://%s", "./../internal/test_data/mysql/migrations/"))
	require.NoError(t, err)

	// Connect to the database
	conn, err := sql.Open("mysql", container.Conn())
	require.NoError(t, err)

	// Query the users table
	rows, err := conn.QueryContext(context.Background(), "SELECT * FROM users")
	require.NoError(t, err)

	// Check if users table is not empty
	require.True(t, rows.Next())

	// Close the connection
	err = conn.Close()
	require.NoError(t, err)

	// Stop the container
	err = container.Stop()
	require.NoError(t, err)
}
