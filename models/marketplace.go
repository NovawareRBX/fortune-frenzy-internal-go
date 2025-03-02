package models

import "time"

type Item struct {
	ID            int       `json:"id"`
	AssetID       int64     `json:"asset_id"`
	Name          string    `json:"name"`
	Creator       string    `json:"creator"`
	Description   string    `json:"description"`
	AveragePrice  int64     `json:"average_price"`
	TotalUnboxed  int64     `json:"total_unboxed"`
	MaximumCopies int64     `json:"maximum_copies"`
	Value         int64     `json:"value"`
	CreatedAt     time.Time `json:"created_at"`
	Color         string    `json:"color"`
}
