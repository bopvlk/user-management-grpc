package httpserver

import (
	"fmt"
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
		err = apperrors.CanNotBindErr.AppendMessage(err)
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
		err = apperrors.CanNotBindErr.AppendMessage(err)
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

	c.JSON(http.StatusOK, gin.H{
		"Message": fmt.Sprint("There is user with id: ", user.ID),
		"User":    mappers.MapPBUserToGetUserResponse(user),
	})
}

func (s *httpServer) Pagination(c *gin.Context) {
	var paginationRequest requests.PaginationRequest
	if err := c.Bind(&paginationRequest); err != nil {
		err = apperrors.CanNotBindErr.AppendMessage(err)
		s.log.err.Print(err)
		mappers.MapAppErrorToErrorResponse(c, err)
	}

	users, totalPages, err := s.userDAO.FetchUsers(c.Request.Context(), paginationRequest.Limit, paginationRequest.Page)
	if err != nil {
		s.log.err.Print(err)
		mappers.MapAppErrorToErrorResponse(c, err)
		return
	}

	pagination := mappers.MapPBUsersToPagination(users)
	pagination.TotalPages = totalPages
	pagination.Limit = paginationRequest.Limit
	pagination.Page = paginationRequest.Page
	urlPath := c.Request.URL.Path
	pagination.FirstPage = fmt.Sprintf("%s?limit=%d&page=%d&sort=%s", urlPath, pagination.Limit, 1, pagination.Sort)
	pagination.LastPage = fmt.Sprintf("%s?limit=%d&page=%d&sort=%s", urlPath, pagination.Limit, pagination.TotalPages, pagination.Sort)

	if pagination.Page > 1 {
		pagination.PreviousPage = fmt.Sprintf("%s?limit=%d&page=%d&sort=%s", urlPath, pagination.Limit, pagination.Page-1, pagination.Sort)
	}

	if pagination.Page < pagination.TotalPages {
		pagination.NextPage = fmt.Sprintf("%s?limit=%d&page=%d&sort=%s", urlPath, pagination.Limit, pagination.Page+1, pagination.Sort)
	}

	c.JSON(http.StatusOK, gin.H{
		"Message":    "Pagination success",
		"Pagination": pagination,
	})
}

func (s *httpServer) UpdateUser(c *gin.Context) {
	var uReq requests.UpdateUsersRequest

	if err := c.Bind(&uReq); err != nil {
		err = apperrors.CanNotBindErr.AppendMessage(err)
		s.log.err.Print(err)
		mappers.MapAppErrorToErrorResponse(c, err)
	}

	id, err := s.userDAO.UpdateUser(c.Request.Context(), uReq.FirstName, uReq.LastName, uReq.Email, uReq.Password)
	if err != nil {
		s.log.err.Print(err)
		mappers.MapAppErrorToErrorResponse(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"Message": fmt.Sprintf("User with id: %d updated", id),
	})
}

func (s *httpServer) DeleteUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		err = apperrors.CanNotBindErr.AppendMessage(err)
		s.log.err.Print(err)
		mappers.MapAppErrorToErrorResponse(c, err)
	}

	respId, err := s.userDAO.DeleteUser(c.Request.Context(), uint(id))
	if err != nil {
		s.log.err.Print(err)
		mappers.MapAppErrorToErrorResponse(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"Message": fmt.Sprintf("User with id: %d deleted", respId),
	})
}
