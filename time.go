package cstools

import "time"

func GetTsInt64() int64 {
	return time.Now().Unix()
}

func GetTsStr() string {
	ts := time.Now().Unix()
	tsStr := Int64ToStr(ts)
	return tsStr
}

func GetTsNano() int64 {
	return time.Now().UnixNano()
}

func GetTsNanoStr() string {
	ts := time.Now().UnixNano()
	tsStr := Int64ToStr(ts)
	return tsStr
}
