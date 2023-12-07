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
