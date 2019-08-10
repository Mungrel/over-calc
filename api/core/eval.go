package core

import (
	"context"
	"net/http"

	"github.com/Mungrel/over-calc/db/repo"
	"github.com/shopspring/decimal"

	"github.com/Mungrel/over-calc/types"
)

var (
	ErrNoOperand        = types.NewKnownError("no operand found for id", http.StatusNotFound)
	ErrInvalidOperation = types.NewKnownError("invalid operation", http.StatusBadRequest)
	ErrDivZero          = types.NewKnownError("look at this chump trying to divide by zero", http.StatusBadRequest)
)

// EvaluateExpression evaluates an expression and returns the result.
// It also makes an entry in the history table.
func EvaluateExpression(ctx context.Context, exp types.EvalExpression) (types.EvalResult, error) {
	left, err := repo.GetOperand(ctx, exp.LeftOperandID)
	if err != nil {
		return types.EvalResult{}, err
	}

	right, err := repo.GetOperand(ctx, exp.RightOperandID)
	if err != nil {
		return types.EvalResult{}, err
	}

	if left == nil || right == nil {
		return types.EvalResult{}, ErrNoOperand
	}

	result, err := eval(left.Value, right.Value, exp.Operation)
	if err != nil {
		return types.EvalResult{}, err
	}

	_, err = repo.CreateHistoryEntry(ctx, exp)
	if err != nil {
		return types.EvalResult{}, err
	}

	return types.EvalResult{
		Result: result,
	}, nil
}

func eval(left, right decimal.Decimal, op types.Operation) (decimal.Decimal, error) {
	var result decimal.Decimal
	switch op {
	case types.Add:
		result = left.Add(right)
	case types.Subtract:
		result = left.Sub(right)
	case types.Multiply:
		result = left.Mul(right)
	case types.Divide:
		if right.IsZero() {
			return decimal.Decimal{}, ErrDivZero
		}
		result = left.Div(right)
	default:
		return decimal.Decimal{}, ErrInvalidOperation
	}

	return result, nil
}
