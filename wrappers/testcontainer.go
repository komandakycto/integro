package wrappers

import (
	"context"
	"fmt"

	"github.com/docker/go-connections/nat"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"

	"github.com/komandakycto/integro"
)

type TestContainer struct {
	migrator      integro.Migrator
	testcontainer testcontainers.Container

	conn string
}

func New(image string) (integro.Container, error) {
	port := defaultPort(image)
	if port == "" {
		// TODO add more friendly error
		return nil, fmt.Errorf("failed to get default port for image: %s", image)
	}

	tcpPort := fmt.Sprintf("%s/tcp", port)
	natPort, err := nat.NewPort("tcp", port)
	if err != nil {
		return nil, fmt.Errorf("failed to create nat port: %w", err)
	}

	req := testcontainers.ContainerRequest{
		Image:        image,
		ExposedPorts: []string{tcpPort},
		WaitingFor:   wait.ForListeningPort(natPort),
		Env:          defaultEnv(image),
	}

	ctx := context.Background()
	container, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to start container: %w", err)
	}

	hostIp, err := container.Host(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get container ip: %w", err)
	}
	mappedPort, err := container.MappedPort(ctx, natPort)
	if err != nil {
		return nil, fmt.Errorf("failed to get container port: %w", err)
	}

	conn := fmt.Sprintf(defaultConnection(image), hostIp, mappedPort.Port())

	return &TestContainer{
		testcontainer: container,
		migrator:      NewTestContainerMigrator(conn),
		conn:          conn,
	}, nil
}

func (c *TestContainer) Conn() string {
	return c.conn
}

func (c *TestContainer) Stop() error {
	err := c.testcontainer.Terminate(context.Background())
	if err != nil {
		return fmt.Errorf("failed to terminate Container, err: %w", err)
	}

	return nil
}

func (c *TestContainer) Migrate(source string) error {
	return c.migrator.Migrate(source)
}
