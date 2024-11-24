package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	_ "github.com/mattn/go-sqlite3" // Driver SQLite
)

var db *sql.DB

type APIResponse struct {
	USDBRL CurrencyData `json:"USDBRL"`
}

type CurrencyData struct {
	Code       string `json:"code"`
	Codein     string `json:"codein"`
	Name       string `json:"name"`
	High       string `json:"high"`
	Low        string `json:"low"`
	VarBid     string `json:"varBid"`
	PctChange  string `json:"pctChange"`
	Bid        string `json:"bid"`
	Ask        string `json:"ask"`
	Timestamp  string `json:"timestamp"`
	CreateDate string `json:"create_date"`
}

func main() {
	var err error

	db, err = setupDB()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	http.HandleFunc("/cotacao", quoteHandler)
	port := ":8080"

	fmt.Printf("Server running on port %s...\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}

func setupDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./usd_prices.db")
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
		return db, err
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS usd_prices (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		bid TEXT NOT NULL,
		timestamp DATETIME DEFAULT CURRENT_TIMESTAMP
	)`)
	if err != nil {
		log.Fatalf("Failed to create table: %v", err)
		return db, err
	}

	return db, err
}

func quoteHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	apiCtx, cancelAPI := context.WithTimeout(ctx, 200*time.Millisecond)
	defer cancelAPI()

	apiURL := "https://economia.awesomeapi.com.br/json/last/USD-BRL"
	req, err := http.NewRequestWithContext(apiCtx, http.MethodGet, apiURL, nil)
	if err != nil {
		http.Error(w, "Failed to create request", http.StatusInternalServerError)
		log.Printf("Error creating API request: %v\n", err)
		return
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		http.Error(w, "Timeout fetching exchange rate", http.StatusRequestTimeout)
		log.Printf("API request timed out: %v\n", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "Error reading response", http.StatusInternalServerError)
		return
	}

	var apiResponse APIResponse
	err = json.Unmarshal(body, &apiResponse)
	if err != nil {
		http.Error(w, "Error decoding API response", http.StatusInternalServerError)
		return
	}

	dbCtx, cancelDB := context.WithTimeout(ctx, 10*time.Millisecond)
	defer cancelDB()

	_, err = db.ExecContext(dbCtx, "INSERT INTO usd_prices (bid) VALUES (?)", apiResponse.USDBRL.Bid)
	if err != nil {
		http.Error(w, "Timeout persisting data", http.StatusRequestTimeout)
		log.Printf("Database insertion timed out: %v\n", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	bidResponse := map[string]string{"bid": apiResponse.USDBRL.Bid}
	json.NewEncoder(w).Encode(bidResponse)
}
