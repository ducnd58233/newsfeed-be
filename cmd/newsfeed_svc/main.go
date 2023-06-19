package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/ducnd58233/newsfeed-be/configs"
	"github.com/ducnd58233/newsfeed-be/internal/app/newsfeed_svc"
	"github.com/ducnd58233/newsfeed-be/pkg/types/proto/pb/newsfeed"
	"google.golang.org/grpc"
)

var path = flag.String("conf", "config.yml", "config path for this service")

func main() {
	conf, err := configs.GetNewsfeedConfig(*path)
	if err != nil {
		log.Fatalf("failed to parse config: %v\n", err)
	}
	service, err := newsfeed_svc.NewNewsfeedService(conf)
	if err != nil {
		log.Fatalf("failed to init service %v\n", err)
	}
	lis, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", conf.Port))
	if err != nil {
		log.Fatalf("failed to listen %v\n", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	newsfeed.RegisterNewsfeedServer(grpcServer, service)
	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatalf("server stopped %v", err)
	}
	log.Println("newsfeed server started.")
}
