package handle

import (
	"net/http"

	"github.com/Mungrel/over-calc/auth"

	"github.com/Mungrel/over-calc/db/repo"

	"github.com/Mungrel/over-calc/types"
)

// SignIn handles sign in POST requests.
func SignIn(w http.ResponseWriter, r *http.Request) {
	var signInRequest types.SignInRequest
	err := readBody(r, &signInRequest)
	if err != nil {
		respondWithErrorAndLog(w, err)
		return
	}

	user, err := repo.AttemptSignIn(r.Context(), signInRequest)
	if err != nil {
		respondWithErrorAndLog(w, err)
		return
	}

	if user == nil {
		respondWithStatus(w, http.StatusUnauthorized)
		return
	}

	userClaims := &auth.UserClaims{
		ID:       user.ID,
		Username: user.Username,
	}

	jwt, err := auth.BuildJWT(userClaims)
	if err != nil {
		respondWithStatus(w, http.StatusUnauthorized)
		return
	}

	response := types.SignInResponse{
		ID:       user.ID,
		Username: user.Username,
		Token:    jwt,
	}

	respondWithJSON(w, response, http.StatusOK)
}
