package service

import (
	"context"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"test/webook/internal/domain"
	"test/webook/internal/repository"
)

var (
	ErrDuplicateEmail        = repository.ErrDuplicateEmail
	ErrInvalidUserOrPassword = errors.New("用户不存在或者密码不对")
)

type UserService struct {
	userRepo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{userRepo: repo}
}

func (userService *UserService) SignUp(context context.Context, user domain.User) error {
	//加密操作
	hashPwd, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	//加密失败
	if err != nil {
		return err
	}
	user.Password = string(hashPwd)
	err = userService.userRepo.Create(context, user)
	return err
}

/*
用户登录逻辑:

	1.通过唯一信息(邮箱)找到对应的用户(可能为空,需要返回用户不存在)
	2.比较密码是否正确
*/
func (userService *UserService) Login(context context.Context, email, password string) (domain.User, error) {
	//@1 通过邮箱找到用户
	user, err := userService.userRepo.FindByEmail(context, email)
	if err == repository.ErrUserNotFound {
		return domain.User{}, ErrInvalidUserOrPassword
	}
	if err != nil {
		return domain.User{}, err
	}
	//@2 校验密码是否正确
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return domain.User{}, ErrInvalidUserOrPassword
	}
	return user, nil
}
