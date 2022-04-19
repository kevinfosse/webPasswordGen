package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

type passGenerated struct {
	Password string `json:"password"`
}

func Passfunc(w http.ResponseWriter, r *http.Request) {

	var response passGenerated

	response.Password = createPassword()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	jsonresponse, err := json.Marshal(response)
	if err != nil {
		return
	}
	w.Write(jsonresponse)

}

func main() {
	fmt.Println("Démarrage du serveur sur le port :8080")

	router := mux.NewRouter()
	router.HandleFunc("/api/generate", Passfunc).Methods("GET")
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./server/")))
	http.Handle("/", router)

	corsObj := handlers.AllowedOrigins([]string{"*"})
	log.Fatal(http.ListenAndServe(":8080", handlers.CORS(corsObj)(router)))
	//http.ListenAndServe(":8080", router)

}
