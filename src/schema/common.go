package schema

import (
	"encoding/json"
	"fmt"
)

// PercentRule ...
type PercentRule struct {
	Kind string `json:"kind"`
	Rule struct {
		Value int `json:"value"`
	} `json:"rule"`
}

// Validate ...
func (r *PercentRule) Validate() error {
	if r.Kind != "userPercent" {
		return fmt.Errorf("invalid kind: %s", r.Kind)
	}
	if r.Rule.Value < 0 || r.Rule.Value > 100 {
		return fmt.Errorf("invalid percent rule value: %d", r.Rule.Value)
	}
	return nil
}

// ToRule ...
func (r *PercentRule) ToRule() string {
	if b, err := json.Marshal(r.Rule); err == nil {
		return string(b)
	}
	return ""
}

// ToPercentRule ...
func ToPercentRule(kind, rule string) *PercentRule {
	r := &PercentRule{Kind: "userPercent"}
	r.Rule.Value = -1
	if kind == "userPercent" && rule != "" {
		if err := json.Unmarshal([]byte(rule), &r.Rule); err != nil {
			r.Rule.Value = -1
		}

		if err := r.Validate(); err != nil {
			r.Rule.Value = -1
		}
	}

	return r
}

// ToRuleObject ...
func ToRuleObject(kind, rule string) interface{} {
	switch kind {
	case "userPercent":
		return ToPercentRule(kind, rule).Rule
	default:
		return struct{}{}
	}
}
