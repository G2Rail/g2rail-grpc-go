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

// SuggestionsClient is a client to call online confiramtions service RPCs
type SuggestionsClient struct {
	service g2rail.SuggestionsClient
}

// Query calls Suggestions RPC to retrieve possible city/stations by a given criteria.
func (suggestions *SuggestionsClient) Query(query string) *g2rail.SuggestionsResponse {
	req := &g2rail.SuggestionRequest{}

	// set timeout
	ctx, cancel := context.WithTimeout(context.Background(), 400*time.Second)
	defer cancel()

	res, err := suggestions.service.Query(ctx, req)
	if err != nil {
		st, ok := status.FromError(err)
		if ok && st.Code() == codes.PermissionDenied {
			log.Fatal("permission denied")
		} else {
			log.Fatal("cannot perform search: ", err)
		}
		return &g2rail.SuggestionsResponse{}
	}

	log.Printf("Possible matches are: %s", res)
	return res
}

// NewSuggestionsClientsClient returns a new online solutions client
func NewSuggestionsClientsClient(cc *grpc.ClientConn) *SuggestionsClient {
	service := g2rail.NewSuggestionsClient(cc)
	return &SuggestionsClient{service}
}
