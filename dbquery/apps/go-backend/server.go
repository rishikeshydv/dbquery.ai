package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	openai "github.com/sashabaranov/go-openai"
)

// RequestPayload defines the structure of the expected JSON body
type RequestBody struct {
	Prompt string `json:"prompt"`
}

// use godot package to load/read the .env file and
// return the value of the key
func loadEnv() error {
	err := godotenv.Load(".env")
	if err != nil {
		return fmt.Errorf("error loading .env file: %w", err)
	}
	return nil
}

// goDotEnvVariable returns the value of the given key from the environment variables
func goDotEnvVariable(key string) string {
	// Load .env file
	err := loadEnv()
	if err != nil {
		fmt.Printf("Completion error: ", err)
	}

	// Get the value of the key
	value := os.Getenv(key)
	if value == "" {
		fmt.Errorf("key %s not found in the environment variables", key)
	}

	return value
}

func resultGen(prompt_ string) (string, error) {
	// godotenv package
	secretKey := goDotEnvVariable("OPENAI_SECRET_KEY")

	// create a new client
	client := openai.NewClient(secretKey)
	ctx := context.Background()
	resp, err := client.CreateChatCompletion(
		ctx,
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: prompt_,
				},
			},
		},
	)

	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
		return "", err
	}

	fmt.Println(resp.Choices[0].Message.Content)
	return resp.Choices[0].Message.Content, nil
}

func getResponse(w http.ResponseWriter, r *http.Request) {

	// Parse the JSON body
	var reqBody RequestBody
	err := json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Process the prompt
	fmt.Printf("Received: %v\n", reqBody.Prompt)

	// Generate the response
	w.Header().Set("Content-Type", "application/json")
	gptOutput, err := resultGen(reqBody.Prompt)
	if err != nil {
		http.Error(w, "Error generating response", http.StatusInternalServerError)
		return
	}
	response := map[string]string{"result": gptOutput}
	json.NewEncoder(w).Encode(response)
}

func main() {

	//handling requests
	myRouter := mux.NewRouter()
	myRouter.HandleFunc("/api/v1/dbquery/firebase", getResponse).Methods("POST")
	fmt.Printf("Server started at http://localhost:8080\n")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
