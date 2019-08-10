package repo

import (
	"context"
	"database/sql"

	"github.com/Mungrel/over-calc/db"

	"github.com/Mungrel/over-calc/types"
	"github.com/gofrs/uuid"
	"golang.org/x/crypto/bcrypt"
)

// AttemptSignIn checks a user's password and username when signing in.
// It will return the user if username and password match, otherwise it will return nil.
func AttemptSignIn(ctx context.Context, signInRequest types.SignInRequest) (*types.User, error) {
	const getUser = `
		SELECT *
		FROM user
		WHERE username = ?`

	var user types.User
	err := db.ContextDB(ctx).GetContext(ctx, &user, getUser, signInRequest.Username)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	if !checkPassword(signInRequest.Password, user.Password) {
		return nil, nil
	}

	return &user, nil
}

// CreateUser creates a user in the DB.
// It will generate the user's ID itself.
func CreateUser(ctx context.Context, user types.User) (types.User, error) {
	id, err := uuid.NewV4()
	if err != nil {
		return types.User{}, err
	}

	user.ID = id.String()
	user.Password, err = encryptPassword(user.Password)
	if err != nil {
		return types.User{}, err
	}

	const insertUser = `
		INSERT INTO user (
			id,
			username,
			password
		) VALUES (
			:id,
			:username,
			:password
		)`

	_, err = db.ContextDB(ctx).NamedExecContext(ctx, insertUser, user)
	if err != nil {
		return types.User{}, err
	}

	return user, nil
}

// encryptPassword returns a bcrypt encrypted version of the supplied password.
func encryptPassword(password string) (string, error) {
	encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return "", err
	}

	return string(encryptedPassword), nil
}

// checkPassword returns true if the bcrypt encrypted version of the supplied password attempt
// matches the stored encrypted password.
func checkPassword(password, storedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(storedPassword), []byte(password))
	return err == nil
}
