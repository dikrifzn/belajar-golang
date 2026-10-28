package belajar_golang_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func SayHello(writer http.ResponseWriter, request *http.Request) {
	name := request.URL.Query().Get("name")
	if name == "" {
		fmt.Fprint(writer, "Hello!")
	} else {
		fmt.Fprintf(writer, "Hello %s!", name)
	}
}

func TestQueryParam(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/hello", nil)
	recorder := httptest.NewRecorder()

	SayHello(recorder, request)

	response := recorder.Result()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		t.Fatal(err)
	}
	bodyString := string(body)
	fmt.Println(bodyString)
}

func MultiQueryParam(writer http.ResponseWriter, request *http.Request) {
	query := request.URL.Query()
	firstName := query.Get("first_name")
	lastName := query.Get("last_name")

	if firstName == "" || lastName == "" {
		fmt.Fprint(writer, "Hello!")
	} else {
		fmt.Fprintf(writer, "Hello %s %s!", firstName, lastName)
	}
}

func TestMultiQueryParam(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/hello?first_name=Dikri&last_name=Pratama", nil)
	recorder := httptest.NewRecorder()
	MultiQueryParam(recorder, request)

	response := recorder.Result()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		t.Fatal(err)
	}
	bodyString := string(body)
	fmt.Println(bodyString)
}

func MultiQueryParamValue(writer http.ResponseWriter, request *http.Request) {
	query := request.URL.Query()
	names := query["name"]
	fmt.Fprint(writer, "Hello "+strings.Join(names, " "))
}

func TestMultiQueryParamValue(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/hello?name=Dikri&name=Pratama&name=Golang", nil)
	recorder := httptest.NewRecorder()
	MultiQueryParamValue(recorder, request)

	response := recorder.Result()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		t.Fatal(err)
	}
	bodyString := string(body)
	fmt.Println(bodyString)
}
