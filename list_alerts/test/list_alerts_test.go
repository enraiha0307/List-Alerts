package test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"gerrit.ext.net.nokia.com/AANM/sandbox-team3/blr-intern-projects/Akanksha/list_alerts/pkg/api"
	"github.com/gorilla/mux"
)

func TestListAlertsBySeverity(t *testing.T) {
	t.Parallel()

	req, _ := http.NewRequest("GET", "/MINOR", nil)
	rr := httptest.NewRecorder()

	vars := map[string]string{
		"severity": "MINOR",
	}

	req = mux.SetURLVars(req, vars)

	api.ListAlertsBySeverity(rr, req)

	expected := `{"status":"success","data":{"alerts":[{"labels":{"alert_area":"http_endpoints","alert_delay":"60","alert_display_name":"ServiceHttpEndpointServerErrors","alert_id":"1000","alert_instance":"neo0125node07/ne3s-plugin-deployment-7cfc65b54-ptwmp","alert_mapping_id":"ne3s-plugin-deployment-7cfc65b54-ptwmp","alert_mapping_id_type":"kubernetes.pod_name","alert_mapping_symptom_source_dt":"default","alert_mapping_symptom_source_is":"default","alert_mapping_symptom_source_prom":"default","alert_mapping_symptom_start_offset":"3600","alert_probable_cause":"343","alert_public":"false","alert_resolution":"TBD","alert_service":"ne3s-plugin-service","alert_severity":"MINOR","alert_text":"Service ne3s-plugin-service instance ne3s-plugin-deployment-7cfc65b54-ptwmp on node neo0125node07 has over 20% of endpoint UNKNOWN requests failed to server errors during last 10 minutes (status code 5xx, gateway errors are excluded).","alert_threshold":"20","alert_threshold_unit":"%","alert_type":"qualityOfService","alertname":"ServiceHttpEndpointServerErrorsMinor","kubernetes_io_hostname":"neo0125node07","kubernetes_name":"ne3s-plugin-service","kubernetes_pod_name":"ne3s-plugin-deployment-7cfc65b54-ptwmp","severity":"2"},"annotations":{"description":"Service ne3s-plugin-service instance ne3s-plugin-deployment-7cfc65b54-ptwmp on node neo0125node07 has over 20% (100%) of endpoint UNKNOWN requests failed to server errors during last 10 minutes (status code 5xx, gateway errors are excluded).","summary":"Service has over 20% of endpoint requests failed to server errors."},"state":"firing","activeAt":"2021-09-21T06:15:23.54002335Z","value":"1e+02"}]}}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body:\n\n\n got %v \n\n\n want %v",
			rr.Body.String(), expected)
	}
}
