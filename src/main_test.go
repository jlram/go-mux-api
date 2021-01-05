package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"
	"testing"
)

var a App

func TestMain(m *testing.M) {

	a.Initialize(
		"openpg",
		"openpgpwd",
		"postgres")

	ensureTableExists()
	code := m.Run()
	clearTable()
	os.Exit(code)
}

// Expected results: 200 code and empty array
func TestEmptyTable(t *testing.T) {
	clearTable()

	req, _ := http.NewRequest("GET", "/products", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code) // Expected: 200

	if body := response.Body.String(); body != "[]" {
		t.Errorf("Expected an empty array. Got %s", body)
	}
}

// Expected results: 404 code and not found error message
func TestNonExistentProduct(t *testing.T) {
	clearTable()
	req, _ := http.NewRequest("GET", "/product/1", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusNotFound, response.Code) // Expected: 404

	var m map[string]string
	// https://golang.org/pkg/encoding/json/#Unmarshal
	json.Unmarshal(response.Body.Bytes(), &m)
	if m["error"] != "Product not found" { // Checks that the response contains an error with the message "Product not found"
		t.Errorf("Expected the 'error' key of the response to be set to 'Product not found'. Got '%s'", m["error"])
	}
}

// Expected results: 201 code, ID 1 and price 11.22
func TestCreateProduct(t *testing.T) {
	clearTable()
	var jsonStr = []byte(`{"name":"test product", "price": 11.22}`)
	req, _ := http.NewRequest("POST", "/product", bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	response := executeRequest(req)
	checkResponseCode(t, http.StatusCreated, response.Code) // Expected: 201

	var product map[string]interface{} // empty interface, any type accepted
	json.Unmarshal(response.Body.Bytes(), &product)

	if product["name"] != "test product" {
		t.Errorf("Expected product name to be 'test product'. Got '%v'", product["name"])
	}

	if product["price"] != 11.22 {
		t.Errorf("Expected product price to be '11.22'. Got '%v'", product["price"])
	}

	// JSON unmarshal converts numbers to floats, this is the way we parse empty interfaces to int
	if id, ok := product["id"].(int); id != 1 && ok == true {
		t.Errorf("Expected product ID to be '1'. Got '%v'", product["id"])
	}
}

// Expected results: 200 code, name: Product 1, id: 1
func TestRetrieveProduct(t *testing.T) {
	clearTable()
	addProducts(1)
	req, _ := http.NewRequest("GET", "/product/1", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusCreated, response.Code) // Expected: 201

	var product map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &product)

	if product["name"] != "Product 1" {
		t.Errorf("Expected the 'name' key of the response to be set to 'Product 1'. Got '%s'", product["name"])
	}

	if id, ok := product["id"].(int); id != 1 && ok == false {
		t.Errorf("Expected product ID to be '1'. Got '%v'", product["id"])
	}
}

// Expected results: 200 code, same id, different name and price
func TestUpdateProduct(t *testing.T) {
	clearTable()
	addProducts(1)

	req, _ := http.NewRequest("GET", "/product/1", nil)
	response := executeRequest(req)

	// gets product before updating it, for checking purposes
	var originalProduct map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &originalProduct)

	// updates product
	var jsonStr = []byte(`{"name":"Product", "price": 55.00}`)
	req, _ = http.NewRequest("PUT", "/product/1", bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	response = executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)
	var updatedProduct map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &updatedProduct)

	if updatedProduct["id"] != originalProduct["id"] {
		t.Errorf("Expected the id to remain the same (%v). Got %v", originalProduct["id"], updatedProduct["id"])
	}

	if updatedProduct["name"] == originalProduct["name"] {
		t.Errorf("Expected the name to change from '%v' to '%v'. Got '%v'", originalProduct["name"], updatedProduct["name"], updatedProduct["name"])
	}

	if updatedProduct["price"] == originalProduct["price"] {
		t.Errorf("Expected the price to change from '%v' to '%v'. Got '%v'", originalProduct["price"], updatedProduct["price"], updatedProduct["price"])
	}
}

// Expected results: 204 code, then 404 code.
func TestDeleteProduct(t *testing.T) {
	clearTable()
	addProducts(1)

	req, _ := http.NewRequest("DELETE", "/product/1", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusNoContent, response.Code) // Expected 204

	req, _ = http.NewRequest("GET", "/product/1", nil)
	response = executeRequest(req)

	checkResponseCode(t, http.StatusNotFound, response.Code) // Expected 404
}

// Receives number of products to add, then iterates for loop and adds them
func addProducts(count int) {
	if count < 1 {
		count = 1
	}

	for i := 1; i < count; i++ {
		a.DB.Exec("INSERT INTO products(name, price) VALUES($1, $2)", "Product "+strconv.Itoa(i), (i+1.0)*10)
	}
}

// Calls given request and returns its response, using httptest
func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	a.Router.ServeHTTP(rr, req)

	return rr
}

// Check response's response code and comopares with expected
func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}

// Remove every instance on our test database
func clearTable() {
	a.DB.Exec("DELETE FROM products")
	a.DB.Exec("ALTER SEQUENCE products_id_seq RESTART WITH 1")
}

// Check if our table is created
func ensureTableExists() {
	if _, err := a.DB.Exec(tableCreationQuery); err != nil {
		log.Fatal(err)
	}
}

// Create product table, if not exists
const tableCreationQuery = `CREATE TABLE IF NOT EXISTS products
(
    id SERIAL,
    name TEXT NOT NULL,
    price NUMERIC(10,2) NOT NULL DEFAULT 0.00,
    CONSTRAINT products_pkey PRIMARY KEY (id)
)`
