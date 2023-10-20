package main

import (
	"log"
	"plogger/internal/server"
)

func main(){
	srv := server.NewHttpServer(":8080")
	log.Fatal(srv.ListenAndServe())
}