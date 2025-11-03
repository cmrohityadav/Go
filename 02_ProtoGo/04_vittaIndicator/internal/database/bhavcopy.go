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

func BhavcopyScheduler(cfg config.Config) {
	for {
		currentTime := time.Now()

		parsedTime, err := time.Parse("15:04", cfg.Bhavcopyurl.NSE.Time)

		if err != nil {
			log.Println("Error while parsing time format:", err)
			time.Sleep(1 * time.Minute)
			continue
		}

		bulkInsertNextTime := time.Date(
			currentTime.Year(),
			currentTime.Month(),
			currentTime.Day(),
			parsedTime.Hour(),
			parsedTime.Minute(),
			0, 0, currentTime.Location(),
		).Add(1 * time.Minute)

		if bulkInsertNextTime.Before(currentTime) {
			bulkInsertNextTime = bulkInsertNextTime.Add(24 * time.Hour)
		}

		sleepNextTimeForBulkInsert := bulkInsertNextTime.Sub(currentTime)

		log.Printf("[BhavcopyScheduler] Next BhavCopy insert scheduled at: %v (in %v)",
			bulkInsertNextTime.Format(time.RFC1123), sleepNextTimeForBulkInsert)

		time.Sleep(sleepNextTimeForBulkInsert)

		db := connectDb(cfg)
		if db == nil {
			log.Println("DB connection failed — skipping BhavCopyInsert.")
			continue
		}

		if err := ensureBhavcopyTable(db); err != nil {
			log.Fatal("Failed to create bhavcopy table:", err)
		}

		if err := truncateTable(db, "bhavcopy"); err != nil {
			log.Println(err)
		}

		log.Println("Bhavcopy table is ready!")

		log.Println("[BhavcopyScheduler] Running BhavCopyInsert...")

		sCurrentWD, err := os.Getwd()
		if err != nil {
			log.Println("Unable to get Working directory:", err)
			return
		}


		formatedUrl:=filepath.Base(cfg.Bhavcopyurl.NSE.URL);
		formatedUrl=strings.ReplaceAll(formatedUrl,".zip","");
		formatedUrl=utils.ParseFileNameWithDate(formatedUrl,&bulkInsertNextTime);
		sBhavcopyPath := filepath.Join(sCurrentWD, "storage", "download",formatedUrl)

		
		if err := BhavCopyInsert(db, sBhavcopyPath); err != nil {
			log.Println("BhavCopyInsert failed:", err)
		} else {
			log.Println("BhavCopyInsert completed successfully.")
		}

		db.Close()

		// Sleep 1 minute before next iteration to prevent tight loop
		time.Sleep(1 * time.Minute)
	}
}

func ensureBhavcopyTable(db *sql.DB) error {
	query := `
	CREATE TABLE IF NOT EXISTS bhavcopy (
		TradDt DATE,
		BizDt DATE,
		Sgmt TEXT,
		Src TEXT,
		FinInstrmTp TEXT,
		FinInstrmId TEXT,
		ISIN TEXT,
		TckrSymb TEXT,
		SctySrs TEXT,
		FinInstrmNm TEXT,
		OpnPric DOUBLE PRECISION,
		HghPric DOUBLE PRECISION,
		LwPric DOUBLE PRECISION,
		ClsPric DOUBLE PRECISION,
		LastPric DOUBLE PRECISION,
		PrvsClsgPric DOUBLE PRECISION,
		UndrlygPric DOUBLE PRECISION,
		SttlmPric DOUBLE PRECISION,
		OpnIntrst BIGINT,
		ChngInOpnIntrst BIGINT,
		TtlTradgVol BIGINT,
		TtlTrfVal DOUBLE PRECISION,
		TtlNbOfTxsExctd BIGINT,
		SsnId TEXT,
		NewBrdLotQty BIGINT,
		Rmks TEXT,
		Rsvd1 TEXT,
		Rsvd2 TEXT,
		Rsvd3 TEXT,
		Rsvd4 TEXT,
		XpryDt DATE,
		FininstrmActlXpryDt DATE,
		StrkPric DOUBLE PRECISION,
		OptnTp TEXT
	);`
	_, err := db.Exec(query)
	return err
}

func BhavCopyInsert(db *sql.DB, csvPath string) error {
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
		b := parseBhavcopyRow(rec)
		err := insertBhavcopy(db, b)
		if err != nil {
			log.Printf("row %d insert failed: %v", i+1, err)
		}
	}
	return nil
}




func parseBhavcopyRow(rec []string) types.Bhavcopy {
	toFloat := func(s string) *float64 {
		if s == "" {
			return nil
		}
		v, _ := strconv.ParseFloat(s, 64)
		return &v
	}
	toInt := func(s string) *int64 {
		if s == "" {
			return nil
		}
		v, _ := strconv.ParseInt(s, 10, 64)
		return &v
	}
	toDate := func(s string) *time.Time {
		if s == "" {
			return nil
		}
		t, err := time.Parse("02-01-2006", s)
		if err != nil {
			log.Printf("Invalid date: %v", s)
			return nil
		}
		return &t
	}

	return types.Bhavcopy{
		TradDt:              timeMustParse(rec[0]),
		BizDt:               timeMustParse(rec[1]),
		Sgmt:                rec[2],
		Src:                 rec[3],
		FinInstrmTp:         rec[4],
		FinInstrmId:         rec[5],
		ISIN:                rec[6],
		TckrSymb: cleanSymbol(rec[7]),
		SctySrs:             rec[8],
		XpryDt:              toDate(rec[9]),
		FininstrmActlXpryDt: toDate(rec[10]),
		StrkPric:            toFloat(rec[11]),
		OptnTp:              rec[12],
		FinInstrmNm:         rec[13],
		OpnPric:             toFloat(rec[14]),
		HghPric:             toFloat(rec[15]),
		LwPric:              toFloat(rec[16]),
		ClsPric:             toFloat(rec[17]),
		LastPric:            toFloat(rec[18]),
		PrvsClsgPric:        toFloat(rec[19]),
		UndrlygPric:         toFloat(rec[20]),
		SttlmPric:           toFloat(rec[21]),
		OpnIntrst:           toInt(rec[22]),
		ChngInOpnIntrst:     toInt(rec[23]),
		TtlTradgVol:         toInt(rec[24]),
		TtlTrfVal:           toFloat(rec[25]),
		TtlNbOfTxsExctd:     toInt(rec[26]),
		SsnId:               rec[27],
		NewBrdLotQty:        toInt(rec[28]),
		Rmks:                rec[29],
		Rsvd1:               rec[30],
		Rsvd2:               rec[31],
		Rsvd3:               rec[32],
		Rsvd4:               rec[33],
	}
}

func timeMustParse(s string) time.Time {
	t, err := time.Parse("2006-01-02", s)
	if err != nil {
		log.Fatalf("Invalid required date in CSV: %v", s)
	}
	return t
}

func insertBhavcopy(db *sql.DB, b types.Bhavcopy) error {
	stmt, err := db.Prepare(`
	INSERT INTO bhavcopy (
		TradDt, BizDt, Sgmt, Src, FinInstrmTp, FinInstrmId, ISIN, TckrSymb,
		SctySrs, FinInstrmNm, OpnPric, HghPric, LwPric, ClsPric, LastPric,
		PrvsClsgPric, UndrlygPric, SttlmPric, OpnIntrst, ChngInOpnIntrst,
		TtlTradgVol, TtlTrfVal, TtlNbOfTxsExctd, SsnId, NewBrdLotQty,
		Rmks, Rsvd1, Rsvd2, Rsvd3, Rsvd4,
		XpryDt, FininstrmActlXpryDt, StrkPric, OptnTp
	) VALUES (
		$1,$2,$3,$4,$5,$6,$7,$8,
		$9,$10,$11,$12,$13,$14,$15,
		$16,$17,$18,$19,$20,$21,$22,
		$23,$24,$25,$26,$27,$28,$29,
		$30,$31,$32,$33,$34
	)`)
	if err != nil {
		return fmt.Errorf("prepare statement failed: %v", err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		b.TradDt, b.BizDt, b.Sgmt, b.Src, b.FinInstrmTp, b.FinInstrmId, b.ISIN, b.TckrSymb,
		b.SctySrs, b.FinInstrmNm, b.OpnPric, b.HghPric, b.LwPric, b.ClsPric, b.LastPric,
		b.PrvsClsgPric, b.UndrlygPric, b.SttlmPric, b.OpnIntrst, b.ChngInOpnIntrst,
		b.TtlTradgVol, b.TtlTrfVal, b.TtlNbOfTxsExctd, b.SsnId, b.NewBrdLotQty,
		b.Rmks, b.Rsvd1, b.Rsvd2, b.Rsvd3, b.Rsvd4,
		b.XpryDt, b.FininstrmActlXpryDt, b.StrkPric, b.OptnTp,
	)
	if err != nil {
		return fmt.Errorf("insert failed: %v", err)
	}
	return nil
}
