package backup

import (
    "fmt"
    "path/filepath"
    "time"

    "github.com/Chinzzii/db-backup-utility/internal/database"
)

type IncrementalBackup struct {
    db       database.Database
    filePath string
}

func NewIncrementalBackup(db database.Database) *IncrementalBackup {
    return &IncrementalBackup{
        db:       db,
        filePath: filepath.Join("backups", fmt.Sprintf("incremental_backup_%s.bak", time.Now().Format("20060102150405"))),
    }
}

func (b *IncrementalBackup) Execute() error {
    data, err := b.db.Backup()
    if err != nil {
        return err
    }
    return saveBackup(data, b.filePath)
}

func (b *IncrementalBackup) Validate() error {
    // Implement validation logic here
    return nil
}

func (b *IncrementalBackup) GetFilePath() string {
    return b.filePath
}