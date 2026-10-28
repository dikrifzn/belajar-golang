package belajar_golang_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func RequestHandler(writer http.ResponseWriter, request *http.Request) {
	//logic handler
	contentType := request.Header.Get("Content-Type")
	fmt.Fprint(writer, contentType)
}

func TestHeader(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/hello", nil)

	request.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()

	RequestHandler(recorder, request)

	response := recorder.Result()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		t.Fatal(err)
	}

	bodyString := string(body)

	fmt.Println(bodyString)
}

func ResponseHeader(writer http.ResponseWriter, request *http.Request) {
	//logic handler
	writer.Header().Add("X-Powered-By", "Belajar Golang Web")
	fmt.Fprint(writer, "ok")
}

func TestResponseHeader(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/hello", nil)

	request.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()

	ResponseHeader(recorder, request)

	response := recorder.Result()

	fmt.Println("Response Header:", response.Header.Get("X-Powered-By"))
}