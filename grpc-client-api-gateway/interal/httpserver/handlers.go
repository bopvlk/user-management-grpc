package httpserver

import (
	"net/http"
	"strconv"

	"git.foxminded.com.ua/grpc/grpc-client-api-gateway/interal/apperrors"
	"git.foxminded.com.ua/grpc/grpc-client-api-gateway/interal/domain/mappers"
	"git.foxminded.com.ua/grpc/grpc-client-api-gateway/interal/domain/requests"
	"github.com/gin-gonic/gin"
)

func (s *httpServer) SignUp(c *gin.Context) {
	var signUpRequest requests.SignUpRequest
	if err := c.ShouldBindJSON(&signUpRequest); err != nil {
		s.log.err.Print(err)
		mappers.MapAppErrorToErrorResponse(c, err)
		return
	}

	token, err := s.userDAO.CreateUser(c.Request.Context(), signUpRequest.FirstName, signUpRequest.LastName, signUpRequest.Email, signUpRequest.Password)
	if err != nil {
		s.log.err.Print(err)
		mappers.MapAppErrorToErrorResponse(c, err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Hello in restricted zone!",
		"token":   token})
}

func (s *httpServer) SignIn(c *gin.Context) {
	var signInRequest requests.SignInRequest
	if err := c.ShouldBindJSON(&signInRequest); err != nil {
		s.log.err.Print(err)
		mappers.MapAppErrorToErrorResponse(c, err)
		return
	}

	token, err := s.userDAO.FetchUserByEmail(c.Request.Context(), signInRequest.Email, signInRequest.Password)
	if err != nil {
		s.log.err.Print(err)
		mappers.MapAppErrorToErrorResponse(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Hello in restricted zone!",
		"token":   token})
}

func (s *httpServer) GetOneUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		err = apperrors.CanNotBindErr.AppendMessage(err)
		s.log.err.Print(err)
		mappers.MapAppErrorToErrorResponse(c, err)
	}

	user, err := s.userDAO.FetchUserByID(c.Request.Context(), uint(id))
	if err != nil {
		s.log.err.Print(err)
		mappers.MapAppErrorToErrorResponse(c, err)
		return
	}

	mappers.MapPBUSERToGetUserResponse(c, user)
}
