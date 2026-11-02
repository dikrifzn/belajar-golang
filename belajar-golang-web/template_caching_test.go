package belajar_golang_web

import (
	"embed"
	_ "embed"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

//go:embed templates/*.gohtml
var templateCache embed.FS
var cachedTemplates = template.Must(template.ParseFS(templateCache, "templates/*.gohtml"))

func TemplateCaching(w http.ResponseWriter, r *http.Request) {
	cachedTemplates.ExecuteTemplate(w, "simple.gohtml", "Hello World")
}

func TestTemplateCaching(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateCaching(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}
