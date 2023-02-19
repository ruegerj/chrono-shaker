package adapters

import (
	"errors"

	"github.com/ruegerj/chrono-shaker/internal/common"
)

func Factory(platform common.Platform, filter *common.FilterOptions) (common.PlatformAdapter, error) {
	switch platform {
	case common.CHRONO_24:
		return NewChrono24Adapter(filter), nil
	}

	return nil, errors.New("Invalid platform specified")
}
