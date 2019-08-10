package repo

import (
	"context"

	"github.com/Mungrel/over-calc/auth"
	"github.com/Mungrel/over-calc/db"

	"github.com/Mungrel/over-calc/types"
	"github.com/gofrs/uuid"
)

// CreateHistoryEntry creates a history entry in the DB.
// It will generate the ID itself.
// It will use the user ID attached to the provided context.
func CreateHistoryEntry(ctx context.Context, entry types.EvalExpression) (types.EvalExpression, error) {
	id, err := uuid.NewV4()
	if err != nil {
		return types.EvalExpression{}, err
	}

	entry.ID = id.String()

	const insertHistoryEntry = `
		INSERT INTO history (
			id,
			user_id,
			left_operand_id,
			right_operand_id,
			operation
		) VALUES (
			:id,
			:user_id,
			:left_operand_id,
			:right_operand_id,
			:operation
		)`

	_, err = db.ContextDB(ctx).NamedExecContext(ctx, insertHistoryEntry, entry)
	if err != nil {
		return types.EvalExpression{}, err
	}

	return entry, nil
}

// ListHistoryEntries lists all history entries for a user.
// It will use the user ID attached to the provided context.
func ListHistoryEntries(ctx context.Context) ([]types.EvalExpression, error) {
	const listHistoryEntries = `
		SELECT *
		FROM history
		WHERE user_id = ?`

	userID := auth.ContextUser(ctx).ID

	entries := []types.EvalExpression{}

	err := db.ContextDB(ctx).SelectContext(ctx, &entries, listHistoryEntries, userID)
	if err != nil {
		return []types.EvalExpression{}, err
	}

	return entries, nil
}
