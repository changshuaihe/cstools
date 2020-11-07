package lib

import "sort"

func SortStrArray(target []string) []string {
	sort.Sort(sort.StringSlice(target))
	return target
}
