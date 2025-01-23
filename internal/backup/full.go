package backup

import (
	"fmt"
	"path/filepath"
	"time"

	"github.com/Chinzzii/db-backup-utility/internal/database"
)

type FullBackup struct {
	db       database.Database
	filePath string
}

func NewFullBackup(db database.Database) *FullBackup {
	return &FullBackup{
		db:       db,
		filePath: filepath.Join("backups", fmt.Sprintf("full_backup_%s.bak", time.Now().Format("20060102150405"))),
	}
}

func (b *FullBackup) Execute() error {
	data, err := b.db.Backup()
	if err != nil {
		return err
	}
	return saveBackup(data, b.filePath)
}

func (b *FullBackup) Validate() error {
	// Implement validation logic here
	return nil
}

func (b *FullBackup) GetFilePath() string {
	return b.filePath
}
