package api

import "net/http"

//This is the home route handler
func HomeRoute(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to Home Page"))
}
