package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/hanssen97/assignment1/constants"
	"github.com/hanssen97/assignment1/testdata"
	"github.com/jarcoal/httpmock"
	"google.golang.org/appengine/aetest"
)

var inst aetest.Instance

func TestMain(m *testing.M) {
	inst, _ = aetest.NewInstance(nil)

	m.Run()

	defer tearDown()
}

// Repo() ----------------------------------------------------
func TestRepo(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", constants.GITHUB+"repos/klyve/gritapp",
		httpmock.NewStringResponder(200, testdata.REPOJSON))

	httpmock.RegisterResponder("GET", constants.GITHUB+"repos/klyve/gritapp/contributors",
		httpmock.NewStringResponder(200, testdata.REPOCONTRIBUTORS))

	httpmock.RegisterResponder("GET", constants.GITHUB+"repos/klyve/gritapp/languages",
		httpmock.NewStringResponder(200, testdata.REPOLANGUAGES))

	req, err := inst.NewRequest("GET", "/projectinfo/v1/github.com/klyve/gritapp", nil)
	if err != nil {
		t.Fatal("Can't create request")
	}

	recorder := httptest.NewRecorder()

	handler := http.HandlerFunc(Repo)
	handler.ServeHTTP(recorder, req)

	expected := testdata.REPORES

	if recorder.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			recorder.Body.String(), expected)
	}
}

// getRepoMapShortUrl() -----------------------------------------------
func TestGetRepoMapShortURL(t *testing.T) {
	req, err := inst.NewRequest("GET", "/URLThatIsNotLongEnough", nil)
	if err != nil {
		t.Fatal("Can't create request")
	}

	recorder := httptest.NewRecorder()

	data, err := getRepoMap(recorder, req)

	expected := "Bad Request"

	if err.Error() != expected {
		t.Errorf("Function returned unexpected error: got %v want %v",
			data, expected)
	}
}

// getCommitters() --------------------------------------------
// func TestGetCommiters(t *testing.T) {
// 	httpmock.Activate()
// 	defer httpmock.DeactivateAndReset()
//
// 	httpmock.RegisterResponder("GET", "/getTestCommitters",
// 		httpmock.NewStringResponder(200, testdata.COMMITTERS))
//
// 	req, err := inst.NewRequest("GET", "/getTestCommitters", nil)
// 	if err != nil {
// 		t.Fatal("Can't create request")
// 	}
//
// 	recorder := httptest.NewRecorder()
//
// 	data, err := getCommitters(recorder, req, "/getTestCommitters")
//
// 	expected := []Committer{
// 		{
// 			Name:          "klyve",
// 			Contributions: 276,
// 		},
// 		{
// 			Name:          "Hanssen97",
// 			Contributions: 104,
// 		},
// 		{
// 			Name:          "omeyjey",
// 			Contributions: 57,
// 		},
// 		{
// 			Name:          "henriktre",
// 			Contributions: 12,
// 		},
// 	}
//
// 	if err != nil {
// 		t.Errorf("Unexpected error: got %v want %v",
// 			err, nil)
// 	}
//
// 	if !reflect.DeepEqual(data, expected) {
// 		t.Errorf("Function returned unexpected values: got %v want %v",
// 			data, expected)
// 	}
// }

// getLanguages() ---------------------------------------------
// func TestGetLanguages(t *testing.T) {
// 	httpmock.Activate()
// 	defer httpmock.DeactivateAndReset()
//
// 	httpmock.RegisterResponder("GET", "/getTestLanguages",
// 		httpmock.NewStringResponder(200, `{"JavaScript":110,"C++":69}`))
//
// 	req, err := inst.NewRequest("GET", "/getTestLanguages", nil)
// 	if err != nil {
// 		t.Fatal("Can't create request")
// 	}
//
// 	recorder := httptest.NewRecorder()
//
// 	data, err := getLanguages(recorder, req, "/getTestLanguages")
//
// 	expected := []string{"JavaScript", "C++"}
//
// 	if !reflect.DeepEqual(data, expected) {
// 		t.Errorf("Function returned unexpected values: got %v want %v",
// 			data, expected)
// 	}
// }

// Respond() --------------------------------------------------
func TestRespond(t *testing.T) {
	data := Repository{
		Project:   "github.com/test/test",
		Owner:     "Tester",
		Committer: "Committer",
		Commits:   110,
		Languages: []string{"JavaScript", "Go", "C++"},
	}

	recorder := httptest.NewRecorder()

	respond(recorder, data)

	expected := `{
 "project": "github.com/test/test",
 "owner": "Tester",
 "committer": "Committer",
 "commits": 110,
 "languages": [
  "JavaScript",
  "Go",
  "C++"
 ]
}`

	if recorder.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			recorder.Body.String(), expected)
	}
}

func tearDown() {
	if inst != nil {
		inst.Close()
	}
}
