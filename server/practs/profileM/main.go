package main

import (
	"fmt"
	"log"
	"net"

	"practs/profileM/config"
	"practs/profileM/db"
	"practs/profileM/pb"
	"practs/profileM/src"

	"google.golang.org/grpc"
)

func main() {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	h := db.Init(c.DBUrl)

	lis, err := net.Listen("tcp", c.Port)

	if err != nil {
		log.Fatalln("Failed to listing:", err)
	}

	fmt.Println("Profile Svc on", c.Port)

	s := src.Server{
		H: h,
	}

	grpcServer := grpc.NewServer()

	pb.RegisterProfileServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln("Failed to serve:", err)
	}
}
