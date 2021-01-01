package handlers

import (
	"context"
	"log"
	"net/http"
	"strconv"

	"github.com/ayeressian/go-test2/test/test1/data"
	"github.com/gorilla/mux"
)

type ProductController struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *ProductController {
	return &ProductController{l: l}
}

func (productController *ProductController) GetProducts(respWriter http.ResponseWriter, request *http.Request) {
	lp := data.GetProducts()
	err := lp.ToJSON(respWriter)
	if err != nil {
		http.Error(respWriter, "Unable to parse product list", http.StatusInternalServerError)
	}
}

func (productController *ProductController) AddProduct(respWriter http.ResponseWriter, request *http.Request) {
	product := request.Context().Value(KeyProduct{}).(*data.Product)
	data.AddProduct(product)
}

func (productController *ProductController) UpdateProduct(respWriter http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id, _ := strconv.Atoi(vars["id"])

	product := request.Context().Value(KeyProduct{}).(*data.Product)

	product.ID = id

	err := data.UpdateProduct(product)

	if err == data.NotFoundError {
		http.Error(respWriter, "Product with the given ID doesn't exist", http.StatusNotFound)
	}
}

type KeyProduct struct {}

func (productController *ProductController) MidProductValid(next http.Handler) http.Handler {
	return http.HandlerFunc(func(respWriter http.ResponseWriter, request *http.Request) {
		product := &data.Product{}
		err := product.FromJSON(request.Body)
		if err != nil {
			http.Error(respWriter, "Incorrect product format", http.StatusBadRequest)
			return
		}

		ctx := context.WithValue(request.Context(), KeyProduct{}, product)
		request = request.WithContext(ctx)

		next.ServeHTTP(respWriter, request)
	})
}