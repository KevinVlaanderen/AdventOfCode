package model

import (
	"regexp"
)

var workflowPattern = regexp.MustCompile(`(?m)^(\w+)\{(.+)}$`)

func ParseWorkflows(data string) map[string][]Rule {
	workflowMatches := workflowPattern.FindAllStringSubmatch(data, -1)
	workflows := make(map[string][]Rule, len(workflowMatches))
	for _, workflowMatch := range workflowMatches {
		workflows[workflowMatch[1]] = ParseRules(workflowMatch[2])
	}
	return workflows
}
