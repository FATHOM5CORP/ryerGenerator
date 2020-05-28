package main

type mainDefiner struct {
	Name	string
	Port	int
}

var mainTemplate = &Main {
	`package main

	import (
		"fmt"
		"log"
		"net/http"
	)
	
	// Version in semantic versioning format
	const version = "v0.0.1"
	
	// Environment variable for port connection
	const (
		Port = {{.Port}}
	)
	
	type server struct {
		router *http.ServeMux
	}
	
	func newServer() *server {
		s := &server{}
		s.router = http.DefaultServeMux
		s.routes()
		return s
	}
	
	func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
		s.router.ServeHTTP(w, r)
	}
	
	func main() {
		server := newServer()
		log.Printf("{{.Name}} service listening on localhost port %d\n", Port)
		Address := fmt.Sprintf(":%d", Port)
		if err := http.ListenAndServe(Address, server.router); err != nil {
			log.Printf("{{.Name}} failed with error: %v\n", err)
		}
	}`,
}