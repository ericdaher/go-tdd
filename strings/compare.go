package strings

import "strings"

func Compare(str1, str2 string) bool {
	comparison := strings.Compare(str1, str2)
	return comparison == 0
}
