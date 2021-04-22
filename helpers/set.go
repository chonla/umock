package helpers

import (
	"sort"
)

func HasSomeInCommon(a, b []string) bool {
	sort.Strings(a)
	sort.Strings(b)

	for _, v := range a {
		for _, w := range b {
			if v == w {
				return true
			}
		}
	}
	return false
}
