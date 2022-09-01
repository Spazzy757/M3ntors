package middleware

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang-jwt/jwt/v4"
	"github.com/spazzy757/m3ntors/courses/pkg/config"
	"github.com/stretchr/testify/require"
)

func TestSetUserInContext(t *testing.T) {
	assert := require.New(t)
	cfg := &config.Config{
		AuthSecret: "test",
	}

	m := New(WithConfig(cfg))

	t.Run("authentication user", func(t *testing.T) {
		h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			val := r.Context().Value(userkey)
			assert.Equal(val, "foobar")
			w.WriteHeader(http.StatusCreated)
		})
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"sub": "foobar",
		})
		res := httptest.NewRecorder()
		tokenString, err := token.SignedString([]byte(cfg.AuthSecret))
		assert.NoError(err)
		mockHandler := m.SetUserContext(h)
		req := httptest.NewRequest("GET", "http://testing", nil)
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", tokenString))
		mockHandler.ServeHTTP(res, req)
		assert.Equal(res.Code, http.StatusCreated)
	})
	t.Run("unauthentication user", func(t *testing.T) {
		h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			val := r.Context().Value(userkey)
			assert.Nil(val)
			w.WriteHeader(http.StatusCreated)
		})
		res := httptest.NewRecorder()
		mockHandler := m.SetUserContext(h)
		req := httptest.NewRequest("GET", "http://testing", nil)
		mockHandler.ServeHTTP(res, req)
		assert.Equal(res.Code, http.StatusCreated)
	})
	t.Run("invalid token", func(t *testing.T) {
		h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			val := r.Context().Value(userkey)
			assert.Nil(val)
			w.WriteHeader(http.StatusCreated)
		})
		res := httptest.NewRecorder()
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"sub": "foobar",
		})
		tokenString, err := token.SignedString([]byte("not_the_secret"))
		assert.NoError(err)
		mockHandler := m.SetUserContext(h)
		req := httptest.NewRequest("GET", "http://testing", nil)
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", tokenString))
		mockHandler.ServeHTTP(res, req)
		assert.Equal(res.Code, http.StatusCreated)
	})

	t.Run("invalid claim", func(t *testing.T) {
		h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			val := r.Context().Value(userkey)
			assert.Nil(val)
			w.WriteHeader(http.StatusCreated)
		})
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"wrong": "foobar",
		})
		res := httptest.NewRecorder()
		tokenString, err := token.SignedString([]byte(cfg.AuthSecret))
		assert.NoError(err)
		mockHandler := m.SetUserContext(h)
		req := httptest.NewRequest("GET", "http://testing", nil)
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", tokenString))
		mockHandler.ServeHTTP(res, req)
		assert.Equal(res.Code, http.StatusCreated)
	})
}
