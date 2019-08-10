package types

import "github.com/shopspring/decimal"

// EvalExpression represents an expression evalution request in the API,
// and a history entry in the DB.
type EvalExpression struct {
	ID             string    `json:"id"`
	UserID         string    `json:"user_id,omitempty"`
	LeftOperandID  string    `json:"left_operand_id"`
	RightOperandID string    `json:"right_operand_id"`
	Operation      Operation `json:"operation"`
}

type EvalResult struct {
	Result decimal.Decimal `json:"result"`
}
