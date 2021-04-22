package models

import (
	"encoding/json"
	"net/url"
	"strings"

	"github.com/chonla/umock/helpers"
	"github.com/chonla/umock/logger"
)

type Query []string

func (q Query) Test(values url.Values, log *logger.Logger) bool {
	var valuesMap map[string][]string
	b, _ := json.Marshal(values)
	json.Unmarshal(b, &valuesMap)

	log.Debug("      Form Body Condition(s): %d", len(q))
	qExpected := map[string][]string{}
	for _, qValue := range q {
		qPair := strings.SplitN(qValue, "=", 2)
		if len(qPair) != 2 {
			log.Debug("      Unable to extract key-value from condition: %s", qValue)
			return false
		}
		if _, ok := qExpected[qPair[0]]; !ok {
			qExpected[qPair[0]] = []string{}
		}
		qExpected[qPair[0]] = append(qExpected[qPair[0]], qPair[1])
	}

	for kExpected, vExpected := range qExpected {
		if vActual, ok := valuesMap[kExpected]; ok {
			if !helpers.HasSomeInCommon(vExpected, vActual) {
				log.Debug("      Matching %s ... FAILED", kExpected)
				return false
			}
			log.Debug("      Matching %s ... PASSED", kExpected)
		} else {
			// Request does not contain some expected value.
			log.Debug("      Matching %s ... FAILED", kExpected)
			return false
		}
	}
	return true
}
