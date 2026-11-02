package belajar_golang_web

import (
	"bytes"
	_ "embed"
	"fmt"
	"html/template"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func UploadFile(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFiles("./templates/form.gohtml"))
	t.ExecuteTemplate(writer, "form.gohtml", nil)
}

func UploadFileProcess(writer http.ResponseWriter, request *http.Request) {
	// request.ParseMultipartForm(32 << 20)
	file, fileHeader, err := request.FormFile("file")
	if err != nil {
		panic(err)
	}

	fileDestination, err := os.Create("./resources/" + fileHeader.Filename)
	if err != nil {
		panic(err)
	}
	_, err = io.Copy(fileDestination, file)
	if err != nil {
		panic(err)
	}
	defer fileDestination.Close()

	name := request.PostFormValue("name")
	t := template.Must(template.ParseFiles("./templates/uploadSuccess.gohtml"))
	t.ExecuteTemplate(writer, "uploadSuccess.gohtml", map[string]interface{}{
		"Name": name,
		"File": "/static/" + fileHeader.Filename,
	})
}

// func TestUploadFileProcess(t *testing.T) {
// 	mux := http.NewServeMux()
// 	mux.HandleFunc("/form", UploadFile)
// 	mux.HandleFunc("/upload", UploadFileProcess)
// 	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./resources"))))

// 	server := http.Server{
// 		Addr:    "localhost:8080",
// 		Handler: mux,
// 	}

// 	err := server.ListenAndServe()
// 	if err != nil {
// 		panic(err)
// 	}
// }

//go:embed resources/testing.png
var uploadFileTest []byte

func TestUploadFile(t *testing.T) {
	body := &bytes.Buffer{}

	writer := multipart.NewWriter(body)

	writer.WriteField("name", "Dikri")
	file, _ := writer.CreateFormFile("file", "testing.png")

	file.Write(uploadFileTest)

	writer.Close()

	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/form", body)
	request.Header.Set("Content-Type", writer.FormDataContentType())
	recorder := httptest.NewRecorder()

	UploadFile(recorder, request)

	bodyResponse, _ := io.ReadAll(recorder.Result().Body)

	if !bytes.Contains(bodyResponse, []byte("My Form")) {
		t.Error("body response should contain 'My Form'")
	}
	fmt.Println(string(bodyResponse))
}
