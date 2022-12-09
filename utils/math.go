package utils

func Floor[T ~int | int64](value, precision T) T {
	return (value / precision) * precision
}
