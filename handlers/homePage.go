package handlers

import (
	"fmt"
	"net/http"

	"github.com/Hanssen97/cloud_assignment1/constants"
)

// HomePage returns the homepage
func HomePage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	fmt.Fprint(w, constants.HOMEPAGE)
}
