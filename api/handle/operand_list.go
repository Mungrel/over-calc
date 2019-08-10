package handle

import (
	"net/http"

	"github.com/Mungrel/over-calc/db/repo"
)

// ListOperands handles operand list GET requests.
func ListOperands(w http.ResponseWriter, r *http.Request) {
	operands, err := repo.ListOperands(r.Context())
	if err != nil {
		respondWithErrorAndLog(w, err)
		return
	}

	respondWithJSON(w, operands, http.StatusOK)
}
