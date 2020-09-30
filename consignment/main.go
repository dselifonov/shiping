package main

import (
	pb "github.com/dselifonov/shiping/consignment/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"sync"
)

const port = ":50051"

type repository interface {
	Create(*pb.Consignment) (*pb.Consignment, error)
}

type Repository struct {
	mu           sync.Mutex
	consignments []*pb.Consignment
}
