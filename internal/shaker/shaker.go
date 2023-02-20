package shaker

import (
	"sync"

	"github.com/geziyor/geziyor"
	"github.com/ruegerj/chrono-shaker/internal/adapters"
	"github.com/ruegerj/chrono-shaker/internal/common"
)

type Shaker struct {
	filter *common.FilterOptions
}

func NewShaker(filter *common.FilterOptions) *Shaker {
	return &Shaker{filter: filter}
}

func (shaker *Shaker) ShakeListings() []common.WatchListing {
	var waitGroup sync.WaitGroup
	platformCount := len(shaker.filter.Platforms)
	waitGroup.Add(platformCount)

	outputChan := make(chan common.WatchListing)

	for _, platform := range shaker.filter.Platforms {
		go shaker.scrapePlatform(platform, outputChan, &waitGroup)
	}

	eoiCounter := 0
	var emptyListing common.WatchListing
	listings := make([]common.WatchListing, 0)

	for listing := range outputChan {
		if listing == emptyListing {
			eoiCounter++

			if eoiCounter < platformCount {
				continue
			}

			break
		}

		listings = append(listings, listing)
	}

	waitGroup.Wait()
	close(outputChan)

	return listings
}

func (shaker *Shaker) scrapePlatform(platform common.Platform, output chan common.WatchListing, wg *sync.WaitGroup) {
	defer wg.Done()
	adapter, err := adapters.Factory(platform, shaker.filter)

	if err != nil {
		panic(err)
	}

	scraper := buildScraper(adapter)

	exporter := common.NewChannelPipeExporter(output)
	scraper.Opt.Exporters = append(scraper.Opt.Exporters, exporter)

	scraper.Start()
}

func buildScraper(adapter common.PlatformAdapter) *geziyor.Geziyor {
	defaultOptions := &geziyor.Options{
		LogDisabled: true,
	}
	targetUrl := adapter.CreateListingsUrl()

	// Customize options based on adapter
	defaultOptions.ParseFunc = adapter.Parse
	defaultOptions.StartURLs = append(defaultOptions.StartURLs, targetUrl)

	return geziyor.NewGeziyor(defaultOptions)
}
