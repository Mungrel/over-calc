package handle

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Mungrel/over-calc/types"
)

func readBody(r *http.Request, dest interface{}) error {
	err := json.NewDecoder(r.Body).Decode(dest)
	if err != nil {
		return err
	}

	return r.Body.Close()
}

func respondWithStatus(w http.ResponseWriter, status int) {
	w.WriteHeader(status)
}

func respondWithJSON(w http.ResponseWriter, value interface{}, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	err := json.NewEncoder(w).Encode(value)
	if err != nil {
		panic(err)
	}
}

func respondWithErrorAndLog(w http.ResponseWriter, err error) {
	log.Println(err.Error())

	response := map[string]string{
		"error": "an internal error occurred",
	}
	status := http.StatusInternalServerError

	if knownError, isKnownError := err.(types.KnownError); isKnownError {
		response = map[string]string{
			"error": knownError.Error(),
		}
		status = knownError.StatusCode()
	}

	respondWithJSON(w, response, status)
}
