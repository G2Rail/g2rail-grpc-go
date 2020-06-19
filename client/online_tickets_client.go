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

// OnlineTicketsClient is a client to call online confiramtions service RPCs
type OnlineTicketsClient struct {
	service g2rail.OnlineTicketsClient
}

// Download calls create onlineOrder RPC
func (onlineTickets *OnlineTicketsClient) Download(orderID string) *g2rail.TicketsResponse {
	req := &g2rail.DownloadRequest{
		OrderId: orderID,
	}

	// set timeout
	ctx, cancel := context.WithTimeout(context.Background(), 400*time.Second)
	defer cancel()

	res, err := onlineTickets.service.Download(ctx, req)
	if err != nil {
		st, ok := status.FromError(err)
		if ok && st.Code() == codes.PermissionDenied {
			log.Fatal("permission denied")
		} else {
			log.Fatal("cannot perform search: ", err)
		}
		return &g2rail.TicketsResponse{}
	}

	log.Printf("Your Tickets are: %s", res)
	return res
}

// NewOnlineTicketsClient returns a new online solutions client
func NewOnlineTicketsClient(cc *grpc.ClientConn) *OnlineTicketsClient {
	service := g2rail.NewOnlineTicketsClient(cc)
	return &OnlineTicketsClient{service}
}
