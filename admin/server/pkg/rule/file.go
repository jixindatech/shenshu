package rule

import "gorm.io/datatypes"

type FileRules struct {
	Rules      []*OutputRuleItem      `json:"rules"`
	BatchRules []*OutputBatchRuleItem `json:"batchRule"`
}
type OutputBatchRuleItem struct {
	Type    string         `json:"type"`
	Args    datatypes.JSON `json:"args"`
	Pattern string         `json:"pattern"`
	Action  int            `json:"action"`
	Status  int            `json:"status"`
	Remark  string         `json:"remark"`
}

type OutputRuleItem struct {
	Type     string         `json:"type"`
	Rules    datatypes.JSON `json:"rules" `
	Action   int            `json:"action"`
	Priority int            `json:"priority"`
	Status   int            `json:"status"`
	Remark   string         `json:"remark"`
}
