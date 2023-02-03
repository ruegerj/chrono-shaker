package common

type FilterOptions struct {
	Brand string
	RefNo string
}

func NewFilterOptions(brand string, refNo string) *FilterOptions {
	return &FilterOptions{
		Brand: brand,
		RefNo: refNo,
	}
}
