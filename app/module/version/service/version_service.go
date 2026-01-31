package service

import (
	"github.com/google/uuid"
	"go_fiber_wibowo/app/module/version/repository"
	"go_fiber_wibowo/app/module/version/request"
	"go_fiber_wibowo/app/module/version/response"
	"go_fiber_wibowo/utils/paginator"
)

// VersionService
type versionService struct {
	Repo repository.VersionRepository
}

// define interface of IVersionService
//
//go:generate mockgen -destination=version_service_mock.go -package=service . VersionService
type VersionService interface {
	All(req request.VersionsRequest) (versions []*response.Version, paging paginator.Pagination, err error)
	One(req request.VersionsRequest) (versions *response.Version, err error)
	Show(id uuid.UUID) (version *response.Version, err error)
	Store(req request.VersionRequest) (err error)
	Update(id uuid.UUID, req request.VersionRequest) (err error)
	Destroy(id uuid.UUID) error
}

// init VersionService
func NewVersionService(repo repository.VersionRepository) VersionService {
	return &versionService{
		Repo: repo,
	}
}

// implement interface of VersionService
func (_i *versionService) All(req request.VersionsRequest) (versions []*response.Version, paging paginator.Pagination, err error) {
	results, paging, err := _i.Repo.GetVersions(req)
	if err != nil {
		return
	}

	//versions = response.FromDomain(results)
	for _, result := range results {
		versions = append(versions, response.FromDomain(result))
	}

	return
}

// implement interface of VersionService
func (_i *versionService) One(req request.VersionsRequest) (versions *response.Version, err error) {
	results, err := _i.Repo.GetOneVersion()
	if err != nil {
		return
	}

	versions = response.FromDomain(results)

	return
}

func (_i *versionService) Show(id uuid.UUID) (version *response.Version, err error) {
	result, err := _i.Repo.FindOne(id)
	if err != nil {
		return nil, err
	}

	return response.FromDomain(result), nil
}

func (_i *versionService) Store(req request.VersionRequest) (err error) {
	return _i.Repo.Create(req.ToDomain())
}

func (_i *versionService) Update(id uuid.UUID, req request.VersionRequest) (err error) {
	return _i.Repo.Update(id, req.ToDomain())
}

func (_i *versionService) Destroy(id uuid.UUID) error {
	return _i.Repo.Delete(id)
}
