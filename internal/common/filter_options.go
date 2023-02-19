package common

type FilterOptions struct {
	Brand     string
	RefNo     string
	Platforms []Platform
}

func NewFilterOptions(brand string, refNo string, platforms []Platform) *FilterOptions {
	return &FilterOptions{
		Brand:     brand,
		RefNo:     refNo,
		Platforms: platforms,
	}
}
