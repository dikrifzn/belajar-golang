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

func TestPanicHandler(t *testing.T) {
	router := httprouter.New()

	router.PanicHandler = func(writer http.ResponseWriter, request *http.Request, err interface{}) {
		fmt.Fprint(writer, "Ups Error")
	}

	router.GET("/", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		panic("ups")
	})

	request := httptest.NewRequest("GET", "http://localhost:8080/", nil)

	Recorder := httptest.NewRecorder()

	router.ServeHTTP(Recorder, request)

	response := Recorder.Result()
	body, _ := io.ReadAll(response.Body)

	assert.Equal(t, "Ups Error", string(body))

}
