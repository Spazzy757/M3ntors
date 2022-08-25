package router

import (
	"net/http"

	"github.com/gorilla/mux"
)

type App struct {
	Router *mux.Router
}

//GetRouter returns a mux router for server
func (a *App) GetRouter() {
	a.Router = mux.NewRouter()

	a.Router.HandleFunc("/healthz", a.HealthzHandler)
}

// HealthzHandler
func (a *App) HealthzHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
