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

type LogMiddleware struct {
	http.Handler
}

func (middleware *LogMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("Request URL:", request.URL.Path)
	middleware.Handler.ServeHTTP(writer, request)
}

func TestLogMiddleware(t *testing.T) {
	router := httprouter.New()

	router.GET("/", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		fmt.Fprint(writer, "Hello World")
	})

	middleware := &LogMiddleware{router}

	request := httptest.NewRequest("GET", "http://localhost:8080/", nil)

	Recorder := httptest.NewRecorder()

	middleware.ServeHTTP(Recorder, request)

	response := Recorder.Result()
	body, _ := io.ReadAll(response.Body)

	assert.Equal(t, "Hello World", string(body))
}
