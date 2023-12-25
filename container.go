package integro

import "net"

type Container interface {
	// Conn returns connection string for the container as DSN
	Conn() string
	// Ip returns container ip
	Ip() net.IP
	// Port returns container port
	Port() string
	// Stop stops the container
	Stop() error
	// Migrate applies migrations from source
	Migrate(source string) error
}
