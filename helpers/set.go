package helpers

import (
	"sort"

	"github.com/kr/pretty"
)

func HasSomeInCommon(a, b []string) bool {
	sort.Strings(a)
	sort.Strings(b)

	pretty.Println(a, b)

	for _, v := range a {
		for _, w := range b {
			if v == w {
				return true
			}
		}
	}
	return false
}
