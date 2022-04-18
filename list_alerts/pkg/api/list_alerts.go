package api

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//Base url for alertmanager api
//const BaseURL = "http://localhost:9093/api/v1/"

//HAndler function for getting Alerts
func ListAlertsBySeverity(w http.ResponseWriter, r *http.Request) {

	/*res, err := http.Get(BaseURL + "alerts")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Print("Failed to request alertmanager api")
	}
	fmt.Println(res)
	*/

	vars := mux.Vars(r)
	severity := vars["severity"]

	body, err := ioutil.ReadFile("/go/src/gerrit.ext.net.nokia.com/AANM/sandbox-team3/blr-intern-projects/Akanksha/list_alerts/pkg/api/alerts.json")
	//bytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)

	}

	allAlerts := AlertRequest{}

	err = json.Unmarshal(body, &allAlerts)

	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println(allAlerts.Data.Alerts)
	alertsBySeverity := AlertRequest{}
	var alert Alert
	alertsBySeverity.Status = allAlerts.Status

	for _, alert = range allAlerts.Data.Alerts {
		if alert.Labels.AlertSeverity == severity {

			alertsBySeverity.Data.Alerts = append(alertsBySeverity.Data.Alerts, alert)
		}

	}
	//	fmt.Println(alertsBySeverity)
	w.Header().Set("Content-Type", "application/json")
	res, err := json.Marshal(alertsBySeverity)
	if err != nil {
		log.Fatal(err)
		return
	}
	w.Write(res)
}
