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


// getMigrationFile 通过迁移文件的名称来获取到 MigrationFile 对象
func getMigrationFile(name string) MigrationFile {
    for _, mfile := range migrationFiles {
        if name == mfile.FileName {
            return mfile
        }
    }
    return MigrationFile{}
}

// isNotMigrated 判断迁移是否已执行
func (mfile MigrationFile) isNotMigrated(migrations []Migration) bool {
    for _, migration := range migrations {
        if migration.Migration == mfile.FileName {
            return false
        }
    }
    return true
}
