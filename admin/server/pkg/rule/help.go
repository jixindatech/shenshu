package rule

import (
	"encoding/json"
	"errors"
	"gorm.io/datatypes"
)

type HostConfig struct {
	Host     string          `json:"host"`
	Decoders map[string]bool `json:"decoders"`
	Mode     string          `json:"mode"`
}

type CCItem struct {
	Method string `json:"method"`
	// Uri       string `json:"uri"`
	Mode      string `json:"mode"`
	Match     string `json:"match"`
	Threshold int    `json:"threshold"`
	Duration  int    `json:"duration"`
	Action    string `json:"action"`
}

type HsRuleItem struct {
	Id      uint     `json:"id"`
	Type    string   `json:"type"`
	Args    []string `json:"args"`
	Pattern string   `json:"pattern"`
	Action  int      `json:"action"`
}

type RuleItem struct {
	Variable string `json:"variable"`
	Operator string `json:"operator"`
	Pattern  string `json:"pattern"`
	Header   string `json:"header"`
}

type RuleConfig struct {
	Id     uint       `json:"id"`
	Type   string     `json:"type"`
	Action int        `json:"action"`
	Rules  []RuleItem `json:"rules"`
}

var variable = map[string]bool{
	"IP":         true,
	"METHOD":     true,
	"URI":        true,
	"QUERY":      true,
	"REQ_HEADER": true,
	"FILE":       true,
	"FILE_NAMES": true,
}

var operator = map[string]bool{
	// "not_exist":        true,
	"EQUALS": true,
	// "not_str_equal":    true,
	"STR_CONTAINS": true,
	// "not_str_contains": true,
	// "prefix_equal":     true,
	// "suffix_equal":     true,
	// "ip_contains":      true,
	// "not_ip_contains":  true,
	"GREATER": true,
	"LESS":    true,
	// "num_equal":        true,
	"REGEX": true,
	// "not_re_equal":     true,
}

var hsArgs = map[string]bool{
	"query":  true,
	"post":   true,
	"cookie": true,
}

func ValidateRuleInfo(rules []RuleItem) error {
	for _, v := range rules {
		_, ok := variable[v.Variable]
		if !ok {
			return errors.New("not exist variable")
		}
		_, ok = operator[v.Operator]
		if !ok {
			return errors.New("not exist operator")
		}
	}
	return nil
}

func ValidateRule(rule datatypes.JSON) error {
	var ruleInfo []RuleItem
	err := json.Unmarshal(rule, &ruleInfo)
	if err != nil {
		return err
	}
	return ValidateRuleInfo(ruleInfo)
}

func ValidateCCItem(item CCItem) error {
	if !(item.Duration > 0 && item.Threshold > 0) {
		return errors.New("threshold duration is not natural number")
	}
	_, ok := operator[item.Match]
	if !ok {
		return errors.New("match is not correct")
	}
	return nil
}

func ValidateHsArgs(args []string) error {
	for _, v := range args {
		_, ok := hsArgs[v]
		if !ok {
			return errors.New("not exist args")
		}
	}
	return nil
}

func ValidateBatchArgs(args datatypes.JSON) error {
	var ruleArgs []string
	err := json.Unmarshal(args, &ruleArgs)
	if err != nil {
		return err
	}
	return ValidateHsArgs(ruleArgs)
}
