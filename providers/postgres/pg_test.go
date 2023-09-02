package postgres_test

import (
	"context"
	"fmt"
	"path/filepath"
	"runtime"
	"testing"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"integro/providers/postgres"
)

func TestPGContainerLifecycle(t *testing.T) {
	image := "postgres:latest"
	container := postgres.NewPGContainer()

	t.Run("TestCreateAndTerminate", func(t *testing.T) {
		err := container.Create(image)
		require.NoError(t, err)

		// Check that the container is running
		assert.NotEmpty(t, container.Ip())
		assert.NotZero(t, container.Port())
		assert.NotEmpty(t, container.Conn())

		// Apply migrations
		migrationsSource, err := migrations()
		require.NoError(t, err)
		err = container.Migrate(migrationsSource)
		require.NoError(t, err)

		// create sqlx.DB instance
		ctx := context.Background()
		conn, err := pgx.Connect(ctx, container.Conn())
		require.NoError(t, err)
		defer func(conn *pgx.Conn, ctx context.Context) {
			err := conn.Close(ctx)
			if err != nil {
				t.Errorf("failed to close connection: %v", err)
			}
		}(conn, ctx)

		// Check that the database migrations were applied
		var version int
		err = conn.QueryRow(ctx, "SELECT version FROM schema_migrations ORDER BY version DESC LIMIT 1").Scan(&version)
		require.NoError(t, err)
		assert.Equal(t, 1, version)

		// select username,email from users table order by id
		var (
			username string
			email    string
		)
		err = conn.QueryRow(ctx, "SELECT username, email FROM users ORDER BY id LIMIT 1").Scan(&username, &email)
		require.NoError(t, err)
		assert.Equal(t, "user1", username)
		assert.Equal(t, "user1@example.com", email)

		err = container.Terminate()
		require.NoError(t, err)
	})
}

// migrations provide path to test migration.
func migrations() (string, error) {
	_, currFile, _, ok := runtime.Caller(0)
	if !ok {
		return "", fmt.Errorf("failed to get current file")
	}

	return fmt.Sprintf("file://%s", filepath.Join(currFile, "..", "..", "..", "test_data", "pg", "migrations")), nil
}
