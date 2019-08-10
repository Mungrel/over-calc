package types

import "github.com/shopspring/decimal"

// Operand represents an equation's operand in the DB and API.
type Operand struct {
	ID     string          `json:"id"`
	UserID string          `json:"user_id"`
	Value  decimal.Decimal `json:"value"`
}
