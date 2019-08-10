package auth

import (
	"context"
	"net/http"
)

type Authenticator struct {
	inner http.Handler
}

func NewAuthenticator(inner http.Handler) http.Handler {
	return &Authenticator{
		inner: inner,
	}
}

// ServeHTTP implements the http.Handler interface.
func (authenticator *Authenticator) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		authenticator.inner.ServeHTTP(w, r)
		return
	}

	header := r.Header.Get("Authorization")
	if header == "" {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	claims, err := validateJWT(header)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	ctx := contextWithClaims(r.Context(), claims)

	authenticator.inner.ServeHTTP(w, r.WithContext(ctx))
}

type claimsContextKey int

const claimsCtxKey claimsContextKey = iota

func contextWithClaims(ctx context.Context, claims *Claims) context.Context {
	return context.WithValue(ctx, claimsCtxKey, claims)
}

func ContextClaims(ctx context.Context) *Claims {
	return ctx.Value(claimsCtxKey).(*Claims)
}

func ContextUser(ctx context.Context) *UserClaims {
	return ContextClaims(ctx).User
}
