package main

func mapper[T, U any](slice []T, fn func(T) U) []U {
	mappedSlice := make([]U, len(slice))
	for i, v := range slice {
		mappedSlice[i] = fn(v)
	}
	return mappedSlice
}
