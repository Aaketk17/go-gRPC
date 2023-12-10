package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/Aaketk17/go+gRPC/pb"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type LaptopServer struct {
	laptopStore LaptopStore
}

func NewLaptopServer(laptopStore LaptopStore) *LaptopServer {
	return &LaptopServer{
		laptopStore: laptopStore,
	}
}

func (service *LaptopServer) CreateLaptop(ctx context.Context, req *pb.CreateLaptopRequest) (*pb.CreateLaptopResponse, error) {
	laptop := req.GetLaptop()
	fmt.Println("Received a create laptop request")

	if len(laptop.Id) > 0 {
		_, err := uuid.Parse(laptop.Id)
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "Given UUID is not a valid value: %v", err)
		}
	} else {
		uuid, err := uuid.NewRandom()
		if err != nil {
			return nil, status.Errorf(codes.Internal, "Error in creating new Laptop ID: %v", err)
		}
		laptop.Id = uuid.String()
	}

	err := service.laptopStore.Save(laptop)
	if err != nil {
		code := codes.Internal
		if errors.Is(err, ErrAlreadyExists) {
			code = codes.AlreadyExists
		}
		return nil, status.Errorf(code, "Unable to save the laptop to the store: %v", err)
	}
	fmt.Printf("Saved the laptop with ID: %s", laptop.Id)

	res := &pb.CreateLaptopResponse{
		Id: laptop.Id,
	}

	return res, nil
}
