package main

import (
	"encoding/json"
	"log"
	"net/http"
)

var mainNumber int

func main() {
	mainNumber = 0
	http.HandleFunc("/counter", quoteHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
func quoteHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/counter" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}
	keys, ok := r.URL.Query()["operation"]
	if !ok || len(keys) == 0 {
		http.Error(w, "400 bad request.", http.StatusBadRequest)
		return
	}
	if keys[0] != "increase" && keys[0] != "decrease" {
		http.Error(w, "501 not implemented.", http.StatusNotImplemented)
		return
	}
	if r.Method == http.MethodGet {
		operation := keys[0]
		var answer Num
		if operation == "increase" {
			answer = increaseNumber()
		} else {
			answer = decreaseNumber()
		}
		output, err := json.Marshal(answer)
		if err != nil {
			log.Fatal(err)
		}
		w.Header().Set("content-type", "application/json")
		_, err = w.Write(output)
		if err != nil {
			return
		}

	} else {
		http.Error(w, "501 not implemented.", http.StatusNotImplemented)
		return
	}
}
func increaseNumber() Num {
	var answer Num
	answer.OldNumber = mainNumber
	mainNumber += 1
	answer.NewNumber = mainNumber
	return answer
}
func decreaseNumber() Num {
	var answer Num
	answer.OldNumber = mainNumber
	mainNumber -= 1
	answer.NewNumber = mainNumber
	return answer
}
