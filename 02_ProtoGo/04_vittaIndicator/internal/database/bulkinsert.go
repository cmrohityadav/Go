package database

import "vittaindicator/internal/config"

func ScheduledBulkInsert(cfg config.Config) {
	go BhavcopyScheduler(cfg)
	go PriceBandScheduler(cfg)
}