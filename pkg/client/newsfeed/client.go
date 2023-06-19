package newsfeed_client

import (
	"github.com/ducnd58233/newsfeed-be/pkg/types/proto/pb/newsfeed"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type newsfeedClient struct {
	clients []newsfeed.NewsfeedClient
}

func NewClient(hosts []string) (newsfeed.NewsfeedClient, error) {
	var opts = []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	clients := make([]newsfeed.NewsfeedClient, 0, len(hosts))
	for _, host := range hosts {
		conn, err := grpc.Dial(host, opts...)
		if err != nil {
			return nil, err
		}
		client := newsfeed.NewNewsfeedClient(conn)
		clients = append(clients, client)
	}

	return &newsfeedClient{clients: clients}, nil
}