package main

import (
	"context"
	"flag"
	"log"
	"net"

	"github.com/soalno3/Movie/handler"
	"github.com/soalno3/movie-search/Movie/pb"
	"github.com/soalno3/movie-search/Movie/router"
	"google.golang.org/grpc"
)

func main() {
	var (
		gRPCAddr = flag.String("grpc", ":8081",
			"gRPC listen address")
	)
	flag.Parse()
	ctx := context.Background()

	service := handler.MovieSearchService{}

	endpoints := router.Endpoints{
		SearchEndpoint: router.MakeSearchEndpoint(service),
	}

	listener, err := net.Listen("tcp", *gRPCAddr)
	if err != nil {
		log.Println("grpc application failed to start")
		return
	}

	log.Println("grpc application up and running on port 8001")
	handler := router.NewGRPCServer(ctx, endpoints)
	gRPCServer := grpc.NewServer()
	pb.RegisterMovieSearchServer(gRPCServer, handler)
	err = gRPCServer.Serve(listener)
	if err != nil {
		log.Printf("grpc application down. please check. err : %s", err.Error())
	}
}
