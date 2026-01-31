package repository

import (
	"github.com/google/uuid"
	"go_fiber_wibowo/app/database/schema"
	"go_fiber_wibowo/app/module/role/request"
	"go_fiber_wibowo/app/module/role/response"
	paginationRequest "go_fiber_wibowo/app/request"
	"go_fiber_wibowo/internal/bootstrap/database"
	"go_fiber_wibowo/utils/config"
	"go_fiber_wibowo/utils/paginator"
)

type roleRepository struct {
	DB  *database.Database
	Cfg *config.Config
}

//go:generate mockgen -destination=role_repository_mock.go -package=repository . RoleRepository
type RoleRepository interface {
	All(req paginationRequest.Pagination) (role []*response.Role, paging paginator.Pagination, err error)
	FindOne(id string) (role *response.Role, err error)
	CheckExist(id string) (role *schema.Role, err error)
	Create(role *schema.Role) (err error)
	Update(id string, role *schema.Role) (err error)
	Delete(req request.RoleDeleteRequest, id string) (err error)
	FindByUserId(Id string) (user []*schema.Role, err error)
}

func NewRoleRepository(db *database.Database) RoleRepository {
	return &roleRepository{
		DB: db,
	}
}

func (_i *roleRepository) All(req paginationRequest.Pagination) (role []*response.Role, paging paginator.Pagination, err error) {
	var count int64

	query := _i.DB.DB.Model(&schema.Role{})
	query.Count(&count)

	req.Pagination.Count = count
	req.Pagination = paginator.Paging(req.Pagination)

	err = query.Offset(req.Pagination.Offset).Limit(req.Pagination.Limit).Find(&role).Error
	if err != nil {
		return
	}

	paging = *req.Pagination

	return
}

func (_i *roleRepository) FindOne(id string) (role *response.Role, err error) {
	if err := _i.DB.DB.First(&role, uuid.Must(uuid.Parse(id))).Error; err != nil {
		return nil, err
	}

	return role, nil
}

func (_i *roleRepository) CheckExist(id string) (role *schema.Role, err error) {
	if err := _i.DB.DB.Select("id").First(&role, uuid.Must(uuid.Parse(id))).Error; err != nil {
		return nil, err
	}

	return role, nil
}

func (_i *roleRepository) Create(role *schema.Role) (err error) {
	return _i.DB.DB.Create(role).Error
}

func (_i *roleRepository) Update(id string, role *schema.Role) (err error) {
	return _i.DB.DB.Table("roles").
		Where(&schema.Role{Id: uuid.Must(uuid.Parse(id))}).
		Updates(role).Error
}

func (_i *roleRepository) Delete(req request.RoleDeleteRequest, id string) error {
	if err := _i.DB.DB.Table("roles").Where("id", id).Updates(map[string]interface{}{"is_active": req.IsActive, "deleted_at": req.DeletedAt, "deleted_by": req.DeletedBy}).Error; err != nil {
		return err
	}

	return nil
}

func (_i *roleRepository) FindByUserId(Id string) (res []*schema.Role, err error) {
	qState := "SELECT * FROM roles "
	qState = qState + "WHERE created_by = '" + Id + "' "

	err = _i.DB.DB.Raw(qState).Scan(&res).Error
	if err != nil {
		return nil, err
	}

	return res, nil
}
