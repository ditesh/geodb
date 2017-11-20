package main

import (
	"context"
	"log"

	"geodb/api"

	"google.golang.org/grpc"
)

const (
	address = "localhost:12345"
)

func main() {

	p := &api.Point{}

	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())

	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer func() {
		err := conn.Close()

		if err != nil {
			log.Fatal("unable to close connection")
		}
	}()

	c := api.NewAPIClient(conn)
	_, err = c.Write(context.Background(), &api.WriteRequest{P: p, Blob: "this is the potato"})

	if err != nil {
		log.Fatalf("could not connect: (%v)", err)
	}

}
