package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func RespondwithError(w http.ResponseWriter, statusCode int, msg string){
	if statusCode > 499{
		log.Println("5XX Error: ", msg)
	}

	type errResponse struct{
		Error string `json: "error"`
	}

	ResponswithJson(w, statusCode, errResponse{
		Error: msg,
	})
}

func ResponswithJson(w http.ResponseWriter, statsCode int, payload interface{}){
	data, err := json.Marshal(payload)
	if err != nil{
		log.Fatal("Failed to marshal json!")
		w.WriteHeader(500) // Internal server error
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statsCode)
	w.Write(data)

}