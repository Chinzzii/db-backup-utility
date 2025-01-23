package main

import (

    "log"
    "os"

	"github.com/spf13/cobra"

    "github.com/Chinzzii/db-backup-utility/internal/backup"
    "github.com/Chinzzii/db-backup-utility/internal/database"
    "github.com/Chinzzii/db-backup-utility/internal/storage"
    "github.com/Chinzzii/db-backup-utility/internal/logger"
)

func main() {
	var dbType, host, user, password, dbName, backupType, storageType string
    var port int
    var compressionEnabled, notifySlack bool

    rootCmd := &cobra.Command{
        Use:   "db-backup-utility",
        Short: "A CLI tool for database backup",
        Run: func(cmd *cobra.Command, args []string) {
            // Initialize logger
            logger := logger.NewLogger()

            // Validate inputs
            if dbType == "" {
                logger.Fatal("Database type is required")
            }

            // Database connection
            db, err := database.NewDatabase(dbType, host, port, user, password, dbName)
            if err != nil {
                logger.Fatal(err)
            }

            // Perform backup
            err = backup.PerformBackup(db, backupType, compressionEnabled)
            if err != nil {
                logger.Fatal(err)
            }

            // Store backup
            err = storage.StoreBackup(storageType)
            if err != nil {
                logger.Fatal(err)
            }

            // Send notification if enabled
            if notifySlack {
                err = notify.SendSlackNotification()
                if err != nil {
                    logger.Fatal(err)
                }
            }

            logger.Info("Backup completed successfully")
        },
    }

    rootCmd.Flags().StringVarP(&dbType, "db-type", "d", "", "Database type (mysql/postgres/mongodb/sqlite)")
    rootCmd.Flags().StringVarP(&host, "host", "H", "localhost", "Database host")
    rootCmd.Flags().IntVarP(&port, "port", "P", 0, "Database port")
    rootCmd.Flags().StringVarP(&user, "user", "u", "", "Database user")
    rootCmd.Flags().StringVarP(&password, "password", "p", "", "Database password")
    rootCmd.Flags().StringVarP(&dbName, "dbname", "n", "", "Database name")
    rootCmd.Flags().StringVarP(&backupType, "backup-type", "b", "full", "Backup type (full/incremental/differential)")
    rootCmd.Flags().StringVarP(&storageType, "storage", "s", "local", "Storage type (local/s3/gcs/azure)")
    rootCmd.Flags().BoolVarP(&compressionEnabled, "compress", "c", true, "Enable compression")
    rootCmd.Flags().BoolVarP(&notifySlack, "notify-slack", "N", false, "Send Slack notification")

    if err := rootCmd.Execute(); err != nil {
        log.Fatal(err)
        os.Exit(1)
    }
}