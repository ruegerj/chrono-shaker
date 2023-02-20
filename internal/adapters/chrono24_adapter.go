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

const CHRONO_24_DOMAIN = "www.chrono24.com"

type Chrono24Adapter struct {
	filter *common.FilterOptions
}

func NewChrono24Adapter(filter *common.FilterOptions) Chrono24Adapter {
	return Chrono24Adapter{
		filter: filter,
	}
}

func (adapter Chrono24Adapter) CreateListingsUrl() string {
	path := fmt.Sprintf("/%s/ref-%s.htm?pageSize=120&resultview=block",
		strings.ToLower(adapter.filter.Brand),
		strings.ToLower(adapter.filter.RefNo))

	return createChrono24Url(path)
}

func (adapter Chrono24Adapter) Parse(g *geziyor.Geziyor, r *client.Response) {
	r.HTMLDoc.Find("div.article-item-container").Each(func(_ int, s *goquery.Selection) {
		listingAnchor := s.Find("a")
		brand := listingAnchor.AttrOr("data-manufacturer", common.NO_VALUE)
		priceRaw := s.Find("span.currency").Parent().Text()
		listingUrl := listingAnchor.AttrOr("href", common.NO_VALUE)

		priceRaw = strings.TrimPrefix(priceRaw, "\n$")
		priceRaw = strings.ReplaceAll(priceRaw, ",", ".")
		priceValue, _ := decimal.NewFromString(priceRaw)
		price := common.NewPrice(priceValue, "$")

		if brand == common.NO_VALUE || price.Value.IsZero() {
			return
		}

		listing := *common.NewWatchListing(brand,
			adapter.filter.RefNo,
			price,
			common.CHRONO_24,
			createChrono24Url(listingUrl),
		)

		g.Exports <- listing
	})

	if href, ok := r.HTMLDoc.Find("a.paging-next").Attr("href"); ok {
		g.Get(href, adapter.Parse)
	}
}

func createChrono24Url(path string) string {
	return fmt.Sprintf("https://%s%s", CHRONO_24_DOMAIN, path)
}
