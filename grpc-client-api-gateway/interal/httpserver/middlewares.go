package httpserver

import (
	"errors"
	"fmt"
	"strings"

	"git.foxminded.com.ua/grpc/grpc-client-api-gateway/interal/apperrors"
	"git.foxminded.com.ua/grpc/grpc-client-api-gateway/interal/domain/mappers"
	clients "git.foxminded.com.ua/grpc/grpc-client-api-gateway/interal/service-clients"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func (s *httpServer) jwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		strToken := extractToken(c)
		if strToken == "" {
			mappers.MapAppErrorToErrorResponse(c, &apperrors.ExtractTokenErr)
		}
		_, _, _, err := s.parseToken(strToken)
		if err != nil {
			mappers.MapAppErrorToErrorResponse(c, apperrors.ParseTokenErr.AppendMessage(err))
		}
		c.Next()
	}
}

func extractToken(c *gin.Context) string {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		return ""
	}

	headerParts := strings.Split(authHeader, " ")
	if len(headerParts) != 2 {
		return ""
	}

	if headerParts[0] != "Bearer" {
		return ""
	}

	return headerParts[1]
}

func (s *httpServer) parseToken(strToken string) (id uint, firstname, email string, err error) {
	token, err := jwt.ParseWithClaims(strToken, &clients.AuthClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(s.conf.SigningKey), nil
	})

	if err != nil {
		return 0, "", "", err
	}

	if claims, ok := token.Claims.(*clients.AuthClaims); ok && token.Valid {
		return claims.ID, claims.FirstName, claims.Email, nil
	}

	return 0, "", "", errors.New("invalid access token")
}
