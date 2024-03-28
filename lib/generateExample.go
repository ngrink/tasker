package lib

import (
	"strings"
)

func GenerateExample(examples []string) string {
	var builder strings.Builder

	for i, ex := range examples {
		builder.WriteByte(' ')
		builder.WriteByte(' ')
		builder.WriteString(ex)
		if i < len(examples)-1 {
			builder.WriteByte('\n')
		}
	}

	return builder.String()
}
