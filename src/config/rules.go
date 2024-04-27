package config

import (
	errorFactory "IlluminateMyWords/src/utils/errors"

	regexp "github.com/dlclark/regexp2"
)

type MatchingRule struct {
	Regex *regexp.Regexp
	rule  string
	Color string
}

type RuleSet map[string]MatchingRule

func (rs RuleSet) Reset() {
}

func CreateRule(keywords *[]string, color string) (MatchingRule, error) {
	var rule string
	size := len(*keywords)

	if size == 0 {
		return MatchingRule{
			Regex: nil,
			rule:  "",
			Color: color,
		}, nil
	}

	rule = "(?i)(" + (*keywords)[0]

	for i := 1; i < size; i++ {
		rule += "|" + (*keywords)[i]
	}

	rule += ")"

	regex, err := regexp.Compile(rule, regexp.IgnoreCase)
	if err != nil {
		return MatchingRule{
			Regex: nil,
			rule:  rule,
			Color: color,
		}, errorFactory.BuildError(err, "")
	}

	return MatchingRule{
		Regex: regex,
		rule:  rule,
		Color: color,
	}, nil
}
