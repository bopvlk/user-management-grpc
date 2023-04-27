package server

import (
	"context"

	"git.foxminded.com.ua/grpc/grpc-server/interal/domain/mappers"
	"git.foxminded.com.ua/grpc/grpc-server/interal/repository"
	"git.foxminded.com.ua/grpc/grpc-server/proto/pb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type userService struct {
	storage *repository.UserRepository
	pb.UnimplementedApiServiceServer
}

func newUserService(storage *repository.UserRepository) *userService {
	return &userService{
		storage: storage,
	}
}

func (us *userService) Create(ctx context.Context, createRequest *pb.CreateRequest) (*pb.CreateResponse, error) {
	id, err := us.storage.Create(ctx, createRequest.FirstName, createRequest.LastName, createRequest.Email, createRequest.Password)
	if err != nil {
		return nil, err
	}

	return &pb.CreateResponse{
		User: &pb.User{
			ID:        id,
			FirstName: createRequest.FirstName,
			LastName:  createRequest.LastName,
			Email:     createRequest.Email,
			Password:  createRequest.Password,
			CreatedAt: timestamppb.Now(),
			UpdatedAt: timestamppb.Now(),
		},
	}, nil
}

func (us *userService) FetchByEmail(ctx context.Context, fetchByEmailRequest *pb.FetchByEmailRequest) (*pb.FetchByIDResponse, error) {
	u, err := us.storage.FetchByEmail(ctx, fetchByEmailRequest.Email)
	if err != nil {
		return nil, err
	}
	return &pb.FetchByIDResponse{User: mappers.MapUserToPBUser(u)}, nil
}

func (us *userService) FetchByID(ctx context.Context, fetchByIDRequest *pb.FetchByIDRequest) (*pb.FetchByIDResponse, error) {
	u, err := us.storage.FetchByID(ctx, fetchByIDRequest.Id)
	if err != nil {
		return nil, err
	}
	return &pb.FetchByIDResponse{
		User: mappers.MapUserToPBUser(u)}, nil
}
