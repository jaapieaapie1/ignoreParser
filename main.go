package ignoreParser

import (
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

func ParseIgnoreFile(file *os.File) (*Ignore, error) {
	data, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}
	ignoreString := string(data)
	reg := regexp.MustCompile("\r?\n")

	unfilteredLines := reg.Split(ignoreString, -1)
	var lines []*regexp.Regexp

	for _, text := range unfilteredLines {
		if t := strings.Trim(text, " "); t != "" && string(text[0]) != "#" {
			lines = append(lines, regexp.MustCompile(text))
		}
	}

	return &Ignore{
		Ignores: lines,
	}, nil
}