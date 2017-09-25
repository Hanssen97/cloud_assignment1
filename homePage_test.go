package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/Hanssen97/cloud_assignment1/constants"
	"google.golang.org/appengine/aetest"
)

var inst aetest.Instance

func TestMain(m *testing.M) {
	inst, _ = aetest.NewInstance(nil)

	m.Run()

	defer tearDown()
}
func tearDown() {
	if inst != nil {
		inst.Close()
	}
}

func TestHomePage(t *testing.T) {
	req, err := inst.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal("Can't create request")
	}

	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(HomePage)

	handler.ServeHTTP(recorder, req)

	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := constants.HOMEPAGE

	if recorder.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			recorder.Body.String(), expected)
	}
}
func TestHomePageBadRequest(t *testing.T) {
	req, err := inst.NewRequest("GET", "/erybad", nil)
	if err != nil {
		t.Fatal("Can't create request")
	}

	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(HomePage)

	handler.ServeHTTP(recorder, req)

	if status := recorder.Code; status != http.StatusNotFound {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusNotFound)
	}

	expected := `404 page not found`

	if strings.TrimSpace(recorder.Body.String()) != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			recorder.Body.String(), expected)
	}
}
