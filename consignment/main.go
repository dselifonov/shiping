package main

import (
	"context"
	pb "github.com/dselifonov/shiping/consignment/proto"
	pbVessel "github.com/dselifonov/shiping/vessel/proto"
	"github.com/micro/go-micro/v2"
	"log"
)

// Repo Code
type repository interface {
	Create(consignment *pb.Consignment) (*pb.Consignment, error)
	GetAll() []*pb.Consignment
}

type Repository struct {
	consignments []*pb.Consignment
}

func (repo *Repository) Create(consignment *pb.Consignment) (*pb.Consignment, error) {
	repo.consignments = append(repo.consignments, consignment)
	return consignment, nil
}

func (repo *Repository) GetAll() []*pb.Consignment {
	return repo.consignments
}

// Service Code
type consignmentService struct {
	repo repository
}

func (s *consignmentService) CreateConsignment(ctx context.Context, req *pb.Consignment, res *pb.Response) error {
	consignment, err := s.repo.Create(req)
	if err != nil {
		return err
	}
	res.Created = true
	res.Consignment = consignment
	return nil
}

func (s *consignmentService) GetConsignments(ctx context.Context, req *pb.GetRequest, res *pb.Response) error {
	consignments := s.repo.GetAll()
	res.Consignments = consignments
	return nil
}

// Main func
func main() {
	repo := &Repository{}
	service := micro.NewService(micro.Name("consignment"))
	service.Init()

	vesselClient := pbVessel.NewVesselService("vessel.client", service.Client())

	if err := pb.RegisterShippingServiceHandler(service.Server(), &consignmentService{repo}); err != nil {
		log.Panic(err)
	}

	if err := service.Run(); err != nil {
		log.Panic(err)
	}
}
