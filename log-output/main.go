package main

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"flag"
	"time"

	// "fmt"
	"log"
	"net/http"
	"strconv"
)

func generateHash() string {
	key := make([]byte, 32)

	rand.Read(key)

	return base64.StdEncoding.EncodeToString(key)
}

var hashPort int

func init() {
	flag.IntVar(&hashPort, "port", 8081, "port which the server runs on")

	flag.Parse()
}

func main() {
	processHash := generateHash()

	type HashResponse struct {
		ApplicationHash string `json:"appHash"`
		RequestHash     string `json:"reqHash"`
	}

	type TimestampResponse struct {
		ApplicationHash string `json:"appHash"`
		Timestamp       string `json:"timestamp"`
	}

	hashHandler := func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		hashRes := HashResponse{
			ApplicationHash: processHash,
			RequestHash:     generateHash(),
		}

		err := json.NewEncoder(w).Encode(hashRes)
		if err != nil {
			//handle marshalling error
			http.Error(w, "failed to encode response", http.StatusInternalServerError)
			return
		}

	}
	timestampHandler := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		timestampResponse := TimestampResponse{
			ApplicationHash: processHash,
			Timestamp:       time.Now().Format(time.RFC822),
		}

		err := json.NewEncoder(w).Encode(timestampResponse)
		if err != nil {
			http.Error(w, "failed to encode response", http.StatusInternalServerError)
			return
		}
	}

	http.HandleFunc("/hash", hashHandler)
	http.HandleFunc("/timestamp", timestampHandler)
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(hashPort), nil))

}
