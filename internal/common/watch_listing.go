package common

import (
	"fmt"
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

func (listing *WatchListing) String() string {
	return fmt.Sprintf("%s %s: %d (%s / %v)",
		listing.Brand,
		listing.RefNo,
		listing.Price.CoefficientInt64(),
		listing.Platform,
		listing.Date.Format(time.RFC822),
	)
}
