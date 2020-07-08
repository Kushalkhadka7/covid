package main

import (
	"bufio"
	"covid-19/client/controller"
	grpcconn "covid-19/grpc"
	httpconn "covid-19/http"
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
)

func main() {

	port := flag.Int("port", 8080, "Http port to run the client server")
	flag.Parse()

	// Creates http server.
	server := httpconn.New(*port)

	// Connect to the Grpc server on given port.
	grpcConnection, err := grpcconn.NewGrpc(8081)
	if err != nil {
		log.Fatalf("Could not establish connection to grpc server:%v", err)
	}

	// New grpc client.
	covidClient := controller.NewCovid(grpcConnection.ClientConn)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := scanner.Text()

		matched, err := regexp.MatchString("total.*", text)
		if err != nil {
			fmt.Println("total pattern not matched")
		}

		if matched {
			covidClient.HandleTotalCases(text)
		} else {
			if err := covidClient.HandleCountryCases(text); err != nil {
				log.Fatal(err)
			}
		}

	}

	if err := server.Start(); err != nil {
		fmt.Println(err)
	}
	defer server.Close()
}
