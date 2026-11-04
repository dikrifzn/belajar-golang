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

func TestNotAllowed(t *testing.T) {
	router := httprouter.New()

	router.MethodNotAllowed = http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "Custom Not Allowed")
	})

	router.GET("/", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		fmt.Fprint(writer, "Hello World")
	})
	
	request := httptest.NewRequest("POST", "http://localhost:8080/", nil)

	Recorder := httptest.NewRecorder()

	router.ServeHTTP(Recorder, request)

	response := Recorder.Result()
	body, _ := io.ReadAll(response.Body)

	assert.Equal(t, "Custom Not Allowed", string(body))
}
