package main

import (
	"net/http/httptest"
	"testing"
)

// import (
// 	"net/http"
// 	"net/http/httptest"
// 	"testing"
//
// 	"github.com/Hanssen97/cloud_assignment1/constants"
// 	"github.com/Hanssen97/cloud_assignment1/testdata"
// 	"github.com/jarcoal/httpmock"
// 	"google.golang.org/appengine/aetest"
// )
//
// var inst aetest.Instance
//
// func TestMain(m *testing.M) {
// 	inst, _ = aetest.NewInstance(nil)
//
// 	m.Run()
//
// 	defer tearDown()
// }
// func tearDown() {
// 	if inst != nil {
// 		inst.Close()
// 	}
// }
//
// // Repo() ----------------------------------------------------
// func TestRepo(t *testing.T) {
// 	httpmock.Activate()
// 	defer httpmock.DeactivateAndReset()
//
// 	httpmock.RegisterResponder("GET", constants.GITHUB+"repos/klyve/gritapp",
// 		httpmock.NewStringResponder(200, testdata.REPOJSON))
//
// 	httpmock.RegisterResponder("GET", constants.GITHUB+"repos/klyve/gritapp/contributors",
// 		httpmock.NewStringResponder(200, testdata.REPOCONTRIBUTORS))
//
// 	httpmock.RegisterResponder("GET", constants.GITHUB+"repos/klyve/gritapp/languages",
// 		httpmock.NewStringResponder(200, testdata.REPOLANGUAGES))
//
// 	req, err := inst.NewRequest("GET", "/projectinfo/v1/github.com/klyve/gritapp", nil)
// 	if err != nil {
// 		t.Fatal("Can't create request")
// 	}
//
// 	recorder := httptest.NewRecorder()
//
// 	handler := http.HandlerFunc(Repo)
// 	handler.ServeHTTP(recorder, req)
//
// 	expected := testdata.REPORES
//
// 	if recorder.Body.String() != expected {
// 		t.Errorf("handler returned unexpected body: got %v want %v",
// 			recorder.Body.String(), expected)
// 	}
// }
//
// // getRepoMapShortUrl() -----------------------------------------------
// func TestGetRepoMapShortURL(t *testing.T) {
// 	req, err := inst.NewRequest("GET", "/URLThatIsNotLongEnough", nil)
// 	if err != nil {
// 		t.Fatal("Can't create request")
// 	}
//
// 	recorder := httptest.NewRecorder()
//
// 	data, err := getRepoMap(recorder, req)
//
// 	expected := "Bad Request"
//
// 	if err.Error() != expected {
// 		t.Errorf("Function returned unexpected error: got %v want %v",
// 			data, expected)
// 	}
// }
//
// // Respond() --------------------------------------------------
func TestRespond(t *testing.T) {
	data := Repository{
		Project:   "github.com/test/test",
		Owner:     "Tester",
		Committer: "Committer",
		Commits:   110,
		Language:  []string{"JavaScript", "Go", "C++"},
	}

	recorder := httptest.NewRecorder()

	respond(recorder, data)

	expected := `{
 "project": "github.com/test/test",
 "owner": "Tester",
 "committer": "Committer",
 "commits": 110,
 "language": [
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
