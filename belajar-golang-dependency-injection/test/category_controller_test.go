package test

import (
	"belajar-golang-dependency-injection/app"
	"belajar-golang-dependency-injection/controller"
	"belajar-golang-dependency-injection/helper"
	"belajar-golang-dependency-injection/middleware"
	"belajar-golang-dependency-injection/model/domain"
	"belajar-golang-dependency-injection/repository"
	"belajar-golang-dependency-injection/service"
	"context"
	"database/sql"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
)

func setupTestDB() *sql.DB {
	db, err := sql.Open("mysql", "root:1234@tcp(localhost:3306)/belajar_golang_rest_api_test")
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(60 * time.Minute)

	return db
}

func setupRouter(db *sql.DB) http.Handler {
	validate := validator.New()

	categoryRepository := repository.NewCategoryRepositoryImpl()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)

	router := app.NewRouter(categoryController)
	return middleware.NewAuthMiddleware(router)
}

func truncateCategory(db *sql.DB) {
	db.Exec("TRUNCATE category")
}

func TestCreateCategorySuccess(t *testing.T) {
	db := setupTestDB()
	truncateCategory(db)
	router := setupRouter(db)

	requestBody := strings.NewReader(`{"name": "Gadget"}`)

	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/api/categories", requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "SECRET")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	assert.Equal(t, 200, response.StatusCode)

	var responseBody map[string]interface{}
	body, _ := io.ReadAll(response.Body)
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	assert.Equal(t, "OK", responseBody["status"])
	assert.Equal(t, "Gadget", responseBody["data"].(map[string]interface{})["Name"])
}
func TestCreateCategoryFailed(t *testing.T) {
	db := setupTestDB()
	truncateCategory(db)
	router := setupRouter(db)

	requestBody := strings.NewReader(`{"name": ""}`)

	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/api/categories", requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "SECRET")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	assert.Equal(t, 400, response.StatusCode)

	var responseBody map[string]interface{}
	body, _ := io.ReadAll(response.Body)
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 400, int(responseBody["code"].(float64)))
	assert.Equal(t, "BAD REQUEST", responseBody["status"])
}
func TestUpdateCategorySuccess(t *testing.T) {
	db := setupTestDB()
	truncateCategory(db)

	tx, _ := db.Begin()
	categoryRepository := repository.NewCategoryRepositoryImpl()
	category := categoryRepository.Create(context.Background(), tx, domain.Category{
		Name: "Gadget",
	})
	tx.Commit()

	router := setupRouter(db)

	requestBody := strings.NewReader(`{"name": "Gadget"}`)

	request := httptest.NewRequest(http.MethodPut, "http://localhost:8080/api/categories/"+strconv.Itoa(category.Id), requestBody)

	request.Header.Add("X-API-Key", "SECRET")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	assert.Equal(t, 200, response.StatusCode)

	var responseBody map[string]interface{}
	body, _ := io.ReadAll(response.Body)
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	assert.Equal(t, "OK", responseBody["status"])
	assert.Equal(t, category.Id, int(responseBody["data"].(map[string]interface{})["Id"].(float64)))
	assert.Equal(t, "Gadget", responseBody["data"].(map[string]interface{})["Name"])
}
func TestUpdateCategoryFailed(t *testing.T) {
	db := setupTestDB()
	truncateCategory(db)

	tx, _ := db.Begin()
	categoryRepository := repository.NewCategoryRepositoryImpl()
	category := categoryRepository.Create(context.Background(), tx, domain.Category{
		Name: "Gadget",
	})
	tx.Commit()

	router := setupRouter(db)

	requestBody := strings.NewReader(`{"name": ""}`)

	request := httptest.NewRequest(http.MethodPut, "http://localhost:8080/api/categories/"+strconv.Itoa(category.Id), requestBody)

	request.Header.Add("X-API-Key", "SECRET")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	assert.Equal(t, 400, response.StatusCode)

	var responseBody map[string]interface{}
	body, _ := io.ReadAll(response.Body)
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 400, int(responseBody["code"].(float64)))
	assert.Equal(t, "BAD REQUEST", responseBody["status"])
}
func TestGetCategorySuccess(t *testing.T) {
	db := setupTestDB()
	truncateCategory(db)

	tx, _ := db.Begin()
	categoryRepository := repository.NewCategoryRepositoryImpl()
	category := categoryRepository.Create(context.Background(), tx, domain.Category{
		Name: "Gadget",
	})
	tx.Commit()

	router := setupRouter(db)

	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/api/categories/"+strconv.Itoa(category.Id), nil)

	request.Header.Add("X-API-Key", "SECRET")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	assert.Equal(t, 200, response.StatusCode)

	var responseBody map[string]interface{}
	body, _ := io.ReadAll(response.Body)
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	assert.Equal(t, "OK", responseBody["status"])
	assert.Equal(t, category.Id, int(responseBody["data"].(map[string]interface{})["Id"].(float64)))
	assert.Equal(t, category.Name, responseBody["data"].(map[string]interface{})["Name"])
}
func TestGetCategoryfailed(t *testing.T) {
	db := setupTestDB()
	truncateCategory(db)

	router := setupRouter(db)

	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/api/categories/404", nil)

	request.Header.Add("X-API-Key", "SECRET")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	assert.Equal(t, 404, response.StatusCode)

	var responseBody map[string]interface{}
	body, _ := io.ReadAll(response.Body)
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 404, int(responseBody["code"].(float64)))
	assert.Equal(t, "NOT FOUND", responseBody["status"])
}
func TestDeleteCategorySuccess(t *testing.T) {
	db := setupTestDB()
	truncateCategory(db)

	tx, _ := db.Begin()
	categoryRepository := repository.NewCategoryRepositoryImpl()
	category := categoryRepository.Create(context.Background(), tx, domain.Category{
		Name: "Gadget",
	})
	tx.Commit()

	router := setupRouter(db)

	request := httptest.NewRequest(http.MethodDelete, "http://localhost:8080/api/categories/"+strconv.Itoa(category.Id), nil)

	request.Header.Add("X-API-Key", "SECRET")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	assert.Equal(t, 200, response.StatusCode)

	var responseBody map[string]interface{}
	body, _ := io.ReadAll(response.Body)
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	assert.Equal(t, "OK", responseBody["status"])
}
func TestDeleteCategoryFailed(t *testing.T) {
	db := setupTestDB()
	truncateCategory(db)

	router := setupRouter(db)

	request := httptest.NewRequest(http.MethodDelete, "http://localhost:8080/api/categories/404", nil)

	request.Header.Add("X-API-Key", "SECRET")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	assert.Equal(t, 404, response.StatusCode)

	var responseBody map[string]interface{}
	body, _ := io.ReadAll(response.Body)
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 404, int(responseBody["code"].(float64)))
	assert.Equal(t, "NOT FOUND", responseBody["status"])
}
func TestListCategory(t *testing.T) {
	db := setupTestDB()
	truncateCategory(db)

	tx, _ := db.Begin()
	categoryRepository := repository.NewCategoryRepositoryImpl()
	category1 := categoryRepository.Create(context.Background(), tx, domain.Category{
		Name: "Gadget",
	})
	category2 := categoryRepository.Create(context.Background(), tx, domain.Category{
		Name: "Gadget",
	})
	tx.Commit()

	router := setupRouter(db)

	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/api/categories", nil)

	request.Header.Add("X-API-Key", "SECRET")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	assert.Equal(t, 200, response.StatusCode)

	var responseBody map[string]interface{}
	body, _ := io.ReadAll(response.Body)
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	assert.Equal(t, "OK", responseBody["status"])

	categories := responseBody["data"].([]interface{})

	category1Map := categories[0].(map[string]interface{})

	assert.Equal(t, category1.Id, int(category1Map["Id"].(float64)))
	assert.Equal(t, category1.Name, category1Map["Name"].(string))

	category2Map := categories[1].(map[string]interface{})

	assert.Equal(t, category2.Id, int(category2Map["Id"].(float64)))
	assert.Equal(t, category2.Name, category2Map["Name"].(string))
}

func TestUnauthorizedCategory(t *testing.T) {
	db := setupTestDB()
	truncateCategory(db)

	router := setupRouter(db)

	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/api/categories", nil)

	request.Header.Add("X-API-Key", "WRONG_KEY")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	assert.Equal(t, 401, response.StatusCode)

	var responseBody map[string]interface{}
	body, _ := io.ReadAll(response.Body)
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 401, int(responseBody["code"].(float64)))
	assert.Equal(t, "UNAUTHORIZED", responseBody["status"])
}
