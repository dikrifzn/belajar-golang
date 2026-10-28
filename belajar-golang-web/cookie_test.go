package belajar_golang_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func SetCookie(writer http.ResponseWriter, request *http.Request) {
	cookie := new(http.Cookie)
	cookie.Name = "X-PZN-Name"
	cookie.Value = request.URL.Query().Get("name")
	cookie.Path = "/"

	http.SetCookie(writer, cookie)
	fmt.Fprint(writer, "Success Set Cookie")

}

func GetCookie(writer http.ResponseWriter, request *http.Request) {
	cookie, err := request.Cookie("X-PZN-Name")
	if err != nil {
		fmt.Fprint(writer, "No Cookie")
	} else {
		fmt.Fprint(writer, "Hello ", cookie.Value)
	}
}

func TestSetCookie(t *testing.T) {
	//logic handler
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/?name=Dikri", nil)
	recorder := httptest.NewRecorder()

	SetCookie(recorder, request)

	cookie := recorder.Result().Cookies()
	for _, cookie := range cookie {
		fmt.Println(cookie.Name, " : ", cookie.Value)
	}
}

func TestGetCookie(t *testing.T) {
	//logic handler
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/", nil)
	cookie := new(http.Cookie)
	cookie.Name = "X-PZN-Name"
	cookie.Value = "Dikri"
	request.AddCookie(cookie)

	recorder := httptest.NewRecorder()

	GetCookie(recorder, request)

	body, err := io.ReadAll(recorder.Result().Body)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(string(body))
}
