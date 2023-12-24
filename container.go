package integro

type Container interface {
	Conn() string
	Stop() error
	Migrate(source string) error
}
