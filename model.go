package main

import (
	"database/sql"
	"errors"
)

type product struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

// products/id/ --> CREATE
func (p *product) createProduct(db *sql.DB) error {
	return errors.New("Not implemented")
}

// products/id/ --> RETRIEVE
func (p *product) getProduct(db *sql.DB) error {
	return errors.New("Not implemented")
}

// products/id/ --> UPDATE
func (p *product) updateProduct(db *sql.DB) error {
	return errors.New("Not implemented")
}

// products/id/ --> DESTROY
func (p *product) deleteProduct(db *sql.DB) error {
	return errors.New("Not implemented")
}

// products/ --> LIST
func getProducts(db *sql.DB, start, count int) ([]product, error) {
	return nil, errors.New("Not implemented")
}
