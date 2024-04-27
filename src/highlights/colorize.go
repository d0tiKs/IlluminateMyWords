package highlights

import (
	"IlluminateMyWords/src/config"
	"bufio"
	"fmt"
	"io"
	"regexp"

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
	return regex.ReplaceAllString(line, ColorizeKeyword("$1", color))
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
