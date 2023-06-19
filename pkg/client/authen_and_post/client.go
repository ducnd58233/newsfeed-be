package authen_and_post_client

import (
	"github.com/ducnd58233/newsfeed-be/pkg/types/proto/pb/authen_and_post"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type aapClient struct {
	clients []authen_and_post.AuthenticateAndPostClient
}

func NewClient(hosts []string) (authen_and_post.AuthenticateAndPostClient, error) {
	var opts = []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	clients := make([]authen_and_post.AuthenticateAndPostClient, 0, len(hosts))
	for _, host := range hosts {
		conn, err := grpc.Dial(host, opts...)
		if err != nil {
			return nil, err
		}
		client := authen_and_post.NewAuthenticateAndPostClient(conn)
		clients = append(clients, client)
	}

	return &aapClient{clients: clients}, nil
}