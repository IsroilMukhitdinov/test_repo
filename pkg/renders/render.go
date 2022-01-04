package renders

import (
	"fmt"
	"html/template"
	"net/http"
)

func Render(response http.ResponseWriter, tmpl string) {
	rendered, err := template.ParseFiles("./templates/" + tmpl)

	if err != nil {
		fmt.Fprintf(response, "error parsing the template")
		return
	}

	rendered.Execute(response, nil)
}
