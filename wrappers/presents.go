package wrappers

import "strings"

const (
	pgPort    = "5432"
	mysqlPort = "3306"
)

func defaultPort(image string) string {
	if strings.Contains(image, "postgres") {
		return pgPort
	}
	if strings.Contains(image, "mysql") {
		return mysqlPort
	}

	return ""
}

func defaultConnection(image string) string {
	if strings.Contains(image, "postgres") {
		return "postgres://postgres:example@%s:%s/postgres?sslmode=disable"
	}
	if strings.Contains(image, "mysql") {
		return "mysql://root:example@tcp(%s:%s)/mysql"
	}

	return ""
}

func defaultEnv(image string) map[string]string {
	if strings.Contains(image, "postgres") {
		return map[string]string{
			"POSTGRES_HOST_AUTH_METHOD": "trust",
		}
	}
	if strings.Contains(image, "mysql") {
		return map[string]string{
			"MYSQL_ROOT_PASSWORD": "example",
		}
	}

	return nil
}
