package service

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"go_fiber_wibowo/app/module/user/repository"
	"go_fiber_wibowo/app/module/user/request"
	"go_fiber_wibowo/app/module/user/response"
	globalRequest "go_fiber_wibowo/app/request"
	"go_fiber_wibowo/utils/paginator"
	"time"
)

// UserService
type userService struct {
	userRepo repository.UserRepository
}

// define interface of IUserService
//
//go:generate mockgen -destination=user_service_mock.go -package=service . UserService
type UserService interface {
	All(req globalRequest.Pagination, c *fiber.Ctx) (user []*response.UserResponse, paging paginator.Pagination, err error)
	Show(id string) (article *response.UserResponse, err error)
	Store(req request.UserRequest) (err error)
	Update(id string, req request.UserRequest) (err error)
	Delete(userId, id string) error
}

// init UserService
func NewUserService(user repository.UserRepository) UserService {
	return &userService{
		userRepo: user,
	}
}

func (_i *userService) All(req globalRequest.Pagination, c *fiber.Ctx) (user []*response.UserResponse, paging paginator.Pagination, err error) {
	results, paging, err := _i.userRepo.All(req, c)
	if err != nil {
		return
	}

	for _, result := range results {
		user = append(user, response.FromDomain(result))
	}

	return
}

func (_i *userService) Store(req request.UserRequest) (err error) {
	return _i.userRepo.Create(req.ToDomain())
}

func (_i *userService) Show(id string) (article *response.UserResponse, err error) {
	result, err := _i.userRepo.FindOne(id)
	if err != nil {
		return nil, err
	}

	return response.FromDomain(result), nil
}

func (_i *userService) Update(id string, req request.UserRequest) (err error) {
	_, err = _i.userRepo.CheckExist(id)
	if err != nil {
		err = errors.New("user by id:" + id + ", not found")
		return err
	}

	return _i.userRepo.Update(id, req.ToDomainUpdate())
}

func (_i *userService) Delete(userId, id string) error {
	_, err := _i.userRepo.CheckExist(id)
	if err != nil {
		err = errors.New("user by id:" + id + ", not found")
		return err
	}

	req := request.UserDeleteRequest{}
	req.DeletedAt = time.Now()
	req.DeletedBy = uuid.Must(uuid.Parse(userId))
	req.IsActive = false

	return _i.userRepo.Delete(req, id)
}
