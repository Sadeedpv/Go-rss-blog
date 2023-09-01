package main

import "net/http"

func HandlerReady(w http.ResponseWriter, r *http.Request) {
	ResponswithJson(w,200, struct{}{})
}

func HandlerError(w http.ResponseWriter, r *http.Request) {
	RespondwithError(w,200, "Something went wrong!")
}