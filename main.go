package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var urls = []string{
	"https://run.mocky.io/v3/c51441de-5c1a-4dc2-a44e-aab4f619926b",
	"https://run.mocky.io/v3/4ec58fbc-e9e5-4ace-9ff0-4e893ef9663c",
	"https://run.mocky.io/v3/e6c77e5c-aec9-403f-821b-e14114220148",
}

//Item Struct ->
type Item struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Quantity int    `json:"quantity"`
	Price    string `json:"price"`
}

//Init Items var as a slice Item struct ->
var Items []Item

//Slice for cache elements ->
var cache = make([]string, 10)

func main() {
	//Init Router ->
	r := mux.NewRouter()

	//Route Handlers - Endpoints ->
	r.HandleFunc("/", Home).Methods("GET")
	r.HandleFunc("/buyItem/{name}", getItems).Methods("GET")
	r.HandleFunc("/buyItemqty/{name}&{quantity}", getItemsByQty).Methods("GET")
	r.HandleFunc("/buyItemqtyprice/{name}&{quantity}&{price}", getItemsByPrice).Methods("GET")
	r.HandleFunc("/view", view).Methods("GET")

	fmt.Println("Server started on port:3000")
	log.Fatal(http.ListenAndServe(":3000", r))
}

func Home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1><strong>Food Aggregator</strong></h1>"))
}

func getItems(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	flag := 0
	for _, web := range urls {
		res, err := http.Get(web)

		if err != nil {
			panic(err)
		}

		dataBytes, err := ioutil.ReadAll(res.Body)
		if err != nil {
			panic(err)
		}

		data := dataBytes
		json.Unmarshal(data, &Items)

		params := mux.Vars(r)

		for _, item := range Items {

			if item.Name == params["name"] {
				cache = append(cache, string(item.ID), string(item.Name), string(item.Price), fmt.Sprint(item.Quantity))
				flag = 1
				json.NewEncoder(w).Encode(item)
				return
			}

		}
	}
	fmt.Println(cache)

	if flag == 0 {
		json.NewEncoder(w).Encode("NOT_FOUND")
		cache = append(cache, string("ERROR"))
	}
}

func getItemsByQty(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	flag := 0
	for _, web := range urls {
		res, err := http.Get(web)

		if err != nil {
			panic(err)
		}

		dataBytes, err := ioutil.ReadAll(res.Body)
		if err != nil {
			panic(err)
		}

		data := dataBytes
		json.Unmarshal(data, &Items)

		params := mux.Vars(r)
		qty, _ := strconv.Atoi(params["quantity"])

		for _, item := range Items {
			if item.Name == params["name"] && item.Quantity >= qty {
				cache = append(cache, string(item.ID), string(item.Name), string(item.Price), fmt.Sprint(item.Quantity))
				flag = 1
				json.NewEncoder(w).Encode(item)
				return
			}
		}
	}

	fmt.Println(cache)

	if flag == 0 {
		json.NewEncoder(w).Encode("NOT_FOUND")
		cache = append(cache, string("ERROR"))
	}
}

func getItemsByPrice(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	flag := 0
	for _, web := range urls {
		res, err := http.Get(web)

		if err != nil {
			panic(err)
		}

		dataBytes, err := ioutil.ReadAll(res.Body)
		if err != nil {
			panic(err)
		}

		data := dataBytes
		json.Unmarshal(data, &Items)

		params := mux.Vars(r)
		qty, _ := strconv.Atoi(params["quantity"])

		for _, item := range Items {
			if item.Name == params["name"] && item.Quantity >= qty && item.Price == params["price"] {
				cache = append(cache, string(item.ID), string(item.Name), string(item.Price), fmt.Sprint(item.Quantity))
				flag = 1
				json.NewEncoder(w).Encode(item)
			}

		}
	}

	fmt.Println(cache)

	if flag == 0 {
		json.NewEncoder(w).Encode("NOT_FOUND")
		cache = append(cache, string("ERROR"))
	}
}

func view(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "\tProducts Are-\n")
	fmt.Fprintf(w, "\n ID \t\t Name \t Quantity \t Price")
	fmt.Fprintf(w, "\n 24-583-0264 \t Apple \t 30 \t\t $62.02")
	fmt.Fprintf(w, "\n 75-588-0160 \t Apple \t 28 \t\t $86.41")
	fmt.Fprintf(w, "\n 28-996-2788 \t Banana  21 \t\t $99.41")
	fmt.Fprintf(w, "\n 76-152-3057 \t Carrot  13 \t\t $71.01")
	fmt.Fprintf(w, "\n 74-033-7213 \t okra \t 15 \t\t $61.42")
	fmt.Fprintf(w, "\n 87-108-0068 \t Onion \t 20 \t\t $76.30")
	fmt.Fprintf(w, "\n 66-907-8874 \t wheat \t 22 \t\t $89.96")
	fmt.Fprintf(w, "\n 51-268-1902 \t barley  26 \t\t $50.92")
	fmt.Fprintf(w, "\n 68-684-1026 \t rye \t 14 \t\t $80.90")
	fmt.Println("Endpoint Hit: viewPage")
}
