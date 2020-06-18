package client

import (
	"context"
	g2rail "g2rail-grpc-go/g2rail/lib/proto"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// OnlineOrdersClient is a client to call online solutions service RPCs
type OnlineOrdersClient struct {
	service g2rail.OnlineOrdersClient
}

// Book calls create onlineOrder RPC
func (onlineOrder *OnlineOrdersClient) Book(bookingCode string) string {
	req := &g2rail.BookRequest{
		Sections: []string{bookingCode},
		Passengers: []*g2rail.Passenger{&g2rail.Passenger{
			Gender:    g2rail.Passenger_male,
			FirstName: "QINWEN",
			LastName:  "SHI",
			Passport:  "E12341813",
			Phone:     "+8527892123",
			Email:     "wen@g2rail.com",
			Birthdate: "1986-06-01",
		}},
	}

	// set timeout
	ctx, cancel := context.WithTimeout(context.Background(), 400*time.Second)
	defer cancel()

	res, err := onlineOrder.service.Book(ctx, req)
	if err != nil {
		st, ok := status.FromError(err)
		if ok && st.Code() == codes.PermissionDenied {
			log.Fatal("permission denied")
		} else {
			log.Fatal("cannot perform search: ", err)
		}
		return ""
	}

	log.Printf("Async Result Will Be At: %s", res.AsyncKey)
	return res.AsyncKey
}

// QueryAsyncOnlineOrder calls retrieve onlineOrder RPC
func (onlineOrder *OnlineOrdersClient) QueryAsyncOnlineOrder(asyncKey string) *g2rail.OnlineOrderResponse {
	req := &g2rail.OnlineOrderAsyncQueryRequest{
		AsyncKey: asyncKey,
	}

	// set timeout
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	res, err := onlineOrder.service.QueryAsyncOnlineOrder(ctx, req)
	if err != nil {
		st, ok := status.FromError(err)
		if ok && st.Code() == codes.PermissionDenied {
			log.Fatal("permission denied")
		} else {
			log.Fatal("cannot perform book: ", err)
		}
		return &g2rail.OnlineOrderResponse{}
	}

	log.Printf("Async Result Is: %s", res)
	return res
}

// NewOnlineOrdersClient returns a new online solutions client
func NewOnlineOrdersClient(cc *grpc.ClientConn) *OnlineOrdersClient {
	service := g2rail.NewOnlineOrdersClient(cc)
	return &OnlineOrdersClient{service}
}
