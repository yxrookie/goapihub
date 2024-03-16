// Package migrate: cope database migration
package migrate

import (
	"database/sql"

	"gorm.io/gorm"
)

// migrationFunc: define up and down callback methods' type
type migrationFunc func(gorm.Migrator, *sql.DB)

// migrationFiles: all migration file array
var migrationFiles []MigrationFile

// MigrationFile: represent single migration file
type MigrationFile struct {
	Up migrationFunc
	Down migrationFunc
	FileName string
} 
// Add: add a migration file, all migration files need to register using this method
func Add(name string, up migrationFunc, down migrationFunc) {
	migrationFiles = append(migrationFiles, MigrationFile{
		FileName: name,
		Up: up,
		Down: down,
	})
}
