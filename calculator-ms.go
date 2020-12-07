package main

import (
	"./service"
	"encoding/json"
	"strconv"
	"strings"
	"fmt"
	"net/http"
	"errors"
)

var operationHistory []service.Operation

func urlParser(url string) (float64, float64, error){
	splitURL := strings.Split(url, "/")
	if len(splitURL) == 5 {
		if a, err := strconv.ParseFloat(splitURL[3], 64); err != nil {
			return 0.0, 0.0, errors.New("Your first input ["+splitURL[3]+"] is not a valid number (ex: 3.0)")
		} else{
			if b, err := strconv.ParseFloat(splitURL[4], 64); err != nil {
				return 0.0, 0.0, errors.New("Your second input ["+splitURL[4]+"] is not a valid number (ex: 3.0)")
			} else{
				return a, b, nil
			}
		}
	} else {
		return 0.0, 0.0, errors.New("Invalid Operation")
	}
}

func newOperation(w http.ResponseWriter, operation service.Operation){
	operationHistory = append(operationHistory, operation)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(operation)
}

func description(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Available operations: \n/calc/sum/$a/$b\n/calc/sub/$a/$b\n/calc/mul/$a/$b\n/calc/div/$a/$b\n/calc/history")
}

func history(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(operationHistory)
}

func sum(w http.ResponseWriter, r *http.Request) {
	if a, b, err := urlParser(r.URL.Path); err == nil {
		newOperation(w, service.Sum(a,b))
	} else {
		fmt.Fprint(w, "Available operations: \n/calc/sum/$a/$b\n/calc/sub/$a/$b\n/calc/mul/$a/$b\n/calc/div/$a/$b\n/calc/history\n\nError: "+err.Error())
	}
}

func sub(w http.ResponseWriter, r *http.Request) {
	if a, b, err := urlParser(r.URL.Path); err == nil {
		newOperation(w, service.Sub(a,b))
	} else {
		fmt.Fprint(w, "Available operations: \n/calc/sum/$a/$b\n/calc/sub/$a/$b\n/calc/mul/$a/$b\n/calc/div/$a/$b\n/calc/history\n\nError: "+err.Error())
	}
}

func mul(w http.ResponseWriter, r *http.Request) {
	if a, b, err := urlParser(r.URL.Path); err == nil {
		newOperation(w, service.Mul(a,b))
	} else {
		fmt.Fprint(w, "Available operations: \n/calc/sum/$a/$b\n/calc/sub/$a/$b\n/calc/mul/$a/$b\n/calc/div/$a/$b\n/calc/history\n\nError: "+err.Error())
	}
}

func div(w http.ResponseWriter, r *http.Request) {
	if a, b, err := urlParser(r.URL.Path); err == nil {
		newOperation(w, service.Div(a,b))
	} else {
		fmt.Fprint(w, "Available operations: \n/calc/sum/$a/$b\n/calc/sub/$a/$b\n/calc/mul/$a/$b\n/calc/div/$a/$b\n/calc/history\n\nError: "+err.Error())
	}
}

func main() {
	operationHistory = []service.Operation{}
	fmt.Print("Calculator-Ms at http://0.0.0.0:8080")
	http.HandleFunc("/calc/sum/", sum)
	http.HandleFunc("/calc/sub/", sub)
	http.HandleFunc("/calc/mul/", mul)
	http.HandleFunc("/calc/div/", div)
	http.HandleFunc("/calc/history", history)
	http.HandleFunc("/calc", description)
	http.ListenAndServe(":8080", nil)
}
