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

// OnlineConfirmationsClient is a client to call online confiramtions service RPCs
type OnlineConfirmationsClient struct {
	service g2rail.OnlineConfirmationsClient
}

// Confirm calls create onlineOrder RPC
func (onlineConfirmation *OnlineConfirmationsClient) Confirm(orderID string) string {
	req := &g2rail.ConfirmRequest{
		OrderId: orderID,
	}

	// set timeout
	ctx, cancel := context.WithTimeout(context.Background(), 400*time.Second)
	defer cancel()

	res, err := onlineConfirmation.service.Confirm(ctx, req)
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

// QueryAsyncOnlineConfirmation calls retrieve onlineOrder RPC
func (onlineConfirmation *OnlineConfirmationsClient) QueryAsyncOnlineConfirmation(asyncKey string) *g2rail.OnlineConfirmationResponse {
	req := &g2rail.OnlineConfirmationAsyncQueryRequest{
		AsyncKey: asyncKey,
	}

	// set timeout
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	res, err := onlineConfirmation.service.QueryAsyncOnlineConfirmation(ctx, req)
	if err != nil {
		st, ok := status.FromError(err)
		if ok && st.Code() == codes.PermissionDenied {
			log.Fatal("permission denied")
		} else {
			log.Fatal("cannot perform book: ", err)
		}
		return &g2rail.OnlineConfirmationResponse{}
	}

	log.Printf("Async Result Is: %s", res)
	return res
}

// NewOnlineConfirmationsClient returns a new online solutions client
func NewOnlineConfirmationsClient(cc *grpc.ClientConn) *OnlineConfirmationsClient {
	service := g2rail.NewOnlineConfirmationsClient(cc)
	return &OnlineConfirmationsClient{service}
}
