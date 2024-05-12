package util

func Set[T any](receiver **T, value T) {
	*receiver = &value
}
