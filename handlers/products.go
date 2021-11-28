package handlers

import (
	"log"
	"net/http"

	"github.com/karol-janik/go-microservices/product-api/data"
)

type Products struct{
	l *log.Logger
}

func NewProdcuts(l*log.Logger) *Products {
	return &Products{l}
}

func (p* Products) ServeHTTP(rw http.ResponseWriter, r *http.Request){
	if r.Method == http.MethodGet {
		p.getProducts(rw, r)
		return
	}

	// catch all
	rw.WriteHeader(http.StatusMethodNotAllowed)
}

func (p *Products) getProducts(rw http.ResponseWriter, r *http.Request){
	lp := data.GetProducts()
	err := lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}