package models

import "strings"

type Method string

func (m Method) Test(targetMethod string) bool {
	return strings.ToUpper(string(m)) == strings.ToUpper(targetMethod)
}
