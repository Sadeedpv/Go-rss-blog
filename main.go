package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main(){
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	port := os.Getenv("PORT")
	if port == ""{
		log.Fatal("NO PORT FOUND")
	}

	log.Print("PORT NO: \n", port)

	router := chi.NewRouter()
	// Basic CORS
	router.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins:   []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))
	v1Router := chi.NewRouter()
	v1Router.Get("/handle", HandlerReady)
	v1Router.Get("/error", HandlerError)

	router.Mount("/v1", v1Router)
	srv := &http.Server{
		Handler: router,
		Addr: ":" + port,
	}

	log.Printf("Server Listening on PORT: %v", port)

	err2 := srv.ListenAndServe()
	if err2 != nil{
		log.Fatal("Server failed to Load!")
	}
}