package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

type ServerResponse struct {
	Bid string `json:"bid"`
}

func main() {
	url := "http://localhost:8080/cotacao"

	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Millisecond)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		log.Fatalf("Failed to create request: %v\n", err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Printf("Request timed out: %v\n", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Printf("Server returned an error: %s\n", resp.Status)
		return
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Failed to read server response: %v\n", err)
	}

	var serverResponse ServerResponse
	err = json.Unmarshal(body, &serverResponse)
	if err != nil {
		log.Fatalf("Error decoding server response: %v\n", err)
	}

	fileContent := fmt.Sprintf("Dólar: %s\n", serverResponse.Bid)
	err = os.WriteFile("cotacao.txt", []byte(fileContent), 0644)
	if err != nil {
		log.Fatalf("Failed to write to file: %v\n", err)
	}

	log.Println("Exchange rate saved to 'cotacao.txt'")
}
