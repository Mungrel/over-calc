package handle

import (
	"net/http"

	"github.com/Mungrel/over-calc/core"
	"github.com/Mungrel/over-calc/types"
)

// Eval handles evaluation POST requests.
func Eval(w http.ResponseWriter, r *http.Request) {
	var exp types.EvalExpression
	err := readBody(r, &exp)
	if err != nil {
		respondWithErrorAndLog(w, err)
		return
	}

	ctx := r.Context()
	result, err := core.EvaluateExpression(ctx, exp)
	if err != nil {
		respondWithErrorAndLog(w, err)
		return
	}

	respondWithJSON(w, result, http.StatusOK)
}
