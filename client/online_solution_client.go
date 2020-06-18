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

// OnlineSolutionsClient is a client to call online solutions service RPCs
type OnlineSolutionsClient struct {
	service g2rail.OnlineSolutionsClient
}

// Search calls create onlineSolution RPC
func (onlineSolution *OnlineSolutionsClient) Search() string {
	req := &g2rail.SearchRequest{
		From:  "BERLIN",
		To:    "FRANKFURT",
		Date:  "2020-06-22",
		Time:  "08:00",
		Adult: 1,
		Child: 0,
	}

	// set timeout
	ctx, cancel := context.WithTimeout(context.Background(), 400*time.Second)
	defer cancel()

	res, err := onlineSolution.service.Search(ctx, req)
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

// QueryAsyncOnlineSolutions calls retrieve onlineSolution RPC
func (onlineSolution *OnlineSolutionsClient) QueryAsyncOnlineSolutions(asyncKey string) []*g2rail.RailwaySolution {
	req := &g2rail.OnlineSolutionsAsyncQueryRequest{
		AsyncKey: asyncKey,
	}

	// set timeout
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	res, err := onlineSolution.service.QueryAsyncOnlineSolutions(ctx, req)
	if err != nil {
		st, ok := status.FromError(err)
		if ok && st.Code() == codes.PermissionDenied {
			log.Fatal("permission denied")
		} else {
			log.Fatal("cannot perform search: ", err)
		}
		return make([]*g2rail.RailwaySolution, 0)
	}

	log.Printf("Async Result Is: %s", res.RailwaySolutions)
	return res.RailwaySolutions
}

// NewOnlineSolutionsClient returns a new online solutions client
func NewOnlineSolutionsClient(cc *grpc.ClientConn) *OnlineSolutionsClient {
	service := g2rail.NewOnlineSolutionsClient(cc)
	return &OnlineSolutionsClient{service}
}
