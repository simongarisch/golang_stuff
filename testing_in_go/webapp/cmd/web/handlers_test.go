package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_application_handlers(t *testing.T) {
	var expectedStatusCodes = []struct {
		name               string
		url                string
		expectedStatusCode int
	}{
		{"home", "/", http.StatusOK},
		{"404", "/fish", http.StatusNotFound},
	}

	routes := app.routes()

	// create a test server
	ts := httptest.NewTLSServer(routes)
	defer ts.Close()

	pathToTemplates = "./../../templates/"

	// range through the test data
	for _, e := range expectedStatusCodes {
		resp, err := ts.Client().Get(ts.URL + e.url)
		if err != nil {
			t.Log("err")
			t.Fatal(err)
		}

		if resp.StatusCode != e.expectedStatusCode {
			t.Errorf("for %s: expected status %d, got %d", e.name, e.expectedStatusCode, resp.StatusCode)
		}
	}
}
