package mappers

import (
	"database/sql"

	"git.foxminded.com.ua/grpc/grpc-server/interal/domain/models"
	"git.foxminded.com.ua/grpc/grpc-server/proto/pb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func MapPBUserToUser(pu *pb.User) *models.User {
	return &models.User{
		ID:        uint(pu.ID),
		FirstName: pu.FirstName,
		LastName:  pu.LastName,
		Email:     pu.Email,
		CreatedAt: pu.CreatedAt.AsTime(),
		UpdatedAt: pu.UpdatedAt.AsTime(),
		DeleteAt: sql.NullTime{
			Time:  pu.DeleteAt.Time.AsTime(),
			Valid: pu.DeleteAt.Valid,
		},
	}
}

func MapUserToPBUser(u *models.User) *pb.User {
	return &pb.User{
		ID:        uint32(u.ID),
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Email:     u.Email,
		CreatedAt: timestamppb.New(u.CreatedAt),
		UpdatedAt: timestamppb.New(u.UpdatedAt),
		DeleteAt: &pb.User_DeleteAt{
			Time:  timestamppb.New(u.DeleteAt.Time),
			Valid: u.DeleteAt.Valid,
		}}
}
