package api

type Label struct {
	AlertArea                      string `json:"alert_area,omitempty"`
	AlertDelay                     string `json:"alert_delay,omitempty"`
	AlertDisplayname               string `json:"alert_display_name,omitempty"`
	AlertID                        string `json:"alert_id,omitempty"`
	AlertInstance                  string `json:"alert_instance,omitempty"`
	AlertMappingID                 string `json:"alert_mapping_id,omitempty"`
	AlertMappingIDType             string `json:"alert_mapping_id_type,omitempty"`
	AlertMappingSymptomSourceDt    string `json:"alert_mapping_symptom_source_dt,omitempty"`
	AlertMappingSymptomSourceIs    string `json:"alert_mapping_symptom_source_is,omitempty"`
	AlertMappingSymptomSourceProm  string `json:"alert_mapping_symptom_source_prom,omitempty"`
	AlertMappingSymptomStartOffset string `json:"alert_mapping_symptom_start_offset,omitempty"`
	AlertProbableCause             string `json:"alert_probable_cause,omitempty"`
	AlertPublic                    string `json:"alert_public,omitempty"`
	AlertResolution                string `json:"alert_resolution,omitempty"`
	AlertService                   string `json:"alert_service,omitempty"`
	AlertSeverity                  string `json:"alert_severity,omitempty"`
	AlertText                      string `json:"alert_text,omitempty"`
	AlertThreshold                 string `json:"alert_threshold,omitempty"`
	AlertThresholdUnit             string `json:"alert_threshold_unit,omitempty"`
	AlertType                      string `json:"alert_type,omitempty"`
	AlertName                      string `json:"alertname,omitempty"`
	KubernetesIOHostname           string `json:"kubernetes_io_hostname,omitempty"`
	KubernetesName                 string `json:"kubernetes_name,omitempty"`
	KubernetesPodName              string `json:"kubernetes_pod_name,omitempty"`
	Severity                       string `json:"severity,omitempty"`
}

type Annotation struct {
	Description string `json:"description,omitempty"`
	Summary     string `json:"summary,omitempty"`
}

type Alert struct {
	Labels      Label      `json:"labels,omitempty"`
	Annotations Annotation `json:"annotations,omitempty"`
	State       string     `json:"state,omitempty"`
	ActiveAt    string     `json:"activeAt,omitempty"`
	Value       string     `json:"value,omitempty"`
}

type Data struct {
	Alerts []Alert `json:"alerts,omitempty"`
}

type AlertRequest struct {
	Status string `json:"status,omitempty"`
	Data   Data   `json:"data,omitempty"`
}
