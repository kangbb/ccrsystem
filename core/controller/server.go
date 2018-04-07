package controller

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

// GetServer return web server
func GetServer() *negroni.Negroni {
	r := mux.NewRouter()
	r.HandleFunc("/data", dataProcess)
	static := "static"
	r.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir(static))))

	s := negroni.Classic()
	s.UseHandler(r)
	return s
}
