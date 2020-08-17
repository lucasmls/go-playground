package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/google/uuid"
	"github.com/lucasmls/go-playground/grpc-poc/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// Task ...
type Task struct{}

// Create ...
func (t *Task) Create(ctx context.Context, req *pb.CreateTaskPayload) (*pb.CreateTaskResponse, error) {
	fmt.Printf("Request to create a task with title: %s and status: %v\n", req.Task.Title, req.Task.Done)

	return &pb.CreateTaskResponse{
		TaskId: uuid.New().String(),
	}, nil
}

func main() {
	grpcServer := grpc.NewServer()
	pb.RegisterTaskServiceServer(grpcServer, &Task{})
	reflection.Register(grpcServer)

	listener, err := net.Listen("tcp", ":50055")
	if err != nil {
		log.Fatalf("Cannot start GRPC Server: %v", err)
	}

	grpcServer.Serve(listener)
}
