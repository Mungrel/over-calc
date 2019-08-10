package handle

import (
	"net/http"

	"github.com/Mungrel/over-calc/db/repo"
)

// ListHistory handles history list GET requests.
func ListHistory(w http.ResponseWriter, r *http.Request) {
	entries, err := repo.ListHistoryEntries(r.Context())
	if err != nil {
		respondWithErrorAndLog(w, err)
		return
	}

	respondWithJSON(w, entries, http.StatusOK)
}
