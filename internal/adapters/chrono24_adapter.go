package adapters

import (
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/geziyor/geziyor"
	"github.com/geziyor/geziyor/client"
	"github.com/ruegerj/chrono-shaker/internal/common"
	"github.com/shopspring/decimal"
)

type Chono24Adapter struct {
	filter *common.FilterOptions
}

func NewChrono24Adapter(filter *common.FilterOptions) Chono24Adapter {
	return Chono24Adapter{
		filter: filter,
	}
}

func (adapter Chono24Adapter) CreateListingsUrl() string {
	return fmt.Sprintf("https://www.chrono24.com/%s/%s.htm",
		strings.ToLower(adapter.filter.Brand),
		strings.ToLower(adapter.filter.RefNo))
}

func (adapter Chono24Adapter) Parse(g *geziyor.Geziyor, r *client.Response) {
	r.HTMLDoc.Find("div.article-item-container").Each(func(_ int, s *goquery.Selection) {
		brand := s.Find("a").AttrOr("data-manufacturer", "n/a")
		priceRaw := s.Find("span.currency").Parent().Text()
		price, _ := decimal.NewFromString(priceRaw)

		listing := *common.NewWatchListing(brand, adapter.filter.RefNo, price, common.CHRONO_24)

		g.Exports <- listing
	})
}
