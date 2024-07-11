package main

import (
	"farmz-affiliate-cron/db/models"
	"farmz-affiliate-cron/internal/config"
	"farmz-affiliate-cron/internal/jobs"
	"flag"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"time"
)

func main() {
	// Define the date flag
	datePtr := flag.String("date", "", "optional date parameter in YYYY-MM-DD format")

	// Parse the command-line flags
	flag.Parse()
	var date time.Time
	var err error

	// Check if the date parameter was provided
	if *datePtr != "" {
		// Parse the date
		date, err = time.Parse("2006-01-02", *datePtr)
		if err != nil {
			fmt.Println("Invalid date format. Please use YYYY-MM-DD.")
			return
		}
	} else {
		// If no date is provided, use the current date
		date = time.Now()
	}

	cfg := config.LoadConfig()
	cfg.Date = date

	db, err := gorm.Open(mysql.Open(cfg.MysqlDSN), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}
	err = db.AutoMigrate(&models.GhlTokens{})
	if err != nil {
		log.Fatalf("Error runnning migration: %v", err)
	}

	err = jobs.RunGhlJob(db, cfg)
	if err != nil {
		log.Fatalf("GHL job failed: %v", err)
	}

	log.Println("GHL job completed successfully")
}
