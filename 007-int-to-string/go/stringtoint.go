package stringtoint

import "strconv"

func convert(i int) string {
	var s = strconv.Itoa(i)
	return s
}
