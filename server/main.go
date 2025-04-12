package main

import (
	"net/http"
	"log"
	"os"
	"os/signal"
	"syscall"
	"context"
	"github.com/prabhjotaulakh159/doc-save/db"
)

func main() {
	mongoClient, err := db.GetMongoClient()
	if err != nil {
		log.Fatalf("error getting db conn, %v", err)	
	}
	
	log.Println("connected to database")	
	
	defer func() {
		if err = mongoClient.Disconnect(context.TODO()); err != nil {
			log.Fatalf("error disconnecting db conn, %v", err)	
		}
		log.Println("db conn closed")
	}()
	
	handler := http.NewServeMux()
	server := &http.Server {
		Addr: "localhost:8000", 
		Handler: handler,
	}

	go func(){
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("error in starting server %v", err)
		}
	}()
	
	log.Printf("server listening on %s", server.Addr)
	
	channel := make(chan os.Signal, 1)
	signal.Notify(channel, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	<-channel
	
	log.Println("attempting to stop server")
	
	if err := server.Shutdown(context.Background()); err != nil {
		log.Fatalf("error in closing server %v", err)	
	}
	
	log.Println("server stopped listening")	
}