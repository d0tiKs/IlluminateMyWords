package config

import (
	errorFactory "IlluminateMyWords/src/utils/errors"
	"regexp"
)

type MatchingRule struct {
	regex *regexp.Regexp
	rule  string
}

type RuleSet map[string]MatchingRule

func (rs RuleSet) Reset() {
}

func createRule(keywords *[]string) (MatchingRule, error) {
	var rule string
	size := len(*keywords)

	if size == 0 {
		return MatchingRule{
			regex: nil,
			rule:  "",
		}, nil
	}

	rule = "(" + (*keywords)[0]

	for i := 1; i < size; i++ {
		rule += "|" + (*keywords)[i]
	}

	rule += ")"

	regex, err := regexp.Compile(rule)
	if err != nil {
		return MatchingRule{
			regex: nil,
			rule:  rule,
		}, errorFactory.BuildError(err, "")
	}

	return MatchingRule{
		regex: regex,
		rule:  rule,
	}, nil
}

func (mr *MatchingRule) Evaluate() {
}
