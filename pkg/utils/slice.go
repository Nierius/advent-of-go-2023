package utils

type filterOperation func (any) bool
func filter[T any](slice []T, op filterOperation) []T {
    newSlice := []T{}
    for _, val := range slice {
        if op(val) {
            newSlice = append(newSlice, val)
        }
    }

    return newSlice
}
