package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v4"
	"github.com/spazzy757/m3ntors/courses/pkg/config"
)

type Middleware struct {
	cfg *config.Config
}

type middlewareOptions func(*Middleware)

type userContextKey string

const userkey userContextKey = "user"

func WithConfig(cfg *config.Config) middlewareOptions {
	return func(m *Middleware) {
		m.cfg = cfg
	}
}

// NewMiddleware returns a new Midddleware type with options set
func New(opts ...middlewareOptions) *Middleware {
	m := &Middleware{}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

func (m *Middleware) SetUserContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		//Hack to remove the Bearer part of the header
		tokenString = strings.ReplaceAll(tokenString, "Bearer ", "")

		t, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// as we are just identifying the user and validity of the token
			// and errors or issues will just result in a anonymous user
			return []byte(m.cfg.AuthSecret), nil
		})
		if err == nil && t.Valid {
			if claims, ok := t.Claims.(jwt.MapClaims); ok {
				ctx := context.WithValue(r.Context(), userkey, claims["sub"])
				next.ServeHTTP(w, r.WithContext(ctx))
			}
		} else {
			next.ServeHTTP(w, r)
		}
	})
}
