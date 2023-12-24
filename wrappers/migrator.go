package wrappers

import (
	"errors"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
)

type TestContainerMigrator struct {
	conn string
}

func NewTestContainerMigrator(connString string) *TestContainerMigrator {
	return &TestContainerMigrator{
		conn: connString,
	}
}

func (m *TestContainerMigrator) Migrate(source string) error {
	migrator, err := migrate.New(source, m.conn)
	if err != nil {
		return fmt.Errorf("failed to create migrator: %w", err)
	}
	defer func(m *migrate.Migrate) {
		_, _ = m.Close()
	}(migrator)
	if err := migrator.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return fmt.Errorf("failed to apply migration: %w", err)
	}

	return nil
}
