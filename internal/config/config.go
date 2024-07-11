package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"time"
)

type ClientCredentials struct {
	ClientID     string
	ClientSecret string
}

type GhlProperties struct {
	CompanyCredentials  ClientCredentials
	LocationCredentials ClientCredentials
}

type Config struct {
	ShopifyAPIToken string
	GHL             GhlProperties
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
		GHL: GhlProperties{
			CompanyCredentials: ClientCredentials{
				ClientID:     getEnv("GHL_COMPANY_CLIENT_ID", ""),
				ClientSecret: getEnv("GHL_COMPANY_CLIENT_SECRET", ""),
			},
			LocationCredentials: ClientCredentials{
				ClientID:     getEnv("GHL_LOCATION_CLIENT_ID", ""),
				ClientSecret: getEnv("GHL_LOCATION_CLIENT_SECRET", ""),
			},
		},
		TiktokAPIKey: getEnv("TIKTOK_API_KEY", ""),
		MysqlDSN:     getEnv("MYSQL_DSN", ""),
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
