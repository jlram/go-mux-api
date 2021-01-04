package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
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
	req, _ := http.NewRequest("GET", "product/1", nil)
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

	var m map[string]interface{} // empty interface, any type accepted
	json.Unmarshal(response.Body.Bytes(), &m)

	if m["name"] != "test product" {
		t.Errorf("Expected product name to be 'test product'. Got '%v'", m["name"])
	}

	if m["price"] != 11.22 {
		t.Errorf("Expected product price to be '11.22'. Got '%v'", m["price"])
	}

	// JSON unmarshal converts numbers to floats, this is the way we parse empty interfaces to int
	if m["id"].(int) != 1 {
		t.Errorf("Expected product ID to be '1'. Got '%v'", m["id"])
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
