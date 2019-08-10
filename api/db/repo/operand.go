package repo

import (
	"context"
	"database/sql"

	"github.com/Mungrel/over-calc/db"

	"github.com/Mungrel/over-calc/auth"
	"github.com/Mungrel/over-calc/types"
	"github.com/gofrs/uuid"
)

// CreateOperand creates a new operand in the DB.
// It will generate the ID itself.
// It will use the user ID attached to the provided context.
func CreateOperand(ctx context.Context, operand types.Operand) (types.Operand, error) {
	id, err := uuid.NewV4()
	if err != nil {
		return types.Operand{}, err
	}

	operand.ID = id.String()
	operand.UserID = auth.ContextUser(ctx).ID

	const insertOperand = `
		INSERT INTO operand (
			id,
			user_id,
			value
		) VALUES (
			:id,
			:user_id,
			:value
		)`

	_, err = db.ContextDB(ctx).NamedExecContext(ctx, insertOperand, operand)
	if err != nil {
		return types.Operand{}, err
	}

	return operand, nil
}

// GetOperand returns a single operand by its ID for a user.
// It will use the user ID attached to the provided context.
// It will return nil if no operand was found for that ID.
func GetOperand(ctx context.Context, id string) (*types.Operand, error) {
	const getOperand = `
		SELECT *
		FROM operand
		WHERE id = ?
		AND user_id = ?`

	userID := auth.ContextUser(ctx).ID

	var operand types.Operand
	err := db.ContextDB(ctx).GetContext(ctx, &operand, getOperand, id, userID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	return &operand, nil
}

// ListOperands lists all operands for a user.
// It doesn't limit because cbf with pagination.
// It will use the user ID attached to the provided context.
func ListOperands(ctx context.Context) ([]types.Operand, error) {
	const listOperands = `
		SELECT *
		FROM operand
		WHERE user_id = ?`

	userID := auth.ContextUser(ctx).ID

	operands := []types.Operand{}
	err := db.ContextDB(ctx).SelectContext(ctx, &operands, listOperands, userID)
	if err != nil {
		return nil, err
	}

	return operands, nil
}
