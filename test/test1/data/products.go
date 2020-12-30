package data

import (
	"encoding/json"
	"errors"
	"io"
	"time"
)

type Product struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	CreatedAt string `json:"-"`
	UpdatedAt string `json:"-"`
}

type Products []*Product

func (products *Products) ToJSON(rw io.Writer) error {
	encoder := json.NewEncoder(rw)
	return encoder.Encode(products)
}

func (product *Product) FromJSON(reader io.Reader) error {
	decoder := json.NewDecoder(reader)
	return decoder.Decode(&product)
}

var productList = Products{
	&Product{
		ID:        1,
		Name:      "shampoo",
		CreatedAt: time.Now().UTC().String(),
		UpdatedAt: time.Now().UTC().String(),
	},
	&Product{
		ID:        2,
		Name:      "ghars",
		CreatedAt: time.Now().UTC().String(),
		UpdatedAt: time.Now().UTC().String(),
	},
}

func GetProducts() Products {
	return productList
}

func AddProduct(p *Product) {
	p.ID = getNextId()
	productList = append(productList, p)
}

var NotFoundError = errors.New("No prodcut exist with the given ID")

func UpdateProduct(pu *Product) error {
	for i, p := range productList {
		if p.ID == pu.ID {
			productList[i] = pu
			return nil
		}
	}
	return NotFoundError
}

func getNextId() int {
	return len(productList) + 1
}
