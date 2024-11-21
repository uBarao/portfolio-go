package config

import "os"

func GetDatabaseConnectionString() string {
	return os.Getenv("DB_CONNECTION")
}
