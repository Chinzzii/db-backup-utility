package database

import (
	"fmt"
)

// Config struct for database configuration
type Config struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
}

// Database interface defines the methods required for database operations
type Database interface {
	Connect(config Config) error
	Backup() ([]byte, error)
	Restore(data []byte) error
	Close() error
}

// NewDatabase creates a new database connection based on the dbType
func NewDatabase(dbType string, config Config) (Database, error) {
	switch dbType {
	case "mysql":
		return NewMySQLDatabase(config)
	case "postgres":
		return NewPostgresDatabase(config)
	case "mongodb":
		return NewMongoDatabase(config)
	case "sqlite":
		return NewSQLiteDatabase(config)
	default:
		return nil, fmt.Errorf("unsupported database type: %s", dbType)
	}
}
