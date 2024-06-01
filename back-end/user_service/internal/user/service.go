package user

import (
	"context"
	"crypto/sha256"
	"database/sql"
	"fmt"

	pb "github.com/anh-ngn/group-expense-app/user_service/api/user"
	"github.com/anh-ngn/group-expense-app/user_service/internal/user/repository"
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
	return &pb.LoginWithEmailResponse{UserId: fmt.Sprintf("%d", user.ID), Email: user.Email, AvatarUrl: user.AvatarUrl}, nil
}

func (s *UserService) LoginWithGoogle(ctx context.Context, req *pb.LoginWithGoogleRequest) (*pb.LoginWithGoogleResponse, error) {
	user, err := s.repo.GetUserByGoogleID(ctx, req.GoogleId)
	if err == sql.ErrNoRows {
		user, err = s.repo.CreateUserWithGoogle(ctx, repository.CreateUserWithGoogleParams{
			Email:     req.Email,
			GoogleID:  sql.NullString{String: req.GoogleId, Valid: true},
			AvatarUrl: req.AvatarUrl,
		})
		if err != nil {
			return nil, err
		}
	} else if err != nil {
		return nil, err
	}
	return &pb.LoginWithGoogleResponse{UserId: fmt.Sprintf("%d", user.ID), Email: user.Email, AvatarUrl: user.AvatarUrl}, nil
}

func (s *UserService) LoginWithApple(ctx context.Context, req *pb.LoginWithAppleRequest) (*pb.LoginWithAppleResponse, error) {
	user, err := s.repo.GetUserByAppleID(ctx, req.AppleId)
	if err == sql.ErrNoRows {
		user, err = s.repo.CreateUserWithApple(ctx, repository.CreateUserWithAppleParams{
			Email:     req.Email,
			AppleID:   sql.NullString{String: req.AppleId, Valid: true},
			AvatarUrl: req.AvatarUrl,
		})
		if err != nil {
			return nil, err
		}
	} else if err != nil {
		return nil, err
	}
	return &pb.LoginWithAppleResponse{UserId: fmt.Sprintf("%d", user.ID), Email: user.Email, AvatarUrl: user.AvatarUrl}, nil
}

func (s *UserService) LoginWithMicrosoft(ctx context.Context, req *pb.LoginWithMicrosoftRequest) (*pb.LoginWithMicrosoftResponse, error) {
	user, err := s.repo.GetUserByMicrosoftID(ctx, req.MicrosoftId)
	if err == sql.ErrNoRows {
		user, err = s.repo.CreateUserWithMicrosoft(ctx, repository.CreateUserWithMicrosoftParams{
			Email:       req.Email,
			MicrosoftID: sql.NullString{String: req.MicrosoftId, Valid: true},
			AvatarUrl:   req.AvatarUrl,
		})
		if err != nil {
			return nil, err
		}
	} else if err != nil {
		return nil, err
	}
	return &pb.LoginWithMicrosoftResponse{UserId: fmt.Sprintf("%d", user.ID), Email: user.Email, AvatarUrl: user.AvatarUrl}, nil
}
