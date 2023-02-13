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

type Chrono24Adapter struct {
	filter *common.FilterOptions
}

func NewChrono24Adapter(filter *common.FilterOptions) Chrono24Adapter {
	return Chrono24Adapter{
		filter: filter,
	}
}

func (adapter Chrono24Adapter) CreateListingsUrl() string {
	return fmt.Sprintf("https://www.chrono24.com/%s/ref-%s.htm",
		strings.ToLower(adapter.filter.Brand),
		strings.ToLower(adapter.filter.RefNo))
}

func (adapter Chrono24Adapter) Parse(g *geziyor.Geziyor, r *client.Response) {
	r.HTMLDoc.Find("div.article-item-container").Each(func(_ int, s *goquery.Selection) {
		brand := s.Find("a").AttrOr("data-manufacturer", "n/a")
		priceRaw := s.Find("span.currency").Parent().Text()

		priceRaw = strings.TrimPrefix(priceRaw, "\n$")
		priceRaw = strings.ReplaceAll(priceRaw, ",", ".")
		price, _ := decimal.NewFromString(priceRaw)

		listing := *common.NewWatchListing(brand, adapter.filter.RefNo, price, common.CHRONO_24)

		g.Exports <- listing
	})
}
