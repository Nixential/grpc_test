package main

import (
	"context"
	"log"
	"os"
	"sync"
	"time"

	pb "grpc_test/helloworld"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const numRequests = 10000 // Number of requests to make concurrently

func main() {
	logFile, err := os.Create("grpc_client.log")
	if err != nil {
		log.Fatalf("Failed to create log file: %s", err)
	}
	defer logFile.Close()

	// Redirect log output to the log file
	log.SetOutput(logFile)

	var wg sync.WaitGroup

	for i := 0; i < numRequests; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			measureLatency()
		}()
	}

	wg.Wait()
}

func measureLatency() {
	conn, err := grpc.Dial("127.0.0.1:5000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("Failed to connect to gRPC server: %s", err)
		return
	}
	defer conn.Close()

	c := pb.NewHelloWorldServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*60)
	defer cancel()

	startTime := time.Now()

	r, err := c.SayHello(ctx, &pb.HelloWorldRequest{})
	if err != nil {
		log.Fatalf("Error calling SayHello: %s", err)
	}

	endTime := time.Now()
	latency := endTime.Sub(startTime)

	log.Printf("Response from gRPC server's SayHello function: %s", r.GetMessage())
	log.Printf("Latency: %v", latency)
}
