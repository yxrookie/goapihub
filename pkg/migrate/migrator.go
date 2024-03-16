package migrate

import (
	"goapihub/pkg/database"

	"gorm.io/gorm"
)

// Migrator: data migration operations class
type Migrator struct {
	Folder string
	DB *gorm.DB
	Migrator gorm.Migrator
}

//Migration: represent a row of data of the migrations table 
type Migration struct {
	ID uint64 `gorm:"primarykey;autoIncrement;"`
	Migration string `gorm:"type:varchar(255);not null;unique;"`
	Batch int
}

// NewMigrator: create Migrator instance, implement the migration operation 
func NewMigrator() *Migrator {

	// initialize field
	migrator := &Migrator {
		Folder: "database/migrations",
		DB: database.DB,
		Migrator: database.DB.Migrator(),
	}
	// migrations: creation migrations if it not exist
	migrator.createMigrationsTanle()
	return migrator
}

// create migrations table
func(migrator *Migrator) createMigrationsTanle() {
	migration := Migration{}

	// create if not exist 
	if !migrator.Migrator.HasTable(&migration) {
		migrator.Migrator.CreateTable(&migration)
	}
}