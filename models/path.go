package models

import (
	"errors"
	"fmt"
	"regexp"
)

type Path string

func (p Path) Test(targetPath string) bool {
	patternizedSegment := regexp.MustCompile("((?i):[a-z0-9_-]+)")
	if ok := patternizedSegment.MatchString(string(p)); ok {
		patternizedPath := patternizedSegment.ReplaceAllString(string(p), "[^/ ]+")
		pattern := regexp.MustCompile(fmt.Sprintf("^%s$", patternizedPath))
		return pattern.MatchString(targetPath)
	}
	return string(p) == targetPath
}

func (p Path) ExtractParameters(targetPath string) (map[string]string, error) {
	patternizedSegment := regexp.MustCompile("((?i):[a-z0-9_-]+)")
	if ok := patternizedSegment.MatchString(string(p)); ok {
		patternizedPath := patternizedSegment.ReplaceAllStringFunc(string(p), func(key string) string {
			return fmt.Sprintf("(?P<%s>[^/]+)", key[1:])
		})
		pattern := regexp.MustCompile(fmt.Sprintf("^%s$", patternizedPath))
		matches := pattern.FindStringSubmatch(targetPath)
		if len(matches) == 0 {
			return map[string]string{}, errors.New("unable to extract parameters from path")
		}

		output := map[string]string{}
		for i, name := range pattern.SubexpNames() {
			if i != 0 && name != "" {
				output[name] = matches[i]
			}
		}
		return output, nil
	}
	return map[string]string{}, errors.New("path does not contain parameters")
}

// * wut เข้าป่า
