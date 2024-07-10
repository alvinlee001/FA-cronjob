package models

type ShopifyOrder struct {
	ID              int64  `json:"id" gorm:"primaryKey"`
	CancelReason    string `json:"cancel_reason"`
	CreatedAt       string `json:"created_at"`
	Email           string `json:"email"`
	OrderNumber     int    `json:"order_number"`
	ProcessedAt     string `json:"processed_at"`
	SubtotalPrice   string `json:"subtotal_price"`
	TotalPrice      string `json:"total_price"`
	TotalTax        string `json:"total_tax"`
	FinancialStatus string `json:"financial_status"`
	//LineItems       []ShopifyOrderLineItem `json:"line_items" gorm:"foreignKey:ID,references:ID,constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	LineItems []ShopifyOrderLineItem `json:"line_items" gorm:"foreignKey:ID"`
}

type ShopifyOrderLineItem struct {
	ID            int64  `json:"id" gorm:"primaryKey"`
	Price         string `json:"price" sql:"gorm:decimal(20,2)"`
	TotalDiscount string `json:"total_discount" sql:"gorm:decimal(20,2)"`
	Sku           string `json:"sku"`
	VariantId     int64  `json:"variant_id"`
	OrderID       int64
}
