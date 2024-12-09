package utils

import "slices"

func FilterEmptyLines(data []string) []string {
	out := make([]string, 0, len(data))
	for _, line := range data {
		if line != "" {
			out = append(out, line)
		}
	}
	return out
}

func DeleteIndex[T any](array []T, idx int) []T {
	return slices.Concat(array[:idx], array[idx+1:])
}

func ReplaceStringAtIndex(s string, idx int, ch byte) string {
	out := ""
	for i := 0; i < len(s); i++ {
		if i == idx {
			out += string(ch)
		} else {
			out += string(s[i])
		}
	}
	return out
}
