package ignoreParser

import "regexp"

type Ignore struct {
	Ignores []*regexp.Regexp
}


func (i *Ignore) ShouldIgnore(fileName string) bool  {
	for _, reg := range i.Ignores {
		if reg.MatchString(fileName) {
			return true
		}
	}
	return false
}