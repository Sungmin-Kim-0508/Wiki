package myutils

import (
	"strings"
)

// SplitURLBySlash parse request url by slash
func SplitURLBySlash(url string) []string {
	parsedURL := strings.Split(url, "/")
	return parsedURL
}

// SliceNil slices the elements of nil
func SliceNil(arr []string) []string {
	parsedURL := make([]string, 0)
	i := 0
	for _, value := range arr {
		if len(value) > 0 {
			parsedURL = append(parsedURL, value)
			i++
		}
	}
	return parsedURL
}

// GetLastElement returns the last element of array
func GetLastElement(arr []string) string {
	return arr[len(arr)-1]
}
