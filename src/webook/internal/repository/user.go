package repository

import (
	"context"
	"test/webook/internal/domain"
	"test/webook/internal/repository/dao"
)

var (
	ErrDuplicateEmail = dao.ErrDuplicateEmail
	ErrUserNotFound   = dao.ErrRecordNotFound
)

type UserRepository struct {
	dao *dao.UserDAO
}

func NewUserRepository(dao *dao.UserDAO) *UserRepository {
	return &UserRepository{
		dao: dao,
	}
}

func (userRepo *UserRepository) Create(context context.Context, user domain.User) error {
	return userRepo.dao.Insert(context, dao.User{
		Email:    user.Email,
		Password: user.Password,
	})
}

func (userRepo *UserRepository) FindByEmail(context context.Context, email string) (domain.User, error) {
	u, err := userRepo.dao.FindByEmail(context, email)
	if err != nil {
		return domain.User{}, err
	}
	return domain.User{
		Id:       u.Id,
		Email:    u.Email,
		Password: u.Password,
	}, nil
}
