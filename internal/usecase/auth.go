package usecase

import (
	"context"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"

	"SafeBitesServer/internal/entity"
)

type AuthRepository interface {
	GetByLogin(ctx context.Context, login string) (*entity.User, error)
	GetByID(ctx context.Context, id uuid.UUID) (*entity.User, error)
	Create(ctx context.Context, user *entity.User) error
}

type AuthUsecase struct {
	repo   AuthRepository
	secret []byte
}

func NewAuthUsecase(repo AuthRepository, secret []byte) *AuthUsecase {
	return &AuthUsecase{repo: repo, secret: secret}
}

func (uc *AuthUsecase) Register(ctx context.Context, login, password, name string) (*entity.User, error) {
	_, err := uc.repo.GetByLogin(ctx, login)
	if err == nil {
		return nil, errors.New("пользователь уже существует")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &entity.User{
		ID:           uuid.New(),
		Login:        login,
		PasswordHash: string(hash),
		Name:         name,
	}
	if err := uc.repo.Create(ctx, user); err != nil {
		return nil, errors.New("пользователь уже существует")
	}
	return user, nil
}

func (uc *AuthUsecase) Login(ctx context.Context, login, password string) (string, error) {
	user, err := uc.repo.GetByLogin(ctx, login)
	if err != nil {
		return "", errors.New("неверный логин или пароль")
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)); err != nil {
		return "", errors.New("неверный логин или пароль")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID.String(),
		"exp":     time.Now().Add(7 * 24 * time.Hour).Unix(),
	})
	return token.SignedString(uc.secret)
}

//func (uc *AuthUsecase) Me(ctx context.Context, id uuid.UUID) (*entity.User, error) {
//	return uc.repo.GetByID(ctx, id)
//}
