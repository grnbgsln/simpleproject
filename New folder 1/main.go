package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

type data struct {
	Data interface{} `json:"data"`
}

type product struct {
	Merk_mobil  string `json:merk_mobil`
	Jenis_mobil string `json:jenis_mobil`
	Transmisi   string `json:transmisi`
	Harga       int    `json:harga`
}

var cars = []*product{
	{"Ford", "Fiesta", "Manual", 165000000},
	{"Ford", "Fiesta", "Manual", 175000000},
	{"Ford", "Fiesta", "Automatic", 180000000},
	{"Ford", "Fiesta", "Manual", 155000000},
	{"VW", "Polo", "Manual", 170000000},
	{"VW", "Beetle", "Manual", 265000000},
	{"VW", "Polo", "Automatic", 165000000},
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", GetMobil).Methods("GET")
	router.HandleFunc("/merk/{merk_mobil}", GetMerkMobil).Methods("GET")
	router.HandleFunc("/jenis/{jenis_mobil}", getJenisMobil).Methods("GET")
	http.ListenAndServe(":8080", router)
}

func GetMobil(w http.ResponseWriter, r *http.Request) {
	res := data{Data: cars}
	jsonData, err := json.Marshal(res)
	if err != nil {
		fmt.Println(err.Error())
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(jsonData)
}

func GetMerkMobil(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	merkMobil := vars["merk_mobil"]

	mySlice := []*product{}
	for _, v := range cars {
		if strings.ToLower(v.Merk_mobil) == strings.ToLower(merkMobil) {
			x := &product{v.Merk_mobil, v.Jenis_mobil, v.Transmisi, v.Harga}
			mySlice = append(mySlice, x)
		}
	}

	res := data{
		Data: mySlice,
	}
	jsonData, err := json.Marshal(res)
	if err != nil {
		fmt.Println(err.Error())
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(jsonData)
}
func getJenisMobil(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	getJenisMobil := vars["jenis_mobil"]

	mySlice := []*product{}
	for _, v := range cars {
		if strings.ToLower(v.Jenis_mobil) == strings.ToLower(getJenisMobil) {
			x := &product{v.Jenis_mobil, v.Merk_mobil, v.Transmisi, v.Harga}
			mySlice = append(mySlice, x)
		}
	}

	res := data{
		Data: mySlice,
	}

	jsonData, err := json.Marshal(res)
	if err != nil {
		fmt.Println(err.Error())
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(jsonData)
}
