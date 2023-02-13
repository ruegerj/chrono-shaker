package common

type Platform uint

const (
	CHRONO_24 Platform = iota
)

func (platform Platform) String() string {
	switch platform {
	case CHRONO_24:
		return "Chrono24"
	}

	return "n/a"
}
