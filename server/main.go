package main

import (
	"fmt"
	"log"
	"net/http"
	"toDO_list/server/middleware"
	"toDO_list/server/routes"
	"toDO_list/server/models"
)

func main(){
	r := routes.Router()
	fmt.Println("Starting server on the port 8080...")
	log.Fatal(http.ListenAndServe(":8080", r))
}