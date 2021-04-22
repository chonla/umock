package models

import "strings"

type ContentType string

func (c ContentType) Test(targetContentType string) bool {
	t := string(c)
	if t == "" {
		t = "x-www-form-urlencoded"
	}
	return strings.ToUpper(t) == strings.ToUpper(targetContentType)
}
