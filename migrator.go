package integro

type Migrator interface {
	Migrate(source string) error
}
