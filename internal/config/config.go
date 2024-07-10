package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"time"
)

type Config struct {
	ShopifyAPIToken string
	GHLAPIKey       string
	TiktokAPIKey    string
	MysqlDSN        string
	// Add other configuration fields as needed
	Date time.Time
}

func LoadConfig() *Config {

	// Load .env file if it exists
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	config := &Config{
		ShopifyAPIToken: getEnv("SHOPIFY_API_TOKEN", ""),
		GHLAPIKey:       getEnv("GHL_API_KEY", ""),
		TiktokAPIKey:    getEnv("TIKTOK_API_KEY", ""),
		MysqlDSN:        getEnv("MYSQL_DSN", ""),
	}

	//log.Printf("Loaded configuration %s", config)

	return config
}

// Helper function to get an environment variable or return a default value
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
