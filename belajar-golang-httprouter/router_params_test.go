package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"strings"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

func TestParamsNamedParams(t *testing.T) {
	router := httprouter.New()

	router.GET("/products/:id/items/:item", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		id := params.ByName("id")
		itemId := params.ByName("item")
		text := "Product " + id + " Item " + itemId
		fmt.Fprint(writer, text)
	})

	request := httptest.NewRequest("GET", "http://localhost:8080/products/1/items/1", nil)

	Recorder := httptest.NewRecorder()

	router.ServeHTTP(Recorder, request)

	response := Recorder.Result()
	body, _ := io.ReadAll(response.Body)

	assert.Equal(t, "Product 1 Item 1", string(body))

}

func TestRouterCatchAll(t *testing.T) {
	router := httprouter.New()

	router.GET("/images/*image", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		image := params.ByName("image")
		image = strings.TrimPrefix(image, "/")
		text := "Image :" + image
		fmt.Fprint(writer, text)
	})

	request := httptest.NewRequest("GET", "http://localhost:8080/images/example.jpg", nil)

	Recorder := httptest.NewRecorder()

	router.ServeHTTP(Recorder, request)

	response := Recorder.Result()
	body, _ := io.ReadAll(response.Body)

	assert.Equal(t, "Image :example.jpg", string(body))

}
