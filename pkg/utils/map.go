package utils

func MapFilter[K comparable, V any](rawMap map[K]V, op func(K, V) bool) map[K]V {
	newMap := map[K]V{}
	for key, value := range rawMap {
		if op(key, value) {
            newMap[key] = value
		}
	}

	return newMap
}

func MapEvery[K comparable, V any](rawMap map[K]V, op func(K, V) bool) bool {
	for key, val := range rawMap {
		if !op(key, val) {
			return false
		}
	}

	return true
}

func MapSome[K comparable, V any](rawMap map[K]V, op func(K, V) bool) bool {
	for key, val := range rawMap {
		if op(key, val) {
			return true
		}
	}

	return false
}
