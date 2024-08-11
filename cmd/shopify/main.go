package main

import (
	"farmz-affiliate-cron/internal/jobs"
	"flag"
)

func main() {
	// Define the date flag
	datePtr := flag.String("date", "", "optional date parameter in YYYY-MM-DD format")

	// Parse the command-line flags
	flag.Parse()

	jobs.RunShopifyJob(datePtr)
}
