package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/", index)
	// This will serve files under http://localhost/static/<filename>
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	port, ok := os.LookupEnv("PORT")
	if !ok {
		port = "80"
	}

	addr := fmt.Sprintf(":%s", port)
	server := http.Server{
		Addr:         addr,
		Handler:      router,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	log.Println("Web server started on port ", port)
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Impossible to start web server : %v\n", err)
	}

}
