package main

import (
	"encoding/json"
	"fmt"
	"log"
	"testing"
)

func LogJson(data interface{}) {
	bytes, err := json.Marshal(data)
	if err != nil {
		log.Println("Error encoding JSON:", err)
		return
	}
	log.Println(string(bytes))
}

func TestLogJson(t *testing.T) {
	LogJson("Dikri")
	LogJson(123)
	LogJson(true)
	LogJson([]string{"Dikri", "Golang", "JSON"})
}

type Customer struct {
	FirstName  string
	MiddleName string
	LastName   string
	Age        int
	Married    bool
	Hobbies    []string
}

func TestJsonObject(t *testing.T) {
	customer := Customer{
		FirstName:  "Dikri",
		MiddleName: "Golang",
		LastName:   "JSON",
		Age:        20,
		Married:    false,
	}

	bytes, err := json.Marshal(customer)
	if err != nil {
		log.Println("Error encoding JSON:", err)
		return
	}
	fmt.Println(string(bytes))
}

func TestDecodeJson(t *testing.T) {
	jsonString := `{
	"FirstName":"Dikri",
	"MiddleName":"Golang",
	"LastName":"JSON",
	"Age":20,
	"Married":false,
	}`

	jsonBytes := []byte(jsonString)

	customer := &Customer{}
	err := json.Unmarshal(jsonBytes, customer)

	if err != nil {
		log.Println("Error decoding JSON:", err)
		return
	}

	fmt.Println(customer)
	fmt.Println(customer.FirstName)
	fmt.Println(customer.MiddleName)
	fmt.Println(customer.LastName)
	fmt.Println(customer.Age)
	fmt.Println(customer.Married)
}

func TestJsonArray(t *testing.T) {
	customer := Customer{
		FirstName:  "Dikri",
		MiddleName: "Golang",
		LastName:   "JSON",
		Age:        20,
		Married:    false,
		Hobbies: []string{
			"Gaming",
			"Programming",
			"Reading",
		},
	}

	bytes, err := json.Marshal(customer)
	if err != nil {
		log.Println("Error encoding JSON:", err)
		return
	}
	fmt.Println(string(bytes))
}

type Address struct {
	Street     string
	City       string
	PostalCode string
}

type CustomerWithAddress struct {
	FirstName  string
	MiddleName string
	LastName   string
	Age        int
	Married    bool
	Hobbies    []string
	Addresses  []Address
}

func TestJsonNestedObject(t *testing.T) {
	customer := CustomerWithAddress{
		FirstName:  "Dikri",
		MiddleName: "Golang",
		LastName:   "JSON",
		Age:        20,
		Married:    false,
		Hobbies: []string{
			"Gaming",
			"Programming",
			"Reading",
		},
		Addresses: []Address{
			{
				Street:     "Jalan Merdeka",
				City:       "Jakarta",
				PostalCode: "12345",
			},
		},
	}

	bytes, err := json.Marshal(customer)
	if err != nil {
		log.Println("Error encoding JSON:", err)
		return
	}
	fmt.Println(string(bytes))
}

func TestDecodeJsonNestedObject(t *testing.T) {
	jsonString := `{
	"FirstName":"Dikri",
	"MiddleName":"Golang",
	"LastName":"JSON",
	"Age":20,
	"Married":false,
	"Hobbies":["Gaming","Programming","Reading"],
	"Addresses":[{"Street":"Jalan Merdeka","City":"Jakarta","PostalCode":"12345"}]
	}`

	customer := &CustomerWithAddress{}
	err := json.Unmarshal([]byte(jsonString), customer)
	if err != nil {
		log.Println("Error decoding JSON:", err)
		return
	}
	fmt.Println(customer)
	fmt.Println(customer.FirstName)
	fmt.Println(customer.MiddleName)
	fmt.Println(customer.LastName)
	fmt.Println(customer.Age)
	fmt.Println(customer.Married)
	fmt.Println(customer.Hobbies)
	fmt.Println(customer.Addresses)
}

func TestOnlyJsonArray(t *testing.T) {
	jsonString := `[
	{
		"street":"Jalan Merdeka",
		"city":"Jakarta",
		"postalCode":"12345"
	},
	{
		"street":"Jalan Sudirman",
		"city":"Bandung",
		"postalCode":"54321"
	}
	]`

	addresses := []Address{}
	err := json.Unmarshal([]byte(jsonString), &addresses)
	if err != nil {
		log.Println("Error decoding JSON:", err)
		return
	}
	fmt.Println(addresses)
}

func TestOnlyJsonArrayEncode(t *testing.T) {
	addreses := []Address{
		{
			Street:     "Jalan Merdeka",
			City:       "Jakarta",
			PostalCode: "12345",
		},
		{
			Street:     "Jalan Sudirman",
			City:       "Bandung",
			PostalCode: "54321",
		},
	}

	bytes, err := json.Marshal(addreses)
	if err != nil {
		log.Println("Error encoding JSON:", err)
		return
	}
	fmt.Println(string(bytes))
}

type Product struct {
	Id string 		`json:"id"`
	Name  string 	`json:"name"`
	Price int64 	`json:"price"`
	ImageUrl string	`json:"image_url"` // unexported field, will not be marshaled
}

func TestJsonTag(t *testing.T) {
	product := Product{
		Id: "1",
		Name: "Product 1",
		Price: 10000,
		ImageUrl: "https://example.com/product1.jpg",
	}
	bytes, err := json.Marshal(product)
	if err != nil {
		log.Println("Error encoding JSON:", err)
		return
	}
	fmt.Println(string(bytes))
}

func TestDecodeJsonTag(t *testing.T) {
	jsonString := `{
	"id":"1",
	"name":"Product 1",
	"price":10000,
	"image_url":"https://example.com/product1.jpg"
	}`
	product := &Product{}
	err := json.Unmarshal([]byte(jsonString), product)
	if err != nil {
		log.Println("Error decoding JSON:", err)
		return
	}
	fmt.Println(product)
	fmt.Println(product.Id)
	fmt.Println(product.Name)
	fmt.Println(product.Price)
	fmt.Println(product.ImageUrl)
}

func TestMap(t *testing.T) {
	jsonstring := `{
	"id":"1",
	"name":"Product 1",
	"price":10000,
	"image_url":"https://example.com/product1.jpg"
	}`

	jsonbytes := []byte(jsonstring)

	var result map[string]interface{}

	json.Unmarshal(jsonbytes, &result)

	fmt.Println(result)
	fmt.Println(result["id"])
	fmt.Println(result["name"])
	fmt.Println(result["price"])
	fmt.Println(result["image_url"])
}

func TestMapEncode(t *testing.T) {
	product := map[string]interface{}{
		"id": "1",
		"name": "Product 1",
		"price": 10000,
		"image_url": "https://example.com/product1.jpg",
	}

	bytes, err := json.Marshal(product)
	if err != nil {
		log.Println("Error encoding JSON:", err)
		return
	}

	fmt.Println(string(bytes))
}