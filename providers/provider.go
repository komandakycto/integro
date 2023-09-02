package providers

// Provider is an interface for database providers.
type Provider interface {
	// Create creates a new database container.
	Create(image string) error
	// Terminate terminates the database container.
	Terminate() error
	// Conn returns a connection string to the database.
	Conn() string
	// Ip returns an ip address of the database container.
	Ip() string
	// Port returns a port of the database container.
	Port() int
	// Migrate migrates the database.
	Migrate(source string) error
}
