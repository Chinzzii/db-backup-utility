package backup

import (
	"fmt"
	"os"

	"github.com/Chinzzii/db-backup-utility/internal/database"
)

// Backup interface defines the methods required for backup operations
type Backup interface {
	Execute() error
	Validate() error
	GetFilePath() string
}

// PerformBackup performs the backup based on the backup type
func PerformBackup(db database.Database, backupType string, compressionEnabled bool) error {
	var backup Backup

	switch backupType {
	case "full":
		backup = NewFullBackup(db)
	case "incremental":
		backup = NewIncrementalBackup(db)
	case "differential":
		backup = NewDifferentialBackup(db)
	default:
		return fmt.Errorf("unsupported backup type: %s", backupType)
	}

	if err := backup.Validate(); err != nil {
		return err
	}

	if err := backup.Execute(); err != nil {
		return err
	}

	if compressionEnabled {
		data, err := os.ReadFile(backup.GetFilePath())
		if err != nil {
			return err
		}

		compressedData, err := compressData(data)
		if err != nil {
			return err
		}

		return saveBackup(compressedData, backup.GetFilePath())
	}

	return nil
}

func compressData(data []byte) ([]byte, error) {
	// Implement compression logic here
	return data, nil
}

func saveBackup(data []byte, filePath string) error {
	return os.WriteFile(filePath, data, 0644)
}
