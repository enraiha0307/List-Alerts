package api

import (
	"gerrit.ext.net.nokia.com/AANM/sandbox-team3/blr-intern-projects/Akanksha/list_alerts/pkg/api"
	"github.com/gorilla/mux"
)

func Handler(router *mux.Router) {

	// home route
	router.HandleFunc("/", api.HomeRoute)
	// getAlertsBySeverity
	router.HandleFunc("/{severity}", api.ListAlertsBySeverity)

}
