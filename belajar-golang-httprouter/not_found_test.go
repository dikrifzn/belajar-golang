package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

func TestNotFoundHandler(t *testing.T) {
	router := httprouter.New()

	router.NotFound = http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "Custom Not Found")
	})

	request := httptest.NewRequest("GET", "http://localhost:8080/", nil)

	Recorder := httptest.NewRecorder()

	router.ServeHTTP(Recorder, request)

	response := Recorder.Result()
	body, _ := io.ReadAll(response.Body)

	assert.Equal(t, "Custom Not Found", string(body))

}
