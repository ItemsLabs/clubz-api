package config

import "fmt"

func DatabaseURL() string {
	return fmt.Sprintf("host=%s port=%d dbname=%s user=%s password=%s",
		cfg.DatabaseHost,
		cfg.DatabasePort,
		cfg.DatabaseName,
		cfg.DatabaseUser,
		cfg.DatabasePassword,
	)
}

func DatabaseMaxIdleConnections() int {
	return cfg.DatabaseMaxIdleConnections
}

func DatabaseMaxOpenConnections() int {
	return cfg.DatabaseMaxOpenConnections
}
