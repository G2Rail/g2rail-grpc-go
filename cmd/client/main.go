package main

import (
	"flag"
	"log"
	"time"

	"g2rail-grpc-go/client"
	"g2rail-grpc-go/g2rail"

	"google.golang.org/grpc"
)

const (
	refreshDuration = 30 * time.Second
)

func main() {
	serverAddress := flag.String("address", "", "the server address")
	apiKey := flag.String("apikey", "", "The API Key From G2Rail")
	apiSecret := flag.String("secret", "", "The API Secret From G2Rail")
	enableTLS := flag.Bool("tls", false, "enable SSL/TLS")
	flag.Parse()
	log.Printf("dial server %s, TLS = %t", *serverAddress, *enableTLS)

	transportOption := grpc.WithInsecure()

	interceptor, err := client.NewAuthInterceptor(*apiKey, *apiSecret)
	if err != nil {
		log.Fatal("cannot create auth interceptor: ", err)
	}

	cc2, err := grpc.Dial(
		*serverAddress,
		transportOption,
		grpc.WithUnaryInterceptor(interceptor.Unary()),
	)
	if err != nil {
		log.Fatal("cannot dial server: ", err)
	}
	onlineSolutionsClient := client.NewOnlineSolutionsClient(cc2)
	asyncKey := onlineSolutionsClient.Search()

	for {
		time.Sleep(2 * time.Second)
		solutions := onlineSolutionsClient.QueryAsyncOnlineSolutions(asyncKey)
		if !resultIsLoading(solutions) {
			break
		}
	}
}

func resultIsLoading(solutions []*g2rail.RailwaySolution) bool {
	for _, solution := range solutions {
		if solution.Loading {
			return true
		}
	}
	return false
}
