package utils

import "strconv"

func JoinInts(v []int, sep string) string {
	r := ""
	for _, v := range v {
		if len(r) > 0 {
			r += sep
		}
		r += strconv.Itoa(v)
	}
	return r
}
