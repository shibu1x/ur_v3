package main

import (
	"context"
	"log"
	"os"
	"time"

	pb "github.com/shibu1x/ur_v3/pkg/grpc"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial(os.Getenv("HOST"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewRoomServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := &pb.GetRoomsByPrefCodeRequest{
		PrefCode: "tokyo",
	}

	res, err := c.GetRoomsByPrefCode(ctx, req)
	if err != nil {
		log.Fatalf("could not get rooms: %v", err)
	}

	for _, room := range res.Rooms {
		log.Println(room)
	}
}
