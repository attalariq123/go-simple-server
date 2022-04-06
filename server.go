package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

const port = ":8000"

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/", rootPage)
	router.HandleFunc("/products/{fetchCountPercentage}", products).Methods("GET")

	fmt.Println("Serving @ http://127.0.0.1" + port)
	log.Fatal(http.ListenAndServe(port, router))

}

func rootPage(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("This is root page"))
}

func products(w http.ResponseWriter, r *http.Request) {

	fetchCountPercentage, errInput := strconv.ParseFloat(mux.Vars(r)["fetchCountPercentage"], 64)

	fetchCount := 0

	if errInput != nil {
		fmt.Println(errInput.Error())
	} else {
		fetchCount = int(float64(len(productList)) * fetchCountPercentage / 100)
		if fetchCount > len(productList) {
			fetchCount = len(productList)
		}
	}

	// write to response
	jsonList, err := json.Marshal(productList[0:fetchCount])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

	} else {
		w.Header().Set("content-type", "application/json")
		w.Write(jsonList)
	}

}

type product struct {
	Name  string
	Price float64
	Qty   int
}

var productList = []product{

	{"Cabai", 5000.0, 30},
	{"Wortel", 2000.0, 10},
	{"Timun", 3000.0, 320},
	{"Bawang", 25000.0, 20},
	{"Bayam", 5000.0, 340},
	{"Beras", 110000.0, 300},
	{"Tomat", 8000.0, 230},
	{"Kentang", 6000.0, 120},
	{"Terong", 400.0, 10},
	{"Kol", 400.0, 20},
	{"Cabai", 5000.0, 30},
	{"Wortel", 2000.0, 10},
	{"Timun", 3000.0, 320},
	{"Bawang", 25000.0, 20},
	{"Bayam", 5000.0, 340},
	{"Beras", 110000.0, 300},
	{"Tomat", 8000.0, 230},
	{"Kentang", 6000.0, 120},
	{"Terong", 400.0, 10},
	{"Kol", 400.0, 20},
}
