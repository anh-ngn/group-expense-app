// internal/user/service.go
package user

import (
	"context"
	"crypto/sha256"
	"database/sql"
	"fmt"
	"path/to/internal/user/repository"

	pb "github.com/anh-ngn/group-expense-app/user_service/api/user"
)

type UserService struct {
	repo repository.Querier
}

func NewUserService(repo repository.Querier) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) RegisterWithEmail(ctx context.Context, req *pb.RegisterWithEmailRequest) (*pb.RegisterWithEmailResponse, error) {
	passwordHash := fmt.Sprintf("%x", sha256.Sum256([]byte(req.Password)))
	user, err := s.repo.CreateUser(ctx, repository.CreateUserParams{
		Email:        req.Email,
		PasswordHash: sql.NullString{String: passwordHash, Valid: true},
	})
	if err != nil {
		return nil, err
	}
	return &pb.RegisterWithEmailResponse{UserId: fmt.Sprintf("%d", user.ID)}, nil
}

func (s *UserService) LoginWithEmail(ctx context.Context, req *pb.LoginWithEmailRequest) (*pb.LoginWithEmailResponse, error) {
	user, err := s.repo.GetUserByEmail(ctx, req.Email)
	if err != nil {
		return nil, err
	}
	passwordHash := fmt.Sprintf("%x", sha256.Sum256([]byte(req.Password)))
	if user.PasswordHash.String != passwordHash {
		return nil, fmt.Errorf("invalid credentials")
	}
	return &pb.LoginWithEmailResponse{UserId: fmt.Sprintf("%d", user.ID), Email: user.Email}, nil
}

func (s *UserService) LoginWithGoogle(ctx context.Context, req *pb.LoginWithGoogleRequest) (*pb.LoginWithGoogleResponse, error) {
	user, err := s.repo.GetUserByGoogleID(ctx, req.GoogleId)
	if err == sql.ErrNoRows {
		user, err = s.repo.CreateUserWithGoogle(ctx, repository.CreateUserWithGoogleParams{
			Email:    req.Email,
			GoogleID: sql.NullString{String: req.GoogleId, Valid: true},
		})
		if err != nil {
			return nil, err
		}
	} else if err != nil {
		return nil, err
	}
	return &pb.LoginWithGoogleResponse{UserId: fmt.Sprintf("%d", user.ID), Email: user.Email}, nil
}
