package models

import (
	"encoding/json"
	"net/url"
	"strings"

	"github.com/chonla/umock/helpers"
	"github.com/chonla/umock/logger"
)

type Form []string

func (f Form) Test(values url.Values, log *logger.Logger) bool {
	var valuesMap map[string][]string
	b, _ := json.Marshal(values)
	json.Unmarshal(b, &valuesMap)

	log.Debug("    Form Body Condition(s): %d", len(f))
	fExpected := map[string][]string{}
	for _, fValue := range f {
		fPair := strings.SplitN(fValue, "=", 2)
		if len(fPair) != 2 {
			log.Debug("    Unable to extract key-value from condition: %s", fValue)
			return false
		}
		if _, ok := fExpected[fPair[0]]; !ok {
			fExpected[fPair[0]] = []string{}
		}
		fExpected[fPair[0]] = append(fExpected[fPair[0]], fPair[1])
	}

	for kExpected, vExpected := range fExpected {
		if vActual, ok := valuesMap[kExpected]; ok {
			if !helpers.HasSomeInCommon(vExpected, vActual) {
				log.Debug("    Matching %s ... FAILED", kExpected)
				return false
			}
			log.Debug("    Matching %s ... PASSED", kExpected)
		} else {
			// Request does not contain some expected value.
			log.Debug("    Matching %s ... FAILED", kExpected)
			return false
		}
	}
	return true
}
