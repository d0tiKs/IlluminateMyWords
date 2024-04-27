package highlights

import (
	"IlluminateMyWords/src/config"
	errorfactory "IlluminateMyWords/src/utils/errors"
	"bufio"
	"fmt"
	"io"

	regexp "github.com/dlclark/regexp2"

	"github.com/fatih/color"
)

func ColorizeKeyword(keyword string, colorname string) string {
	colorCode := config.ColorCode[colorname]

	if colorCode == 0 {
		colorCode = config.ColorCode[config.DEFAULT_COLOR]
	}

	colorizer := color.New(colorCode)
	return colorizer.Sprintf("%s", keyword)
}

func ColorizeLine(line string, regex *regexp.Regexp, color string) string {
	newline, err := regex.Replace(line, ColorizeKeyword("$1", color), -1, -1)
	if err != nil {
		errorfactory.BuildAndLogError(err, "error while replacing with colored keywords")
		return line
	}
	return newline
}

func ColorizeStream(reader io.Reader) {
	config := config.GetConfig()
	scanner := bufio.NewScanner(reader)
	for i := 0; scanner.Scan(); i++ {
		line := scanner.Text()

		for keyword, rule := range config.Rules {
			newLine := ColorizeLine(line, rule.Regex, config.Rules[keyword].Color)
			line = newLine
		}
		fmt.Println(line)
	}
}
