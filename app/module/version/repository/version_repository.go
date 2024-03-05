package repository

import (
	"github.com/google/uuid"
	"go_fiber_wibowo/app/database/schema"
	"go_fiber_wibowo/app/module/version/request"
	"go_fiber_wibowo/internal/bootstrap/database"
	"go_fiber_wibowo/utils/paginator"
)

type versionRepository struct {
	DB *database.Database
}

//go:generate mockgen -destination=version_repository_mock.go -package=repository . VersionRepository
type VersionRepository interface {
	GetVersions(req request.VersionsRequest) (versions []*schema.Version, paging paginator.Pagination, err error)
	GetOneVersion() (versions *schema.Version, err error)
	FindOne(id uuid.UUID) (version *schema.Version, err error)
	Create(version *schema.Version) (err error)
	Update(id uuid.UUID, version *schema.Version) (err error)
	Delete(id uuid.UUID) (err error)
}

func NewVersionRepository(db *database.Database) VersionRepository {
	return &versionRepository{
		DB: db,
	}
}

// implement interface of IVersionRepository
func (_i *versionRepository) GetVersions(req request.VersionsRequest) (versions []*schema.Version, paging paginator.Pagination, err error) {
	var count int64

	query := _i.DB.DB.Model(&schema.Version{})
	query.Count(&count)

	req.Pagination.Count = count
	req.Pagination = paginator.Paging(req.Pagination)

	err = query.Offset(req.Pagination.Offset).Limit(req.Pagination.Limit).Find(&versions).Error
	if err != nil {
		return
	}

	paging = *req.Pagination

	return
}

func (_i *versionRepository) GetOneVersion() (versions *schema.Version, err error) {
	query := _i.DB.DB.Model(&schema.Version{})

	err = query.First(&versions).Error
	if err != nil {
		return
	}

	return
}

func (_i *versionRepository) FindOne(id uuid.UUID) (version *schema.Version, err error) {
	if err := _i.DB.DB.First(&version, id).Error; err != nil {
		return nil, err
	}

	return version, nil
}

func (_i *versionRepository) Create(version *schema.Version) (err error) {
	return _i.DB.DB.Create(version).Error
}

func (_i *versionRepository) Update(id uuid.UUID, version *schema.Version) (err error) {
	return _i.DB.DB.Model(&schema.Version{}).
		Where(&schema.Version{Id: id}).
		Updates(version).Error
}

func (_i *versionRepository) Delete(id uuid.UUID) error {
	return _i.DB.DB.Delete(&schema.Version{}, id).Error
}
