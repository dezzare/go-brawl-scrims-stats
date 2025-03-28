package util

import (
	"strings"
)

func ParsePlayerTag(tag string) string {
	return "%23" + strings.TrimPrefix(strings.TrimPrefix(tag, "#"), "%23")
}
