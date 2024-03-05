package service

import (
	"errors"
	"github.com/google/uuid"
	"go_fiber_wibowo/app/module/role/repository"
	"go_fiber_wibowo/app/module/role/request"
	"go_fiber_wibowo/app/module/role/response"
	globalRequest "go_fiber_wibowo/app/request"
	"go_fiber_wibowo/utils/paginator"
	"time"
)

// RoleService
type roleService struct {
	role repository.RoleRepository
}

// define interface of IRoleService
//
//go:generate mockgen -destination=role_service_mock.go -package=service . RoleService
type RoleService interface {
	All(req globalRequest.Pagination) (role []*response.Role, paging paginator.Pagination, err error)
	Store(req request.RoleRequest) (err error)
	Show(id string) (role *response.Role, err error)
	Update(id string, req request.RoleRequest) (err error)
	Delete(userId, id string) error
}

// init RoleService
func NewRoleService(role repository.RoleRepository) RoleService {
	return &roleService{
		role: role,
	}
}

func (_i *roleService) All(req globalRequest.Pagination) (role []*response.Role, paging paginator.Pagination, err error) {
	results, paging, err := _i.role.All(req)
	if err != nil {
		return
	}

	for _, result := range results {
		role = append(role, response.FromDomain(result))
	}

	return
}

func (_i *roleService) Store(req request.RoleRequest) (err error) {
	return _i.role.Create(req.ToDomain())
}

func (_i *roleService) Show(id string) (article *response.Role, err error) {
	result, err := _i.role.FindOne(id)
	if err != nil {
		return nil, err
	}

	return response.FromDomain(result), nil
}

func (_i *roleService) Update(id string, req request.RoleRequest) (err error) {
	_, err = _i.role.CheckExist(id)
	if err != nil {
		err = errors.New("role by id:" + id + ", not found")
		return
	}

	return _i.role.Update(id, req.ToDomainUpdate())
}

func (_i *roleService) Delete(userId, id string) error {
	_, err := _i.role.CheckExist(id)
	if err != nil {
		err = errors.New("role by id:" + id + ", not found")
		return err
	}

	req := request.RoleDeleteRequest{}
	req.DeletedAt = time.Now()
	req.DeletedBy = uuid.Must(uuid.Parse(userId))
	req.IsActive = false

	return _i.role.Delete(req, id)
}
