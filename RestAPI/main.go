package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Iklan struct {
	ID       string    `json:"id"`
	TipeID   string    `json:"TipeID"`
	Judul    string    `json:"judul"`
	IsiIklan *IsiIklan `json:"isiiklan"`
}

type IsiIklan struct {
	FirstText string `json:"firsttext"`
	LastText  string `json:"lasttext"`
}

var iklans []Iklan

func getIklans(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(iklans)
}

func getIklan(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // Gets params
	for _, item := range iklans {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Iklan{})
}

// Add
func createIklan(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var iklan Iklan
	_ = json.NewDecoder(r.Body).Decode(&iklan)
	iklan.ID = strconv.Itoa(rand.Intn(10))
	iklans = append(iklans, iklan)
	json.NewEncoder(w).Encode(iklan)
}

// Update
func updateIklan(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range iklans {
		if item.ID == params["id"] {
			iklans = append(iklans[:index], iklans[index+1:]...)
			var iklan Iklan
			_ = json.NewDecoder(r.Body).Decode(&iklan)
			iklan.ID = params["id"]
			iklans = append(iklans, iklan)
			json.NewEncoder(w).Encode(iklan)
			return
		}
	}
}

// delete
func deleteIklan(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range iklans {
		if item.ID == params["id"] {
			iklans = append(iklans[:index], iklans[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(iklans)
}

// Main function
func main() {
	// Init router
	r := mux.NewRouter()

	// Route
	r.HandleFunc("/iklans", getIklans).Methods("GET")
	r.HandleFunc("/iklans/{id}", getIklan).Methods("GET")
	r.HandleFunc("/iklans", createIklan).Methods("POST")
	r.HandleFunc("/iklans/{id}", updateIklan).Methods("PUT")
	r.HandleFunc("/iklans/{id}", deleteIklan).Methods("DELETE")

	// Start server
	log.Fatal(http.ListenAndServe(":8000", r))
}
