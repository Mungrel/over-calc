package handle

import (
	"net/http"

	"github.com/Mungrel/over-calc/db/repo"
	"github.com/Mungrel/over-calc/types"
)

// CreateOperand handles operand creation POST requests.
func CreateOperand(w http.ResponseWriter, r *http.Request) {
	var operand types.Operand
	err := readBody(r, &operand)
	if err != nil {
		respondWithErrorAndLog(w, err)
		return
	}

	savedOperand, err := repo.CreateOperand(r.Context(), operand)
	if err != nil {
		respondWithErrorAndLog(w, err)
		return
	}

	respondWithJSON(w, savedOperand, http.StatusCreated)
}
