package model

import (
	"strings"
)

type Rule struct {
	Condition *Condition
	Next      string
}

func ParseRules(data string) []Rule {
	ruleStrings := strings.Split(data, ",")
	rules := make([]Rule, len(ruleStrings))
	for ruleIndex, ruleString := range ruleStrings {
		ruleParts := strings.Split(ruleString, ":")
		switch len(ruleParts) {
		case 1:
			rule := Rule{nil, ruleParts[0]}
			rules[ruleIndex] = rule
		case 2:
			condition := ParseCondition(ruleParts[0])
			rule := Rule{&condition, ruleParts[1]}
			rules[ruleIndex] = rule
		}
	}
	return rules
}
