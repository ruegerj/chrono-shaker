package common

import (
	"fmt"
	"time"
)

type WatchListing struct {
	Brand    string
	RefNo    string
	Price    *Price
	Platform Platform
	Url      string
	Date     time.Time
}

func NewWatchListing(brand string, refNo string, price *Price, platform Platform, url string) *WatchListing {
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
	return fmt.Sprintf("[%s] %s %s: %s (%s)",
		listing.Platform,
		listing.Brand,
		listing.RefNo,
		listing.Price,
		listing.Url,
	)
}
