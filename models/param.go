package models

import (
	"strings"

	"github.com/chonla/umock/logger"
)

type Param []string

func (p Param) Test(path map[string]string, log *logger.Logger) bool {
	log.Debug("    Path Parameters Condition(s): %d", len(p))
	pExpected := map[string]string{}
	for _, pValue := range p {
		pPair := strings.SplitN(pValue, "=", 2)
		if len(pPair) != 2 {
			log.Debug("    Unable to extract key-value from condition: %s", pValue)
			return false
		}
		if _, ok := pExpected[pPair[0]]; ok {
			log.Debug("    Duplicate key from condition: %s", pValue)
			return false
		}
		pExpected[pPair[0]] = pPair[1]
	}

	for kExpected, vExpected := range pExpected {
		if vActual, ok := path[kExpected]; ok {
			if vExpected != vActual {
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
