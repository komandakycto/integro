package wrappers

import "strings"

const (
	pgPort    = "5432"
	mysqlPort = "3306"
	mongoPort = "27017"
)

func defaultPort(image string) string {
	if strings.Contains(image, "postgres") {
		return pgPort
	}
	if strings.Contains(image, "mysql") {
		return mysqlPort
	}
	if strings.Contains(image, "mongo") {
		return mongoPort
	}

	return ""
}

func defaultConnection(image string) string {
	if strings.Contains(image, "postgres") {
		return "postgres://postgres:example@%s:%s/postgres?sslmode=disable"
	}
	if strings.Contains(image, "mysql") {
		return "mysql://root:root@tcp(%s:%s)/public"
	}
	if strings.Contains(image, "mongo") {
		return "mongodb://root:root@%s:%s/public?authSource=admin"
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
			"MYSQL_ROOT_PASSWORD": "root",
			"MYSQL_DATABASE":      "public",
		}
	}
	if strings.Contains(image, "mongo") {
		return map[string]string{
			"MONGO_INITDB_DATABASE":      "public",
			"MONGO_INITDB_ROOT_USERNAME": "root",
			"MONGO_INITDB_ROOT_PASSWORD": "root",
		}
	}

	return nil
}
