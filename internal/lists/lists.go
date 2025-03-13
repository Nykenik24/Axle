package lists

import (
	"fmt"
	"strings"
)

const (
	reset  = "\033[0m"
	green  = "\033[32m"
	yellow = "\033[33m"
	cyan   = "\033[36m"
)

func UnorderedList(elements []any, bullet string) string {
	var sb strings.Builder

	for _, element := range elements {
		if element != nil {
			sb.WriteString(fmt.Sprintf("%s %v\n", bullet, element))
		}
	}

	return sb.String()
}

func OrderedList(elements []any) string {
	var sb strings.Builder

	for index, element := range elements {
		if element != nil {
			sb.WriteString(fmt.Sprintf("%v. %v\n", index, element))
		}
	}

	return sb.String()
}

func KeyList(elements map[string]any) string {
	var sb strings.Builder

	for key, element := range elements {
		if element != nil {
			sb.WriteString(fmt.Sprintf("%v: %v\n", key, element))
		}
	}

	return sb.String()
}
