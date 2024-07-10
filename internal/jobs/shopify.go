package jobs

import (
	"farmz-affiliate-cron/internal/api"
	"farmz-affiliate-cron/internal/config"
	"gorm.io/gorm"
)

func RunShopifyJob(db *gorm.DB, cfg *config.Config) error {
	client := api.NewShopifyClient(cfg.ShopifyAPIToken)
	return client.DownloadDailyTransaction(db, cfg.Date)
}
