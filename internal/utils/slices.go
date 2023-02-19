package utils

func GroupBy[TItem any, TKey comparable](items []TItem, getKey func(TItem) TKey) map[TKey][]TItem {
	groups := make(map[TKey][]TItem, 0)

	for _, item := range items {
		key := getKey(item)

		if groups[key] == nil {
			groups[key] = make([]TItem, 0)
		}

		groups[key] = append(groups[key], item)
	}

	return groups
}
