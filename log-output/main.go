package main

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"flag"
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

	type RequestResponse struct {
		ApplicationHash string `json:"appHash"`
		RequestHash     string `json:"reqHash"`
	}

	hashHandler := func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		hashRes := RequestResponse{
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

	http.HandleFunc("/hash", hashHandler)
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(hashPort), nil))

}
