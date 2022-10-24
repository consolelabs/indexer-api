package utils

func InlinePointer[T any](value T) *T {
	return &value
}
