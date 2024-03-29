package client

import (
	_ "embed"
	"fmt"
)

type Spec struct {
	Accounts          []AccountSpec `json:"accounts,omitempty" jsonschema:"required,minItems=1"`
	BillHistoryMonths int           `json:"bill_history_months,omitempty"`
	Concurrency       int           `json:"concurrency,omitempty"`
}

type AccountSpec struct {
	Name      string   `json:"name,omitempty" jsonschema:"required,minLength=1"`
	Regions   []string `json:"regions,omitempty" jsonschema:"required,minItems=1"`
	AccessKey string   `json:"access_key,omitempty" jsonschema:"required,minLength=1"`
	SecretKey string   `json:"secret_key,omitempty" jsonschema:"required,minLength=1"`
}

func (s *Spec) SetDefaults() {
	if s.Concurrency == 0 {
		s.Concurrency = 50000
	}
	if s.BillHistoryMonths == 0 {
		s.BillHistoryMonths = 12
	}
}

func (s *Spec) Validate() error {
	if len(s.Accounts) == 0 {
		return fmt.Errorf("missing alicloud accounts in configuration")
	}
	accountNames := map[string]struct{}{}
	for _, account := range s.Accounts {
		if account.Name == "" {
			return fmt.Errorf("missing alicloud account name in configuration")
		}
		if _, found := accountNames[account.Name]; found {
			return fmt.Errorf("duplicate alicloud account name %s in configuration", account.Name)
		}
		accountNames[account.Name] = struct{}{}
		if account.AccessKey == "" {
			return fmt.Errorf("missing access_key in account configuration for account %s", account.Name)
		}
		if account.SecretKey == "" {
			return fmt.Errorf("missing secret_key in account configuration for account %s", account.Name)
		}
		if len(account.Regions) == 0 {
			return fmt.Errorf("missing regions in account configuration for account %s", account.Name)
		}
	}
	return nil
}

//go:embed schema.json
var JSONSchema string
