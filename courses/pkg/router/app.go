package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/graphql-go/handler"
	"github.com/spazzy757/m3ntors/courses/pkg/config"
	"github.com/spazzy757/m3ntors/courses/pkg/graphql"
	"github.com/spazzy757/m3ntors/courses/pkg/middleware"
)

type App struct {
	Router *mux.Router
	Cfg    *config.Config
}

//GetRouter returns a mux router for server
func (a *App) GetRouter() {

	gh := handler.New(&handler.Config{
		Schema:   &graphql.GetGraphQLSetup(graphql.WithConfig(a.Cfg)).Schema,
		Pretty:   true,
		GraphiQL: false,
	})

	m := middleware.New(
		middleware.WithConfig(a.Cfg),
	)

	a.Router = mux.NewRouter()
	a.Router.HandleFunc("/healthz", a.HealthzHandler)
	a.Router.HandleFunc("/sandbox", a.SandboxHandler)
	a.Router.Handle("/graphql", m.SetUserContext(gh))
}

// HealthzHandler
func (a *App) HealthzHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

// SandboxHanlder returns a apollo graphql sandbox for local development
func (a *App) SandboxHandler(w http.ResponseWriter, r *http.Request) {
	var sandboxHTML = []byte(`
<!DOCTYPE html>
<html lang="en">
<body style="margin: 0; overflow-x: hidden; overflow-y: hidden">
<div id="sandbox" style="height:100vh; width:100vw;"></div>
<script src="https://embeddable-sandbox.cdn.apollographql.com/_latest/embeddable-sandbox.umd.production.min.js"></script>
<script>
new window.EmbeddedSandbox({
  target: "#sandbox",
  initialEndpoint: "http://localhost:8001/graphql",
});
</script>
</body>
 
</html>`)
	w.Write(sandboxHTML)
}
