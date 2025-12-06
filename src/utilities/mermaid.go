package utilities

import "strings"

func ConvertMermaidString(s string) string {
	return strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(s, "<", "~"), ">", "~"), ",", ".")
}
