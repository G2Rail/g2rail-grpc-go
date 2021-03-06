package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	"g2rail-grpc-go/client"
	g2rail "g2rail-grpc-go/g2rail/lib/proto"

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

	suggestionClient := client.NewSuggestionsClientsClient(cc2)

	suggestions := suggestionClient.Query("BERLIN")

	fmt.Println(suggestions)

	onlineSolutionsClient := client.NewOnlineSolutionsClient(cc2)
	asyncKey := onlineSolutionsClient.Search()
	var solutions []*g2rail.RailwaySolution
	var bookingCode string
	for {
		time.Sleep(2 * time.Second)
		solutions = onlineSolutionsClient.QueryAsyncOnlineSolutions(asyncKey)
		if !resultIsLoading(solutions) {
			break
		}
	}
	if len(solutions) == 0 {
		fmt.Println("No result found!")
		return
	}

	bookingCode = solutions[0].Solutions[0].Sections[0].Offers[1].Services[0].BookingCode
	onlineOrderClient := client.NewOnlineOrdersClient(cc2)

	asyncKey = onlineOrderClient.Book(bookingCode)
	// asyncKey := onlineOrderClient.Book("P_8GKRW0")
	var onlineOrder *g2rail.OnlineOrderResponse
	for {
		onlineOrder = onlineOrderClient.QueryAsyncOnlineOrder(asyncKey)
		if !onlineOrder.Loading {
			break
		}
		time.Sleep(2 * time.Second)
	}
	fmt.Println(onlineOrder)

	onlineConfimationClient := client.NewOnlineConfirmationsClient(cc2)
	asyncKey = onlineConfimationClient.Confirm(onlineOrder.Id)

	var onlineConfirmation *g2rail.OnlineConfirmationResponse
	for {
		onlineConfirmation = onlineConfimationClient.QueryAsyncOnlineConfirmation(asyncKey)
		if !onlineConfirmation.Loading {
			break
		}
		time.Sleep(2 * time.Second)
	}

	onlineTicketsClient := client.NewOnlineTicketsClient(cc2)
	tickets := onlineTicketsClient.Download(onlineOrder.Id)
	fmt.Println(tickets)
}

func resultIsLoading(solutions []*g2rail.RailwaySolution) bool {
	for _, solution := range solutions {
		if solution.Loading {
			return true
		}
	}
	return false
}
