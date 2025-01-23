package backup

import (
    "fmt"
    "path/filepath"
    "time"

    "github.com/Chinzzii/db-backup-utility/internal/database"
)

type DifferentialBackup struct {
    db       database.Database
    filePath string
}

func NewDifferentialBackup(db database.Database) *DifferentialBackup {
    return &DifferentialBackup{
        db:       db,
        filePath: filepath.Join("backups", fmt.Sprintf("differential_backup_%s.bak", time.Now().Format("20060102150405"))),
    }
}

func (b *DifferentialBackup) Execute() error {
    data, err := b.db.Backup()
    if err != nil {
        return err
    }
    return saveBackup(data, b.filePath)
}

func (b *DifferentialBackup) Validate() error {
    // Implement validation logic here
    return nil
}

func (b *DifferentialBackup) GetFilePath() string {
    return b.filePath
}