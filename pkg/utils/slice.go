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
