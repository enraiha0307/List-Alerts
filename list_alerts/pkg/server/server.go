package server

import (
	"log"
	"net/http"

	"gerrit.ext.net.nokia.com/AANM/sandbox-team3/blr-intern-projects/Akanksha/list_alerts/api"

	//"gerrit.ext.net.nokia.com/AANM/sandbox-team3/blr-intern-projects/Akanksha/list_alerts/pkg/api"
	"github.com/gorilla/mux"
)

//Server Started msg

//Start() starts a server on localhost:8080
func Start() {
	router := mux.NewRouter().StrictSlash(true)
	api.Handler(router)

	log.Fatal(http.ListenAndServe("0.0.0.0:8080", router))
}
