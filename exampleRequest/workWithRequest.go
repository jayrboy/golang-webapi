package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type Product struct {
	ID         int     `json:"id"`
	Name       string  `json:"name"`
	Price      float64 `json:"price"`
	Instructor string  `json:"instructor"`
}

var productList []Product

func init() {
	productJSON := `[
		{
			"id": 1,
			"name": "SQL",
			"price": 1970,
			"instructor": "IBM"
		},
		{
			"id": 2,
			"name": "Python",
			"price": 1991,
			"instructor": "CWI"
		},
		{
			"id": 3,
			"name": "Go",
			"price": 2009,
			"instructor": "Google"
		}
	]`
	err := json.Unmarshal([]byte(productJSON), &productList)

	if err != nil {
		log.Fatal(err)
	}
}

func findID(ID int) (*Product, int) {
	for i, product := range productList {
		if product.ID == ID {
			return &product, i
		}
	}
	return nil, 0
}

func productHandler(w http.ResponseWriter, r *http.Request) {
	urlPathSegment := strings.Split(r.URL.Path, "product/")
	ID, err := strconv.Atoi(urlPathSegment[len(urlPathSegment)-1])

	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	product, listItemIndex := findID(ID)
	if product == nil {
		http.Error(w, fmt.Sprintf("no product with id %d", ID), http.StatusNotFound)
		return
	}

	switch r.Method {
	case http.MethodGet:
		productJSON, err := json.Marshal(product)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(productJSON)

	case http.MethodPut:
		var updatedProduct Product
		byteBody, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		json.Unmarshal(byteBody, &updatedProduct)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if updatedProduct.ID != ID {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		product = &updatedProduct
		productList[listItemIndex] = *product
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		w.Write(byteBody)
		return

	case http.MethodDelete:
		productList = append(productList[:listItemIndex], productList[listItemIndex+1:]...)
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{"message":"Product deleted"}`)
		return

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)

	}
}

func getNextID() int {
	highestID := -1

	for _, product := range productList {
		if highestID < product.ID {
			highestID = product.ID
		}
	}
	return highestID + 1
}

func productsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		productJSON, err := json.Marshal(productList)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(productJSON)
	case http.MethodPost:
		var newProduct Product
		bodyBytes, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		err = json.Unmarshal(bodyBytes, &newProduct)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if newProduct.ID != 0 {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		newProduct.ID = getNextID()
		productList = append(productList, newProduct)
		w.WriteHeader(http.StatusCreated)
		w.Header().Set("Content-Type", "application/json")
		w.Write(bodyBytes)
		return
	}
}

func main() {
	port := "5000"

	http.HandleFunc("/product/", productHandler)
	http.HandleFunc("/product", productsHandler)
	http.ListenAndServe(":"+ port, nil)
}
