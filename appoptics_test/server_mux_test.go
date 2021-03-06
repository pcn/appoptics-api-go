package appoptics_test

import (
	"fmt"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/appoptics/appoptics-api-go"
	"github.com/gorilla/mux"
)

var (
	client *appoptics.Client
	server *httptest.Server
)

func setup() {
	router := NewServerTestMux()
	server = httptest.NewServer(router)
	serverURLWithVersion := fmt.Sprintf("%s/v1/", server.URL)
	client = appoptics.NewClient("deadbeef", appoptics.BaseURLClientOption(serverURLWithVersion))
}

func teardown() {
	server.Close()
}

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	teardown()
	os.Exit(code)
}

func NewServerTestMux() *mux.Router {
	router := mux.NewRouter()

	// Measurements

	// Metrics
	router.Handle("/v1/metrics", ListMetricsHandler()).Methods("GET")
	router.Handle("/v1/metrics/{name}", CreateMetricHandler()).Methods("PUT")
	router.Handle("/v1/metrics/{name}", RetrieveMetricHandler()).Methods("GET")
	router.Handle("/v1/metrics", UpdateMetricHandler()).Methods("PUT")
	router.Handle("/v1/metrics/", DeleteMetricHandler()).Methods("DELETE")

	// Spaces
	router.Handle("/v1/spaces", ListSpacesHandler()).Methods("GET")
	router.Handle("/v1/spaces", CreateSpaceHandler()).Methods("POST")
	router.Handle("/v1/spaces/{id}", RetrieveSpaceHandler()).Methods("GET")
	router.Handle("/v1/spaces/{id}", UpdateSpaceHandler()).Methods("PUT")
	router.Handle("/v1/spaces/{id}", DeleteSpaceHandler()).Methods("DELETE")

	// Charts
	router.Handle("/v1/spaces/{spaceId}/charts", ListChartsHandler()).Methods("GET")
	router.Handle("/v1/spaces/{spaceId}/charts", CreateChartHandler()).Methods("POST")
	router.Handle("/v1/spaces/{spaceId}/charts/{chartId}", RetrieveChartHandler()).Methods("GET")
	router.Handle("/v1/spaces/{spaceId}/charts/{chartId}", UpdateChartHandler()).Methods("PUT")
	router.Handle("/v1/spaces/{spaceId}/charts/{chartId}", DeleteChartHandler()).Methods("DELETE")

	// Services
	router.Handle("/v1/services", ListServicesHandler()).Methods("GET")
	router.Handle("/v1/services", CreateServiceHandler()).Methods("POST")
	router.Handle("/v1/services/{serviceId}", RetrieveServiceHandler()).Methods("GET")
	router.Handle("/v1/services/{serviceId}", UpdateServiceHandler()).Methods("PUT")
	router.Handle("/v1/services/{serviceId}", DeleteServiceHandler()).Methods("DELETE")

	// Annotations
	router.Handle("/v1/annotations", ListAnnotationsHandler()).Methods("GET")
	router.Handle("/v1/annotations/{streamName}", CreateAnnotationHandler()).Methods("POST")
	router.Handle("/v1/annotations/{streamName}", UpdateAnnotationStreamHandler()).Methods("PUT")
	router.Handle("/v1/annotations/{streamName}", RetrieveAnnotationsHandler()).Methods("GET")
	router.Handle("/v1/annotations/{streamName}/{eventID}", RetrieveAnnotationEventHandler()).Methods("GET")
	router.Handle("/v1/annotations/{streamName}/{eventID}/links", UpdateAnnotationEventHandler()).Methods("POST")
	router.Handle("/v1/annotations/{streamName}", DeleteAnnotationHandler()).Methods("DELETE")

	// Alerts
	router.Handle("/v1/alerts", ListAlertsHandler()).Methods("GET")
	router.Handle("/v1/alerts", CreateAlertHandler()).Methods("POST")
	router.Handle("/v1/alerts/{alertId}", RetrieveAlertHandler()).Methods("GET")
	router.Handle("/v1/alerts/{alertId}", UpdateAlertHandler()).Methods("PUT")
	router.Handle("/v1/alerts/{alertId}", DeleteAlertHandler()).Methods("DELETE")
	router.Handle("/v1/alerts/{alertId}/status", StatusAlertHandler()).Methods("GET")
	router.Handle("/v1/alerts/{alertId}/services", AssociateAlertHandler()).Methods("POST")
	router.Handle("/v1/alerts/{alertId}/services/{serviceId}", DisassociateAlertHandler()).Methods("DELETE")

	// API Tokens
	router.Handle("/v1/api_tokens", ListApiTokensHandler()).Methods("GET")
	router.Handle("/v1/api_tokens", CreateApiTokenHandler()).Methods("POST")
	router.Handle("/v1/api_tokens/{tokenName}", RetrieveApiTokenHandler()).Methods("GET")
	router.Handle("/v1/api_tokens/{tokenId}", UpdateApiTokenHandler()).Methods("PUT")
	router.Handle("/v1/api_tokens/{tokenId}", DeleteApiTokenHandler()).Methods("DELETE")

	// Jobs
	router.Handle("/v1/jobs/{jobID}", RetrieveJobsHandler()).Methods("GET")

	// Snapshots
	router.Handle("/v1/snapshots", CreateSnapshotHandler()).Methods("POST")
	router.Handle("/v1/snapshots/{snapshotID}", RetrieveSnapshotHandler()).Methods("GET")

	return router
}
