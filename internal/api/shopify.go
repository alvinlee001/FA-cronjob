package api

import (
	"encoding/json"
	"farmz-affiliate-cron/db/models"
	"fmt"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"log"
	"net/http"
	"time"
)

// ShopifyOrdersResponse represents the response from Shopify orders API
type ShopifyOrdersResponse struct {
	Orders []models.ShopifyOrder `json:"orders"`
}

type ShopifyClient struct {
	API_TOKEN string
}

func NewShopifyClient(token string) *ShopifyClient {
	return &ShopifyClient{API_TOKEN: token}
}

func (c *ShopifyClient) DownloadDailyTransaction(db *gorm.DB, date time.Time) error {
	fmt.Println("Downloading daily transactions for", date)

	dayBefore := date.AddDate(0, 0, -1)
	req, err := http.NewRequest("GET", fmt.Sprintf("https://farmz-fresh-to-go-my.myshopify.com/admin/api/2023-04/orders.json"), nil)
	if err != nil {
		return err
	}

	// Set headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Shopify-Access-Token", c.API_TOKEN)
	// Set query parameters
	q := req.URL.Query()
	q.Add("created_at_min", dayBefore.Format("2006-01-02"))
	q.Add("created_at_max", date.Format("2006-01-02"))
	//q.Add("created_at_min", "2023-02-01")
	//q.Add("created_at_max", "2024-10-02")
	req.URL.RawQuery = q.Encode()

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		log.Fatal("Error sending shopify order request", err)
		return err
	}
	defer resp.Body.Close()

	var ordersResponse ShopifyOrdersResponse
	if err := json.NewDecoder(resp.Body).Decode(&ordersResponse); err != nil {
		log.Fatalf("Error decoding response: %v", err)
	}
	if err != nil {
		fmt.Println(err)
		log.Fatal("Cannot read request body", err)
		return err
	}
	defer resp.Body.Close()

	b, err := json.Marshal(ordersResponse)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(b))
	result := db.Clauses(clause.OnConflict{
		DoNothing: true,
	}).Create(&ordersResponse.Orders)
	fmt.Println(result)
	return nil
}
