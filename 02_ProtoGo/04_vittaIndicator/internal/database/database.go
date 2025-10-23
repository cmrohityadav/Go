package database

import (
	"database/sql"
	"fmt"
	"log"
	"vittaindicator/internal/config"

	_ "github.com/lib/pq"
)

func connectDb(cfg config.Config) *sql.DB {
	connStr := fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=%s",
		cfg.Database.User, cfg.Database.Password, cfg.Database.Host, cfg.Database.Port, cfg.Database.DBName, cfg.Database.SSLMode,
	)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Error opening DB:", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal("Cannot connect to DB:", err)
	}

	return db
}

func truncateTable(db *sql.DB, tableName string) error {
    _, err := db.Exec(fmt.Sprintf("TRUNCATE TABLE %s", tableName))
    if err != nil {
        return fmt.Errorf("failed to truncate table %s: %v", tableName, err)
    }
    return nil
}


