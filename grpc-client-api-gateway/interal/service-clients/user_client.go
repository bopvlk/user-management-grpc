package clients

import (
	"context"
	"time"

	"git.foxminded.com.ua/grpc/grpc-client-api-gateway/interal/apperrors"
	"git.foxminded.com.ua/grpc/grpc-client-api-gateway/interal/config"
	"git.foxminded.com.ua/grpc/grpc-client-api-gateway/proto/pb"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc"
)

type DAO interface {
	CreateUser(ctx context.Context, email, password, firstName, LastName string) (string, error)
	FetchUserByEmail(ctx context.Context, email, password string) (string, error)
	FetchUserByID(ctx context.Context, id uint) (*pb.User, error)
	FetchUsers(ctx context.Context, limit, page int) ([]*pb.User, error)
	UpdateUser(ctx context.Context, user *pb.User) (uint32, error)
	DeleteUser(ctx context.Context, id uint) (uint32, error)
}

type AuthClaims struct {
	jwt.RegisteredClaims
	ID        uint   `json:"id"`
	FirstName string `json:"first_name"`
	Email     string `json:"email"`
}

type UserService struct {
	service pb.ApiServiceClient
	conf    *config.Config
}

func NewUserService(c *config.Config, conn *grpc.ClientConn) DAO {
	return &UserService{
		conf:    c,
		service: pb.NewApiServiceClient(conn)}
}

func (us *UserService) CreateUser(ctx context.Context, firstName, lastName, email, password string) (string, error) {
	hashedPassword, err := hashPassword(password)
	if err != nil {
		return "", apperrors.HashingErr.AppendMessage(err)
	}

	createResponse, err := us.service.Create(ctx, &pb.CreateRequest{FirstName: firstName, LastName: lastName, Email: email, Password: hashedPassword})
	if err != nil {
		return "", apperrors.GRPCErr.AppendMessage(err)
	}

	token, err := us.makeSignedToken(uint(createResponse.User.ID), createResponse.User.FirstName, createResponse.User.Email)
	if err != nil {
		return "", apperrors.CanNotCreateTokenErr.AppendMessage(err)
	}
	return token, nil
}

func (us *UserService) FetchUserByEmail(ctx context.Context, email, password string) (string, error) {
	fetchByIDResponse, err := us.service.FetchByEmail(ctx, &pb.FetchByEmailRequest{Email: email})
	if err != nil {
		return "", apperrors.GRPCErr.AppendMessage(err)
	}

	if !checkPasswordHash(password, fetchByIDResponse.User.Password) {
		return "", &apperrors.WrongPasswordErr
	}

	token, err := us.makeSignedToken(uint(fetchByIDResponse.User.ID), fetchByIDResponse.User.FirstName, fetchByIDResponse.User.Email)
	if err != nil {
		return "", apperrors.CanNotCreateTokenErr.AppendMessage(err)
	}
	return token, nil
}

func (us *UserService) FetchUserByID(ctx context.Context, id uint) (*pb.User, error) {
	fetchByIDResponse, err := us.service.FetchByID(ctx, &pb.FetchByIDRequest{Id: uint32(id)})
	if err != nil {
		return nil, apperrors.GRPCErr.AppendMessage(err)
	}

	return fetchByIDResponse.User, nil
}

func (us *UserService) FetchUsers(ctx context.Context, limit, page int) ([]*pb.User, error) {
	fetchUsersResponse, err := us.service.FetchUsers(ctx, &pb.FetchUsersRequest{Limit: int32(limit), Page: int32(page)})
	if err != nil {
		return nil, apperrors.GRPCErr.AppendMessage(err)
	}

	return fetchUsersResponse.Users, nil
}

func (us *UserService) UpdateUser(ctx context.Context, user *pb.User) (uint32, error) {
	updateResponse, err := us.service.Update(ctx, &pb.UpdateRequest{User: user})
	if err != nil {
		return 0, apperrors.GRPCErr.AppendMessage(err)
	}

	return updateResponse.Id, nil
}

func (us *UserService) DeleteUser(ctx context.Context, id uint) (uint32, error) {
	deleteResponse, err := us.service.Delete(ctx, &pb.DeleteRequest{Id: uint32(id)})
	if err != nil {
		return 0, apperrors.GRPCErr.AppendMessage(err)
	}

	return deleteResponse.Id, nil
}

func (us *UserService) makeSignedToken(id uint, firstName, email string) (string, error) {
	claims := AuthClaims{
		ID:        id,
		FirstName: firstName,
		Email:     email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Second * (time.Duration(us.conf.TokenTtl)))),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(us.conf.SigningKey))
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
