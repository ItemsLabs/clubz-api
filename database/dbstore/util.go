package dbstore

import "fmt"

func GenerateSubstituteParams(low, high int) string {
	str := ""
	for i := low; i <= high; i++ {
		if i > low {
			str += ","
		}
		str += fmt.Sprintf("$%d", i)
	}

	return str
}
