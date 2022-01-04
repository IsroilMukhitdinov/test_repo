package handlers

import (
	"net/http"

	"github.com/IsroilMukhitdinov/gomod_test/pkg/renders"
)

func Home(response http.ResponseWriter, request *http.Request) {
	renders.Render(response, "home.page.tmpl")

}

func About(response http.ResponseWriter, request *http.Request) {
	renders.Render(response, "about.page.tmpl")
}
