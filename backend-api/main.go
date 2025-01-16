package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

type BinaryRequest struct {
	Left  int `json:"a"`
	Right int `json:"b"`
}

type Response struct {
	Result int `json:"result"`
}

func getRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got / request\n")
	io.WriteString(w, "This is the root command\n")
}

func handleBinary(w http.ResponseWriter, r *http.Request, op string) {
  w.Header().Set("Content-Type", "application/json")

	decoder := json.NewDecoder(r.Body)

	var bireq BinaryRequest
	err := decoder.Decode(&bireq)
	if err != nil {
		fmt.Printf("could not read body of reaquest: %s\n", err)
	}

  var res int
  switch op {
  case "add":
    res = bireq.Left + bireq.Right
  case "sub":
    res = bireq.Left - bireq.Right
  case "mul":
    res = bireq.Left * bireq.Right
  default:
    res = -9999
  return
  }

	result := Response{
		Result: res,
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(result); err != nil {
		http.Error(w, "Error creating response", http.StatusInternalServerError)
		return
	}
}

func getAdd(w http.ResponseWriter, r *http.Request) {
  handleBinary(w, r, "add")
}

func getSub(w http.ResponseWriter, r *http.Request) {
  handleBinary(w, r, "sub")
}

func getMult(w http.ResponseWriter, r *http.Request) {
  handleBinary(w, r, "mul")
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", getRoot)
	mux.HandleFunc("/add", getAdd)
	mux.HandleFunc("/subtract", getSub)
	mux.HandleFunc("/multiply", getMult)

	err := http.ListenAndServe(":3333", mux)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}
