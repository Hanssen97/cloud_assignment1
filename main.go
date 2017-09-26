package main

import (
	"fmt"
	"net/http"

	"github.com/Hanssen97/cloud_assignment1/constants"
	"github.com/Hanssen97/cloud_assignment1/handlers"

	"google.golang.org/appengine"
)

func main() {
	http.HandleFunc("/", handlers.HomePage)
	http.HandleFunc("/projectinfo/v1/github.com/", handlers.Repo)

	fmt.Println("App running on port " + constants.PORT)
	http.ListenAndServe(":"+constants.PORT, nil)

	appengine.Main()
}
