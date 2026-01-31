package repository

import (
	"github.com/google/uuid"
	"go_fiber_wibowo/app/database/schema"
	"go_fiber_wibowo/app/module/permission/request"
	paginationRequest "go_fiber_wibowo/app/request"
	"go_fiber_wibowo/internal/bootstrap/database"
	"go_fiber_wibowo/utils/config"
	"go_fiber_wibowo/utils/paginator"
)

type permissionRepository struct {
	DB  *database.Database
	Cfg *config.Config
}

//go:generate mockgen -destination=permission_repository_mock.go -package=repository . PermissionRepository
type PermissionRepository interface {
	All(req paginationRequest.Pagination) (permission []*schema.Permission, paging paginator.Pagination, err error)
	FindOne(id string) (permission *schema.Permission, err error)
	CheckExist(id string) (permission *schema.Permission, err error)
	Create(permission *schema.Permission) (err error)
	Update(id string, permission *schema.Permission) (err error)
	Delete(req request.PermissionDeleteRequest, id string) (err error)
	FindByUserId(Id string) (user []*schema.Permission, err error)
}

func NewPermissionRepository(db *database.Database) PermissionRepository {
	return &permissionRepository{
		DB: db,
	}
}

func (_i *permissionRepository) All(req paginationRequest.Pagination) (permission []*schema.Permission, paging paginator.Pagination, err error) {
	var count int64

	query := _i.DB.DB.Model(&schema.Permission{})
	query.Count(&count)

	req.Pagination.Count = count
	req.Pagination = paginator.Paging(req.Pagination)

	err = query.Offset(req.Pagination.Offset).Limit(req.Pagination.Limit).Find(&permission).Error
	if err != nil {
		return
	}

	paging = *req.Pagination

	return
}

func (_i *permissionRepository) FindOne(id string) (permission *schema.Permission, err error) {
	if err := _i.DB.DB.First(&permission, uuid.Must(uuid.Parse(id))).Error; err != nil {
		return nil, err
	}

	return permission, nil
}

func (_i *permissionRepository) CheckExist(id string) (permission *schema.Permission, err error) {
	if err := _i.DB.DB.Select("id").First(&permission, uuid.Must(uuid.Parse(id))).Error; err != nil {
		return nil, err
	}

	return permission, nil
}

func (_i *permissionRepository) Create(permission *schema.Permission) (err error) {
	return _i.DB.DB.Create(permission).Error
}

func (_i *permissionRepository) Update(id string, permission *schema.Permission) (err error) {
	return _i.DB.DB.Table("permissions").
		Where(&schema.Permission{Id: uuid.Must(uuid.Parse(id))}).
		Updates(permission).Error
}

func (_i *permissionRepository) Delete(req request.PermissionDeleteRequest, id string) error {
	if err := _i.DB.DB.Table("permissions").Where("id", id).Updates(map[string]interface{}{"is_active": req.IsActive, "deleted_at": req.DeletedAt, "deleted_by": req.DeletedBy}).Error; err != nil {
		return err
	}

	return nil
}

func (_i *permissionRepository) FindByUserId(Id string) (res []*schema.Permission, err error) {
	qState := "SELECT * FROM permissions "
	qState = qState + "WHERE created_by = '" + Id + "' "

	err = _i.DB.DB.Raw(qState).Scan(&res).Error
	if err != nil {
		return nil, err
	}

	return res, nil
}
