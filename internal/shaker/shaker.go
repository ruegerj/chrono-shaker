package shaker

import (
	"github.com/geziyor/geziyor"
	"github.com/ruegerj/chrono-shaker/internal/adapters"
	"github.com/ruegerj/chrono-shaker/internal/common"
)

type Shaker struct {
	adapter common.PlatformAdapter
	filter  *common.FilterOptions
	scraper *geziyor.Geziyor
}

func NewShaker(platform common.Platform, filter *common.FilterOptions) (*Shaker, error) {
	adapter, err := adapters.Factory(platform, filter)

	if err != nil {
		return nil, err
	}

	defaultOptions := createDefaultOptions()
	targetUrl := adapter.CreateListingsUrl()

	// Customize options based on adapter
	defaultOptions.ParseFunc = adapter.Parse
	defaultOptions.StartURLs = append(defaultOptions.StartURLs, targetUrl)

	scraper := geziyor.NewGeziyor(defaultOptions)

	return &Shaker{adapter: adapter, filter: filter, scraper: scraper}, nil
}

func (shaker *Shaker) ShakeListings() []common.WatchListing {
	exporter := common.NewMemoryExport[common.WatchListing]()

	shaker.scraper.Opt.Exporters = append(shaker.scraper.Opt.Exporters, exporter)

	shaker.scraper.Start()

	return exporter.Results
}

func createDefaultOptions() *geziyor.Options {
	return &geziyor.Options{}
}
