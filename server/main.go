package main

import (
	"log"
	"server/db"
	"server/router"
)

func main() {
	// Initialize Firestore client
	store, err := db.InitFirestore()
	if err != nil {
		log.Println("Error: Firestore client not initialized")
		return
	}
	s := router.NewServer(store)
	
	s.ListenAndServe()
}
