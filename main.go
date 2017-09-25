package main

import (
	"fmt"
	"net/http"

	"github.com/Hanssen97/cloud_assignment1/constants"

	"google.golang.org/appengine"
)

func main() {
	http.HandleFunc("/", HomePage)
	http.HandleFunc("/projectinfo/v1/github.com/", Repo)

	fmt.Println("App running on port " + constants.PORT)
	http.ListenAndServe(":"+constants.PORT, nil)

	appengine.Main()
}
