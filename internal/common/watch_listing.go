package common

import (
	"time"

	"github.com/shopspring/decimal"
)

type WatchListing struct {
	Brand    string
	RefNo    string
	Price    decimal.Decimal
	Platform Platform
	Date     time.Time
}

func NewWatchListing(brand string, refNo string, price decimal.Decimal, platform Platform) *WatchListing {
	return &WatchListing{
		Brand:    brand,
		RefNo:    refNo,
		Price:    price,
		Platform: platform,
		Date:     time.Now(),
	}
}
