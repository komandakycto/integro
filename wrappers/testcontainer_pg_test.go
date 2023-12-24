package wrappers_test

import (
	"context"
	"fmt"
	"testing"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v5"
	"github.com/stretchr/testify/require"

	"github.com/komandakycto/integro/wrappers"
)

func TestNew_Postgres(t *testing.T) {
	image := "postgres:15"

	// Create a new test container with postgres image
	container, err := wrappers.New(image)
	require.NoError(t, err)

	// Apply migration from file
	err = container.Migrate(fmt.Sprintf("file://%s", "./../internal/test_data/pg/migrations/"))
	require.NoError(t, err)

	// Connect to the database
	conn, err := pgx.Connect(context.Background(), container.Conn())
	require.NoError(t, err)

	// Query the users table
	rows, err := conn.Query(context.Background(), "SELECT * FROM users")
	require.NoError(t, err)

	// Check if users table is not empty
	require.True(t, rows.Next())

	// Close the connection
	err = conn.Close(context.Background())
	require.NoError(t, err)

	// Stop the container
	err = container.Stop()
	require.NoError(t, err)
}
