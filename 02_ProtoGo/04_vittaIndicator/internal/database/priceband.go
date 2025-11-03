package database

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
	"vittaindicator/internal/config"
	"vittaindicator/internal/types"
	"vittaindicator/internal/utils"
)

func PriceBandScheduler(cfg config.Config) {
	for {
		currentTime := time.Now()
		parsedTime, err := time.Parse("15:04", cfg.Priceband.NSE.Time)
		if err != nil {
			log.Println("Error parsing scheduler time:", err)
			time.Sleep(1 * time.Minute)
			continue
		}

		nextTime := time.Date(
			currentTime.Year(),
			currentTime.Month(),
			currentTime.Day(),
			parsedTime.Hour(),
			parsedTime.Minute(),
			0, 0, currentTime.Location(),
		).Add(1 * time.Minute)

		if nextTime.Before(currentTime) {
			nextTime = nextTime.Add(24 * time.Hour)
		}

		sleepDuration := nextTime.Sub(currentTime)
		log.Printf("[PriceBandScheduler] Next insert scheduled at %v (in %v)", nextTime, sleepDuration)
		time.Sleep(sleepDuration)

		db := connectDb(cfg)
		if err := ensurePriceBandTable(db); err != nil {
			log.Fatal("Failed to create table:", err)
		}

		if err := truncateTable(db, "priceband"); err != nil {
			log.Println(err)
		}

		sCurrentWD, _ := os.Getwd();
		
		formatedUrl:=filepath.Base(cfg.Priceband.NSE.URL);
		formatedUrl=utils.ParseFileNameWithDate(formatedUrl,&nextTime);
		sBhavcopyPath := filepath.Join(sCurrentWD, "storage", "download",formatedUrl)
		


		log.Println("[PriceBandScheduler] Running PriceBandInsert...")


		if err := PriceBandInsert(db, sBhavcopyPath); err != nil {
			log.Println("PriceBandInsert failed:", err)
		} else {
			log.Println("PriceBandInsert completed successfully.")
		}

		db.Close()
		time.Sleep(1 * time.Minute)
	}
}

func ensurePriceBandTable(db *sql.DB) error {
	query := `
	CREATE TABLE IF NOT EXISTS priceband (
		Symbol TEXT,
		Series TEXT,
		Name TEXT,
		Band INT,
		Remarks TEXT
	);`
	_, err := db.Exec(query)
	return err
}

func PriceBandInsert(db *sql.DB, csvPath string) error {
	file, err := os.Open(csvPath)
	if err != nil {
		return fmt.Errorf("failed to open CSV: %v", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.TrimLeadingSpace = true
	records, err := reader.ReadAll()
	if err != nil {
		return fmt.Errorf("failed to read CSV: %v", err)
	}

	for i, rec := range records[1:] {
		b := parsePriceBandRow(rec)
		if err := insertPriceBand(db, b); err != nil {
			log.Printf("Row %d insert failed: %v", i+1, err)
		}
	}
	return nil
}

func cleanSymbol(s string) string {
    s = strings.TrimSpace(s)            // Remove leading/trailing spaces
    s = strings.Trim(s, "\"")           // Remove double quotes
    s = strings.ToUpper(s)              // Uppercase for uniformity
    return s
}

func parsePriceBandRow(rec []string) types.PriceBand {
	band := 0
	if rec[3] != "" && rec[3] != "-" && rec[3] != "No Band" {
		if val, err := strconv.Atoi(rec[3]); err == nil {
			band = val
		}
	}
	return types.PriceBand{
		Symbol:  cleanSymbol(rec[0]),
        Series:  strings.TrimSpace(rec[1]),
        Name:    strings.TrimSpace(rec[2]),
        Band:    band,
        Remarks: strings.TrimSpace(rec[4]),
	}
}

func insertPriceBand(db *sql.DB, b types.PriceBand) error {
	stmt, err := db.Prepare(`
		INSERT INTO priceband (Symbol, Series, Name, Band, Remarks)
		VALUES ($1,$2,$3,$4,$5)
	`)
	if err != nil {
		return fmt.Errorf("prepare statement failed: %v", err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(b.Symbol, b.Series, b.Name, b.Band, b.Remarks)
	if err != nil {
		return fmt.Errorf("insert failed: %v", err)
	}
	return nil
}
