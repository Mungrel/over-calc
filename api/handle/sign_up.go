package handle

import (
	"net/http"

	"github.com/Mungrel/over-calc/auth"
	"github.com/Mungrel/over-calc/db/repo"
	"github.com/Mungrel/over-calc/types"
)

// SignUp handles sign up POST requests.
func SignUp(w http.ResponseWriter, r *http.Request) {
	var user types.User
	err := readBody(r, &user)
	if err != nil {
		respondWithErrorAndLog(w, err)
		return
	}

	createdUser, err := repo.CreateUser(r.Context(), user)
	if err != nil {
		respondWithErrorAndLog(w, err)
		return
	}

	userClaims := &auth.UserClaims{
		ID:       createdUser.ID,
		Username: createdUser.Username,
	}

	jwt, err := auth.BuildJWT(userClaims)
	if err != nil {
		respondWithStatus(w, http.StatusUnauthorized)
		return
	}

	response := types.SignInResponse{
		ID:       createdUser.ID,
		Username: createdUser.Username,
		Token:    jwt,
	}

	respondWithJSON(w, response, http.StatusCreated)
}
