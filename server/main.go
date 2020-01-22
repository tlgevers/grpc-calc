package main

import (
    "fmt"
    "google.golang.org/grpc"
    pb "github.com/tlgevers/grpc-calc"
    "context"
    "log"
    "net"
)

const (
	port = ":50051"
)

// server is used to implement calc.CalcServer
type server struct {
    pb.UnimplementedCalcServer
}

// Sum implements calc.CalcServer
func (s *server) Sum(ctx context.Context, in *pb.CalcRequest) (*pb.CalcResult, error) {
    log.Printf("Recieved: %v", in.argA, in.argB)
    return &pb.CalcResult{result: in.argA + in.argB}, nil
}

func main() {
    fmt.Println("Starting server")
    lis, err := net.Listen("tcp", port)
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }
    s := grpc.NewServer()
    pb.RegisterCalcServer(s, &server{})
    if err := s.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}
