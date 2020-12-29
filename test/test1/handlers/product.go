package handlers

import (
	"log"
	"net/http"
	"regexp"
	"strconv"

	"github.com/ayeressian/go-test2/test/test1/data"
)

type ProductController struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *ProductController {
	return &ProductController{l: l}
}

func (products *ProductController) ServeHTTP(respWriter http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case http.MethodGet:
		products.getProducts(respWriter, request)
	case http.MethodPost:
		products.addProduct(respWriter, request)
	case http.MethodPut:
		products.updateProduct(respWriter, request)
	default:
		respWriter.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (productController *ProductController) getProducts(respWriter http.ResponseWriter, request *http.Request) {
	lp := data.GetProducts()
	err := lp.ToJSON(respWriter)
	if err != nil {
		http.Error(respWriter, "Unable to parse product list", http.StatusInternalServerError)
	}
}

func (productController *ProductController) addProduct(respWriter http.ResponseWriter, request *http.Request) {
	product := &data.Product{}
	err := product.FromJSON(request.Body)
	if err != nil {
		http.Error(respWriter, "Incorrect product format", http.StatusBadRequest)
		return
	}
	data.AddProduct(product)
}

func (productController *ProductController) updateProduct(respWriter http.ResponseWriter, request *http.Request) {
	regex := regexp.MustCompile(`/([0-9]+)`)
	g := regex.FindAllStringSubmatch(request.URL.Path, -1)

	if len(g) != 1 && len(g[0]) != 1 {
		http.Error(respWriter, "invalid URL", http.StatusBadRequest)
		return
	}

	idString := g[0][1]
	id, _ := strconv.Atoi(idString)

	p := &data.Product{ID: id}

	p.FromJSON(request.Body)

	err := data.UpdateProduct(p)

	if err == data.NotFoundError {
		http.Error(respWriter, "Product with the given ID doesn't exist", http.StatusNotFound)
	}
}