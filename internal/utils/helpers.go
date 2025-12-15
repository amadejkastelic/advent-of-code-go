package utils

import "strings"

func ValOrPanic[T any](val T, err error) T {
	if err != nil {
		panic(err)
	}
	return val
}

func SplitLines(s string) []string {
	return strings.Split(strings.Trim(string(s), "\n"), "\n")
}
