package main

import (
	"context"
	"encoding/json"
	pb "github.com/dselifonov/shiping/consignment/proto"
	"github.com/micro/go-micro/v2"
	"io/ioutil"
	"log"
	"os"
)

const (
	defaultFilename = "consignment.json"
)

func parseFile(file string) (*pb.Consignment, error) {
	var consignment *pb.Consignment
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, &consignment)
	if err != nil {
		return nil, err
	}

	return consignment, err
}

func main() {
	service := micro.NewService(micro.Name("cli"))
	service.Init()
	client := pb.NewShippingService("consignment", service.Client())

	// Contact the server and print out its response.
	file := defaultFilename
	if len(os.Args) > 1 {
		file = os.Args[1]
	}

	consignment, err := parseFile(file)

	if err != nil {
		log.Fatalf("Could not parse file: %v", err)
	}

	r, err := client.CreateConsignment(context.Background(), consignment)
	if err != nil {
		log.Fatalf("Could not greet: %v", err)
	}
	log.Printf("Created: %t", r.Created)

	getAll, err := client.GetConsignments(context.Background(), &pb.GetRequest{})
	if err != nil {
		log.Fatalf("Could not list consignments: %v", err)
	}
	for _, v := range getAll.Consignments {
		log.Println(v)
	}
}
