package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"sort"
	"strings"

	"github.com/hanssen97/assignment1/constants"
	"google.golang.org/appengine"
	"google.golang.org/appengine/urlfetch"
)

// Repository holds data to return
type Repository struct {
	Project   string   `json:"project"`
	Owner     string   `json:"owner"`
	Committer string   `json:"committer"`
	Commits   int      `json:"commits"`
	Languages []string `json:"languages"`
}

// Committer helper struct to parse comitter data
type Committer struct {
	Name          string `json:"login"`
	Contributions int    `json:"contributions"`
}

// Oups holds error data to return
type Oups struct {
	Error string `json:"error"`
}

// Repo sends parsed repositoryinfo
func Repo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	// Gets a map of repository
	repoMap, err := getRepoMap(w, r)
	if err != nil {
		respond(w, Oups{Error: err.Error()}) // sends error response
		return
	}

	// Sets basic props for Repository
	repo := Repository{
		Project: "github.com/" + repoMap["full_name"].(string),
		Owner:   repoMap["owner"].(map[string]interface{})["login"].(string),
	}

	// Gets committers and parses data
	committers, err := getCommitters(w, r, repoMap["contributors_url"].(string))
	if err != nil {
		respond(w, Oups{Error: err.Error()}) // sends error response
		return
	}
	for i := range committers {
		repo.Commits += committers[i].Contributions
	}
	repo.Committer = committers[0].Name

	// Gets languages
	repo.Languages, err = getLanguages(w, r, repoMap["languages_url"].(string))
	if err != nil {
		respond(w, Oups{Error: err.Error()}) // sends error response
		return
	}

	// Everything went well, sending response; statuscode 200 ok
	w.WriteHeader(200)
	respond(w, repo)
}

func getRepoMap(w http.ResponseWriter, r *http.Request) (map[string]interface{}, error) {
	var err error
	var repoMap map[string]interface{}
	path := strings.Split(r.URL.Path, "/")

	if len(path) < 6 {
		err = errors.New(http.StatusText(400))
		w.WriteHeader(400)
		return repoMap, err
	}

	path = path[4:]

	body, err := parseGetBody(w, r, constants.GITHUB+"repos/"+path[0]+"/"+path[1])
	if err != nil {
		return repoMap, err
	}

	err = json.Unmarshal(body, &repoMap)

	if val, ok := repoMap["message"]; ok {
		err = errors.New(val.(string))
	}

	return repoMap, err
}

func getCommitters(w http.ResponseWriter, r *http.Request, path string) ([]Committer, error) {
	var committers []Committer
	var err error
	var temp interface{}

	body, err := parseGetBody(w, r, path)
	if err != nil {
		return committers, err
	}

	if err = json.Unmarshal(body, &temp); err != nil {
		return committers, err
	}

	parsed, ok := temp.(map[string]interface{})
	if !ok {
		err = json.Unmarshal(body, &committers)
	} else {
		err = errors.New(parsed["message"].(string))
	}

	return committers, err
}

func getLanguages(w http.ResponseWriter, r *http.Request, path string) ([]string, error) {
	var temp map[string]interface{}
	var languages []string
	var err error

	body, err := parseGetBody(w, r, path)
	if err != nil {
		return languages, err
	}

	err = json.Unmarshal(body, &temp)

	if val, ok := temp["message"]; ok {
		err = errors.New(val.(string))
	}

	// Sort based on: http://ispycode.com/GO/Sorting/Sort-map-by-value
	hack := map[int]string{}
	hackkeys := []int{}
	for key, val := range temp {
		hack[int(val.(float64))] = key
		hackkeys = append(hackkeys, int(val.(float64)))
	}
	sort.Ints(hackkeys)

	if err == nil {
		for i := len(hackkeys) - 1; i >= 0; i-- {
			languages = append(languages, hack[hackkeys[i]])
		}
	}

	return languages, err
}

func parseGetBody(w http.ResponseWriter, r *http.Request, path string) ([]byte, error) {
	var body []byte
	var err error

	ctx := appengine.NewContext(r)
	client := urlfetch.Client(ctx)
	res, err := client.Get(path)
	if err != nil {
		return body, err
	}

	//Forwards Github API statuscode if error
	if res.StatusCode != 200 {
		w.WriteHeader(res.StatusCode)
	}

	body, err = ioutil.ReadAll(res.Body)
	res.Body.Close()

	return body, err
}

func respond(w http.ResponseWriter, object interface{}) {
	res, err := json.MarshalIndent(object, "", " ")
	if err != nil {
		// Should implement safeguard for recursive overflow
		respond(w, Oups{Error: err.Error()})
	} else {
		fmt.Fprint(w, string(res))
	}
}
