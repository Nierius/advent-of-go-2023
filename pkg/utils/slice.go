package utils

func Filter[T any](slice []T, op func (T) bool) []T {
    newSlice := []T{}
    for _, val := range slice {
        if op(val) {
            newSlice = append(newSlice, val)
        }
    }

    return newSlice
}

func Every[T any](slice []T, op func (T) bool) bool {
    for _, val := range slice {
        if !op(val) {
            return false
        }
    }

    return true
}

func Some[T any](slice []T, op func (T) bool) bool {
    for _, val := range slice {
        if op(val) {
            return true
        }
    }

    return false
}

func Includes[T comparable](slice []T, val T) bool {
    for _, itVal := range slice {
        if itVal == val {
            return true
        }
    }

    return false
}

func Map[T any, V any](slice []T, op func (T) V) []V {
    newSlice := []V{}
    for _, val := range slice {
        newSlice = append(newSlice, op(val))
    }

    return newSlice
}
