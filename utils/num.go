package utils

import (
	"strconv"
)

func StrToInt64(str string) int64 {
	n, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return 0
	}
	return n
}
