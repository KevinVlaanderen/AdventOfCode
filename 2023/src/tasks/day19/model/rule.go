package model

import (
	"strings"
)

type Rule struct {
	Condition   Condition
	Instruction Instruction
}

func ParseRules(data string) []Rule {
	ruleStrings := strings.Split(data, ",")
	rules := make([]Rule, len(ruleStrings))
	for ruleIndex, ruleString := range ruleStrings {
		ruleParts := strings.Split(ruleString, ":")
		switch len(ruleParts) {
		case 1:
			rule := Rule{Always(), ParseInstruction(ruleParts[0])}
			rules[ruleIndex] = rule
		case 2:
			rule := Rule{ParseCondition(ruleParts[0]), ParseInstruction(ruleParts[1])}
			rules[ruleIndex] = rule
		}
	}
	return rules
}
