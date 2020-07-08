package main

import (
	httpconn "covid-19/http"
	controller "covid-19/server/controller"
	"log"

	"os"
)

func main() {

	server := httpconn.New(8081)
	listener, err := server.StartServer()
	if err != nil {
		log.Fatalf("Error in fetching data form api:%v", err)
		os.Exit(1)
	}
	defer server.Close()

	newCovidServer := controller.NewServer(listener)
	if err := newCovidServer.NewGrpcServer(); err != nil {
		log.Fatalf("Could not establish new grpc controller;%v", err)
		os.Exit(1)
	}
}
