package extractor

import (
	"bufio"
	"os"
	"regexp"
	"strings"
)

type Functions struct {
	Name string
	Body string
	File string
}

var funcOpenRegex = regexp.MustCompile(`^func (\(.+\) )?(?P<funcName>\S+)\(.*\)`)
var funcCloseRegex = regexp.MustCompile(`^}`)

func ExtractFunctionsFromFile(filePath string) ([]Functions, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var rtnMe []Functions

	scanner := bufio.NewScanner(file)

	var currBody strings.Builder
	var funcName string
	funcOpen := false

	for scanner.Scan() {
		curr := scanner.Text()
		if funcOpen {
			currBody.WriteString(curr + "\n")
			if funcCloseRegex.MatchString(curr) {
				funcOpen = false
				rtnMe = append(rtnMe, Functions{
					Name: funcName,
					Body: currBody.String(),
					File: filePath,
				})
				currBody = strings.Builder{}
				funcName = ""
			}
		} else {
			if funcOpenRegex.MatchString(curr) {
				funcOpen = true
				currBody.WriteString(curr + "\n")
				match := funcOpenRegex.FindStringSubmatch(curr)
				funcName = match[funcOpenRegex.SubexpIndex("funcName")]

			}
		}

	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return rtnMe, nil
}
