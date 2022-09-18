package config

import (
	"goravel/facades"
)

func init() {
	config := facades.Config
	config.Add("database", map[string]interface{}{
		//Default database connection name, only support Mysql now.
		"default": config.Env("DB_CONNECTION", "mysql"),

		//Database connections
		"connections": map[string]interface{}{
			"mysql": map[string]interface{}{
				"host":     config.Env("DB_HOST", "mysql-meli"),
				"port":     config.Env("DB_PORT", "3306"),
				"database": config.Env("DB_DATABASE", "meli"),
				"username": config.Env("DB_USERNAME", "root"),
				"password": config.Env("DB_PASSWORD", "root"),
				"charset":  "utf8mb4",
			},
		},
	})
}
