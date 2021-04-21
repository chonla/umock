package models

import (
	"encoding/json"
	"net/url"
	"strings"

	"github.com/chonla/umock/helpers"
)

type Form []string

func (f Form) Test(values url.Values) bool {
	var valuesMap map[string][]string
	b, _ := json.Marshal(values)
	json.Unmarshal(b, &valuesMap)

	fExpected := map[string][]string{}
	for _, fValue := range f {
		fPair := strings.SplitN(fValue, "=", 2)
		if _, ok := fExpected[fPair[0]]; !ok {
			fExpected[fPair[0]] = []string{}
		}
		fExpected[fPair[0]] = append(fExpected[fPair[0]], fPair[1])
	}

	for kExpected, vExpected := range fExpected {
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
