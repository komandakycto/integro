package postgres

import (
	"context"
	"errors"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

const defaultExposedPort = "5432/tcp"
const defaultPort = "5432"
const defaultPassword = "example"

type PGContainer struct {
	// ip is an ip address of the postgres database container.
	ip string
	// port is a port of the postgres database container.
	port int
	// container is a postgres database container.
	container testcontainers.Container
	// conn is a connection string to the postgres database.
	conn string
}

// NewPGContainer creates a new postgres database container.
func NewPGContainer() *PGContainer {
	return &PGContainer{}
}

// Create creates a default postgres database container.
func (c *PGContainer) Create(image string) error {
	req := testcontainers.ContainerRequest{
		Image:        image,
		ExposedPorts: []string{defaultExposedPort},
		Env: map[string]string{
			"POSTGRES_PASSWORD":         defaultPassword,
			"POSTGRES_HOST_AUTH_METHOD": "trust",
		},
		WaitingFor: wait.ForListeningPort(defaultExposedPort),
	}

	ctx := context.Background()
	container, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		return fmt.Errorf("failed to start postgresql container: %w", err)
	}

	ip, err := container.Host(ctx)
	if err != nil {
		return fmt.Errorf("failed to get postgresql container ip: %w", err)
	}
	port, err := container.MappedPort(ctx, defaultPort)
	if err != nil {
		return fmt.Errorf("failed to get postgresql container port: %w", err)
	}

	c.container = container
	c.ip = ip
	c.port = port.Int()
	c.conn = "postgres://postgres:" + defaultPassword + "@" + ip + ":" + port.Port() + "/postgres?sslmode=disable"

	return nil
}

// Terminate terminates the postgres database container.
func (c *PGContainer) Terminate() error {
	if !c.container.IsRunning() {
		return fmt.Errorf("container is not running")
	}

	err := c.container.Terminate(context.Background())
	if err != nil {
		return fmt.Errorf("failed to terminate container, err: %w", err)
	}

	return nil
}

// Conn returns a connection string to the postgres database.
func (c *PGContainer) Conn() string {
	return c.conn
}

// Ip returns an ip address of the postgres database container.
func (c *PGContainer) Ip() string {
	return c.ip
}

// Port returns a port of the postgres database container.
func (c *PGContainer) Port() int {
	return c.port
}

// Migrate migrates the postgres database.
func (c *PGContainer) Migrate(source string) error {
	if !c.container.IsRunning() || c.Conn() == "" {
		return fmt.Errorf("no connection")
	}

	m, err := migrate.New(source, c.Conn())
	if err != nil {
		return fmt.Errorf("failed to create migrator: %w", err)
	}

	err = m.Up()
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return fmt.Errorf("failed to migrate: %w", err)
	}

	sourceErr, databaseErr := m.Close()
	if err != nil {
		return fmt.Errorf("failed to close migrator: %w, %w", sourceErr, databaseErr)
	}

	return nil
}
