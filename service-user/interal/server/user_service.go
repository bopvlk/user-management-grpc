package server

import (
	"context"

	"git.foxminded.com.ua/grpc/service-user/interal/domain/mappers"
	"git.foxminded.com.ua/grpc/service-user/interal/repository"
	"git.foxminded.com.ua/grpc/service-user/proto/pb"
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

func (us *userService) FetchUsers(ctx context.Context, fetchUsersRequest *pb.FetchUsersRequest) (*pb.FetchUsersResponse, error) {
	users, totalPages, err := us.storage.FetchUsers(ctx, fetchUsersRequest.Limit, fetchUsersRequest.Page)
	if err != nil {
		return nil, err
	}
	pbUsers := make([]*pb.User, len(users))
	for i := 0; i < len(users); i++ {
		pbUsers[i] = mappers.MapUserToPBUser(&users[i])
	}
	return &pb.FetchUsersResponse{
		Users:      pbUsers,
		TotalPages: int32(totalPages),
	}, nil
}
