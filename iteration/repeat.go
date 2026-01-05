package iteration

import "strings"

func Repeat(character string, times int) string {
	// repeated := ""
	// for range times {
	// 	repeated += character
	// }
	// return repeated

	var repeated strings.Builder
	for range times {
		repeated.WriteString(character)
	}
	return repeated.String()

}
