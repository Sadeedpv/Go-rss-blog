package main

import "net/http"

type Response struct{
	Message []string `json:"post"`
}


func HandlerReady(w http.ResponseWriter, r *http.Request) {
	message := []string {`Hey this is my first blog post`, `Hey this is my second blog post`}
	repsonse := Response{
	Message: message,
}
	ResponswithJson(w,200, repsonse)
}

func HandlerError(w http.ResponseWriter, r *http.Request) {
	RespondwithError(w,200, "Something went wrong!")
}