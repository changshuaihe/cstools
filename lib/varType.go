package lib

import "strconv"

func Int64ToStr(target int64) string {
	return strconv.FormatInt(target, 10)
}

func StrToInt(target string) (int, error) {
	return strconv.Atoi(target)
}

func StrToInt64(target string) (int64, error) {
	return strconv.ParseInt(target, 10, 64)
}

func IntToStr(target int) string {
	return strconv.Itoa(target)
}
