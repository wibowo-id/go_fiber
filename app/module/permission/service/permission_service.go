package service

import (
	"errors"
	"github.com/google/uuid"
	"go_fiber_wibowo/app/database/schema"
	"go_fiber_wibowo/app/module/permission/repository"
	"go_fiber_wibowo/app/module/permission/request"
	"go_fiber_wibowo/app/module/permission/response"
	globalRequest "go_fiber_wibowo/app/request"
	"go_fiber_wibowo/utils/paginator"
	"time"
)

// PermissionService
type permissionService struct {
	permission repository.PermissionRepository
}

// define interface of IPermissionService
//
//go:generate mockgen -destination=permission_service_mock.go -package=service . PermissionService
type PermissionService interface {
	All(req globalRequest.Pagination) (role []*schema.Permission, paging paginator.Pagination, err error)
	Store(req request.PermissionRequest) (err error)
	Show(id string) (role *schema.Permission, err error)
	Update(id string, req request.PermissionRequest) (err error)
	Delete(userId, id string) error
}

// init PermissionService
func NewPermissionService(permission repository.PermissionRepository) PermissionService {
	return &permissionService{
		permission: permission,
	}
}

func (_i *permissionService) All(req globalRequest.Pagination) (role []*schema.Permission, paging paginator.Pagination, err error) {
	results, paging, err := _i.permission.All(req)
	if err != nil {
		return
	}

	for _, result := range results {
		role = append(role, response.FromDomain(result))
	}

	return
}

func (_i *permissionService) Store(req request.PermissionRequest) (err error) {
	return _i.permission.Create(req.ToDomain())
}

func (_i *permissionService) Show(id string) (article *schema.Permission, err error) {
	result, err := _i.permission.FindOne(id)
	if err != nil {
		return nil, err
	}

	return response.FromDomain(result), nil
}

func (_i *permissionService) Update(id string, req request.PermissionRequest) (err error) {
	_, err = _i.permission.CheckExist(id)
	if err != nil {
		err = errors.New("role by id:" + id + ", not found")
		return
	}

	return _i.permission.Update(id, req.ToDomainUpdate())
}

func (_i *permissionService) Delete(userId, id string) error {
	_, err := _i.permission.CheckExist(id)
	if err != nil {
		err = errors.New("role by id:" + id + ", not found")
		return err
	}

	req := request.PermissionDeleteRequest{}
	req.DeletedAt = time.Now()
	req.DeletedBy = uuid.Must(uuid.Parse(userId))
	req.IsActive = false

	return _i.permission.Delete(req, id)
}
