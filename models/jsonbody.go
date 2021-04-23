package models

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/chonla/umock/logger"

	"github.com/tidwall/gjson"
)

type JsonBody []string

func (j JsonBody) Test(r *http.Request, log *logger.Logger) bool {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		// JSON cannot be parsed
		return false
	}

	r.Body = ioutil.NopCloser(bytes.NewBuffer(b))

	body := string(b)

	log.Debug("    JSON Body Condition(s): %d", len(j))
	for _, jValue := range j {
		jPair := strings.SplitN(jValue, "=", 2)
		if len(jPair) != 2 {
			log.Debug("    Unable to extract key-value from condition: %s", jValue)
			return false
		}
		value := gjson.Get(body, jPair[0])
		if jPair[1] != value.String() {
			log.Debug("    Matching %s ... FAILED", jPair[0])
			return false
		}
		log.Debug("    Matching %s ... PASSED", jPair[0])
	}

	return true
}
