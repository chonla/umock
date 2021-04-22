package models

import (
	"io"
	"io/ioutil"
	"strings"

	"github.com/chonla/umock/logger"
	"github.com/tidwall/gjson"
)

type JsonBody []string

func (j JsonBody) Test(r io.ReadCloser, log *logger.Logger) bool {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		// JSON cannot be parsed
		return false
	}

	body := string(b)

	log.Debug("      JSON Body Condition(s): %d", len(j))
	for _, jValue := range j {
		jPair := strings.SplitN(jValue, "=", 2)
		if len(jPair) != 2 {
			log.Debug("      Unable to extract key-value from condition: %s", jValue)
			return false
		}
		value := gjson.Get(body, jPair[0])
		if jPair[1] != value.String() {
			log.Debug("      Matching %s ... FAILED", jPair[1])
			return false
		}
		log.Debug("      Matching %s ... PASSED", jPair[1])
	}

	return true
}
