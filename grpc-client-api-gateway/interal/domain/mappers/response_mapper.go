package mappers

import (
	"git.foxminded.com.ua/grpc/grpc-client-api-gateway/interal/apperrors"
	"git.foxminded.com.ua/grpc/grpc-client-api-gateway/interal/domain/models"
	"git.foxminded.com.ua/grpc/grpc-client-api-gateway/interal/domain/requests"
	"git.foxminded.com.ua/grpc/grpc-client-api-gateway/proto/pb"
	"github.com/gin-gonic/gin"
)

func MapAppErrorToErrorResponse(c *gin.Context, err error) {
	appErr := err.(*apperrors.AppError)
	c.JSON(appErr.HTTPCode, gin.H{"error": err.Error()})
	c.Abort()
}

func MapPBUserToGetUserResponse(u *pb.User) *requests.UserResponse {
	return &requests.UserResponse{
		ID:        uint(u.ID),
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Email:     u.Email,
		CreatedAt: u.CreatedAt.AsTime(),
		UpdatedAt: u.UpdatedAt.AsTime(),
		DeleteAt:  u.DeleteAt.Time.AsTime()}
}

func MapPBUsersToPagination(users []*pb.User) *models.Pagination {
	pagination := models.Pagination{}

	usersResponse := make([]*requests.UserResponse, len(users))
	for i := range users {
		usersResponse[i] = MapPBUserToGetUserResponse(users[i])
	}
	pagination.Rows = usersResponse
	return &pagination
}
