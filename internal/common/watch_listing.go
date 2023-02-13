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
	Url      string
	Date     time.Time
}

func NewWatchListing(brand string, refNo string, price decimal.Decimal, platform Platform, url string) *WatchListing {
	return &WatchListing{
		Brand:    brand,
		RefNo:    refNo,
		Price:    price,
		Platform: platform,
		Url:      url,
		Date:     time.Now(),
	}
}

func (listing *WatchListing) String() string {
	return fmt.Sprintf("[%s] %s %s: %d (%s)",
		listing.Platform,
		listing.Brand,
		listing.RefNo,
		listing.Price.CoefficientInt64(),
		listing.Url,
	)
}
