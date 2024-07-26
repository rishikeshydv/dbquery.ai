package main

import (
	"encoding/json"
	"fmt"
	"go-backend/requestProcess"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

// RequestBody defines the structure of the expected JSON body
type RequestBody struct {
	Database string `json:"database"`
	Prompt   string `json:"prompt"`
}

func GetResponse(w http.ResponseWriter, r *http.Request) {
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

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := map[string]string{"status": "ok"}
	json.NewEncoder(w).Encode(response)
}

func main() {

	//handling requests
	myRouter := mux.NewRouter()
	myRouter.HandleFunc("/api/v1/dbquery/firebase", GetResponse).Methods("POST")
	myRouter.HandleFunc("/api/v1/health", HealthCheck).Methods("GET")
	// Create a CORS handler with the desired options
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3008"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		ExposedHeaders:   []string{"X-Total-Count"},
		AllowCredentials: true,
	})
	handler := c.Handler(myRouter)
	fmt.Printf("Server started at http://localhost:9002")
	log.Fatal(http.ListenAndServe(":9002", handler))
}
