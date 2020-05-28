package main

type routesDefiner struct {
	Name	string
	Port	int
}

var routesTemplate = &Routes {
	`package main

	import (
		"fmt"
		"net/http"
		"os"

		"github.com/gorilla/handlers"
	)

	func (s *server) routes() {
		s.router.Handle("/version", handlers.LoggingHandler(os.Stdout, s.handleVersion()))
	}

	// handleVersion supports GET requests and returns a json object of
	// {"version":"<version>"}
	func (s *server) handleVersion() http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			if r.Method == http.MethodGet {
				resp := fmt.Sprintf("{\"version\":\"%s\"}", version)
				w.WriteHeader(http.StatusOK)
				w.Header().Add("Content-Type", "application/json")
				w.Write([]byte(resp))
			}
		}
	}`,
}