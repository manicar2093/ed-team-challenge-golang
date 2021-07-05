package stringsfmt

import "strings"

func SeparateByCommas(elements ...string) string {
	builder := strings.Builder{}

	comma := len(elements) - 1
	for i, v := range elements {

		builder.WriteString(v)
		if i != comma {
			builder.WriteString(",")
		}
	}

	return builder.String()
}
