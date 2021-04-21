package models

import (
	"encoding/json"
	"net/url"
	"strings"

	"github.com/chonla/umock/helpers"
)

type Query []string

func (q Query) Test(values url.Values) bool {
	var valuesMap map[string][]string
	b, _ := json.Marshal(values)
	json.Unmarshal(b, &valuesMap)

	qExpected := map[string][]string{}
	for _, qValue := range q {
		qPair := strings.SplitN(qValue, "=", 2)
		if _, ok := qExpected[qPair[0]]; !ok {
			qExpected[qPair[0]] = []string{}
		}
		qExpected[qPair[0]] = append(qExpected[qPair[0]], qPair[1])
	}

	for kExpected, vExpected := range qExpected {
		if vActual, ok := valuesMap[kExpected]; ok {
			if !helpers.HasSomeInCommon(vExpected, vActual) {
				return false
			}
		} else {
			// Request does not contain some expected value.
			return false
		}
	}
	return true
}
