package main

import (
	"encoding/json"
	"fmt"
	"go-backend/requestProcess"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// RequestBody defines the structure of the expected JSON body
type RequestBody struct {
	Database string `json:"database"`
	Prompt   string `json:"prompt"`
}

func GetResponse(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	// Parse the JSON body
	var reqBody RequestBody
	err := json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	// Generate the response
	gptOutput, err := requestProcess.ResultGen(reqBody.Database, reqBody.Prompt)
	if err != nil {
		http.Error(w, "Error generating response", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	response := map[string]string{"result": gptOutput}
	json.NewEncoder(w).Encode(response)
}

func main() {

	//handling requests
	myRouter := mux.NewRouter()
	myRouter.HandleFunc("/api/v1/dbquery/firebase", GetResponse).Methods("POST")
	fmt.Printf("Server started at http://localhost:1234")
	log.Fatal(http.ListenAndServe(":1234", myRouter))
}
