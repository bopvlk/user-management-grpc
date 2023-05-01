package mappers

import (
	"fmt"
	"net/http"

	"git.foxminded.com.ua/grpc/grpc-client-api-gateway/interal/apperrors"
	"git.foxminded.com.ua/grpc/grpc-client-api-gateway/interal/domain/requests"
	"git.foxminded.com.ua/grpc/grpc-client-api-gateway/proto/pb"
	"github.com/gin-gonic/gin"
)

func MapAppErrorToErrorResponse(c *gin.Context, err error) {
	appErr := err.(*apperrors.AppError)
	c.JSON(appErr.HTTPCode, gin.H{"error": err.Error()})
}

func MapPBUSERToGetUserResponse(c *gin.Context, u *pb.User) {
	mapedUser := requests.UserResponse{
		ID:        uint(u.ID),
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Email:     u.Email,
		CreatedAt: u.CreatedAt.AsTime(),
		UpdatedAt: u.UpdatedAt.AsTime(),
		DeleteAt: requests.DeleteAt{
			Time:  u.DeleteAt.Time.AsTime(),
			Valid: u.DeleteAt.Valid}}

	c.JSON(http.StatusOK, gin.H{
		"Message": fmt.Sprint("There is user with id", u.ID),
		"User":    mapedUser,
	})
}
