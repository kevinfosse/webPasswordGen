package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

type passGenerated struct {
	Password string `json:"password"`
	Length   string `json:"lenght"`
}

var response passGenerated

type test_struct struct {
	Test string
}

func Passfunc(w http.ResponseWriter, r *http.Request) {

	length, _ := strconv.Atoi(response.Length)
	response.Password = createPassword(length)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	jsonresponse, err := json.Marshal(response)
	if err != nil {
		return
	}
	w.Write(jsonresponse)

}

func Parameterfunc(w http.ResponseWriter, r *http.Request) {

	reqBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBody, &response)
	length, _ := strconv.Atoi(response.Length)
	response.Password = createPassword(length)
	json.NewEncoder(w).Encode(response)

}

func main() {
	fmt.Println("Démarrage du serveur sur le port :8080")
	// Default lenght = 8
	response.Length = "8"
	router := mux.NewRouter()
	router.HandleFunc("/api/generate", Passfunc).Methods("GET")
	router.HandleFunc("/api/parameter", Parameterfunc).Methods("PUT")
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./server/")))
	http.Handle("/", router)

	corsObj := handlers.AllowedOrigins([]string{"*"})
	log.Fatal(http.ListenAndServe(":8080", handlers.CORS(corsObj)(router)))
	//http.ListenAndServe(":8080", router)

}
