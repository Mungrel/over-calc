package auth

import (
	"crypto/rsa"
	"errors"
	"io/ioutil"

	"github.com/dgrijalva/jwt-go"
)

type UserClaims struct {
	ID       string `json:"id"`
	Username string `json:"username"`
}

type Claims struct {
	User *UserClaims `json:"user"`
	jwt.StandardClaims
}

var (
	privateKey *rsa.PrivateKey
	publicKey  *rsa.PublicKey
)

func init() {
	privateKeyContents, err := ioutil.ReadFile("./etc/keys/1_private.pem")
	if err != nil {
		panic(err)
	}

	privateKey, err = jwt.ParseRSAPrivateKeyFromPEM(privateKeyContents)
	if err != nil {
		panic(err)
	}

	publicKeyContents, err := ioutil.ReadFile("./etc/keys/1.pem")
	if err != nil {
		panic(err)
	}

	publicKey, err = jwt.ParseRSAPublicKeyFromPEM(publicKeyContents)
	if err != nil {
		panic(err)
	}
}

// BuildJWT builds a JWT from the provided user.
func BuildJWT(user *UserClaims) (string, error) {
	claims := &Claims{
		User: user,
	}

	return jwt.NewWithClaims(jwt.SigningMethodRS512, claims).SignedString(privateKey)
}

// validateJWT validates the provided token and returns its claims.
// It will return an error if that claims are not valid.
func validateJWT(token string) (*Claims, error) {
	parsedToken, err := jwt.ParseWithClaims(token, new(Claims), func(_ *jwt.Token) (interface{}, error) {
		return publicKey, nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := parsedToken.Claims.(*Claims); ok && parsedToken.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid claims")
}
