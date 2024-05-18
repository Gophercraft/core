package util

// Generic function to set those extremely annoying Proto2 pointer-to fields.
func Set[T any](receiver **T, value T) {
	*receiver = &value
}
