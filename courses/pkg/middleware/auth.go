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

type User struct {
	ID string
}

type middlewareOptions func(*Middleware)

type contextKey string

const UserKey contextKey = "user"

// GetUserFromContext gets data from the context to return the user
func GetUserFromContext(ctx context.Context) *User {
	id := ctx.Value(UserKey)
	if id != nil {
		return &User{
			ID: id.(string),
		}
	}
	return nil
}

// WithConfig sets the configuration on the middleware
// struct
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

// SetUserContext is middleware to set the "user" context with teh sub (id)
// from a JWT token if it exists. It pushes the job of handeling authentication
// to the individual graphql resolvers
func (m *Middleware) SetUserContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		tokenString := r.Header.Get("Authorization")
		//Hack to remove the Bearer part of the header
		tokenString = strings.ReplaceAll(tokenString, "Bearer ", "")

		t, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// as we are just identifying the user and validity of the token
			// and errors or issues will just result in a anonymous user
			return []byte(m.cfg.AuthSecret), nil
		})

		if err == nil {
			ctx = m.addUserToContext(ctx, t)
		}

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// addUserToContext will add the user id to the context and return it
func (m *Middleware) addUserToContext(
	ctx context.Context,
	t *jwt.Token,
) context.Context {
	if t.Valid {
		if claims, ok := t.Claims.(jwt.MapClaims); ok {
			//TODO use the rest of the claims to create a user
			ctx = context.WithValue(ctx, UserKey, claims["sub"])
		}
	}
	return ctx
}
