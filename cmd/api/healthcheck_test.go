package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealthcheckHandler(t *testing.T) {
	app := &application{
		config: configuration{
			env: "testing",
		},
	}

	req := httptest.NewRequest(http.MethodGet, "/v1/healthcheck", nil)
	rr := httptest.NewRecorder()

	app.healthcheckHandler(rr, req)

	// Check status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Check content type
	if contentType := rr.Header().Get("Content-Type"); contentType != "application/json" {
		t.Errorf("handler returned wrong content type: got %v want application/json", contentType)
	}

	// Parse response body
	var response map[string]any
	err := json.Unmarshal(rr.Body.Bytes(), &response)
	if err != nil {
		t.Fatalf("failed to parse response body: %v", err)
	}

	// Check for expected fields
	if status, ok := response["status"]; !ok {
		t.Error("response missing 'status' field")
	} else if status != "available" {
		t.Errorf("status = %v, want 'available'", status)
	}

	if systemInfo, ok := response["system_info"].(map[string]any); !ok {
		t.Error("response missing 'system_info' field")
	} else {
		if env, ok := systemInfo["environment"]; !ok {
			t.Error("system_info missing 'environment' field")
		} else if env != "testing" {
			t.Errorf("environment = %v, want 'testing'", env)
		}

		if _, ok := systemInfo["version"]; !ok {
			t.Error("system_info missing 'version' field")
		}
	}
}