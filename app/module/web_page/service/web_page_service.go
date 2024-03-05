package service

import (
	"dlh_oss_be/app/database/schema"
	"dlh_oss_be/app/module/web_page/repository"
	"dlh_oss_be/app/module/web_page/request"
	"dlh_oss_be/app/module/web_page/response"
	paginationRequest "dlh_oss_be/app/request"
	"dlh_oss_be/utils/paginator"
	"errors"
	"github.com/google/uuid"
	"time"
)

// WebPageService
type webPageService struct {
	repo repository.WebPageRepository
}

// WebPageService define interface of IWebPageService
//
//go:generate mockgen -destination=web_page_service_mock.go -package=service . WebPageService
type WebPageService interface {
	All(req paginationRequest.Pagination) (articles []*response.WebPageResponse, paging paginator.Pagination, err error)
	Show(id string) (article *response.WebPageResponse, err error)
	Store(req request.WebPageRequest) (err error)
	Update(id string, req request.WebPageRequest) (err error)
	Delete(userId, id string) error
	FindByUserId(Id string) (res []*schema.WebPage, err error)
}

// NewWebPageService init WebPageService
func NewWebPageService(webPage repository.WebPageRepository) WebPageService {
	return &webPageService{
		repo: webPage,
	}
}

// All implement interface of WebPageService
func (_i *webPageService) All(req paginationRequest.Pagination) (articles []*response.WebPageResponse, paging paginator.Pagination, err error) {
	results, paging, err := _i.repo.GetAll(req)
	if err != nil {
		return
	}

	for _, result := range results {
		articles = append(articles, response.FromDomain(result))
	}

	return
}

func (_i *webPageService) Show(id string) (article *response.WebPageResponse, err error) {
	result, err := _i.repo.FindOne(id)
	if err != nil {
		return nil, err
	}

	return response.FromDomain(result), nil
}

func (_i *webPageService) Store(req request.WebPageRequest) (err error) {
	return _i.repo.Create(req.ToDomain())
}

func (_i *webPageService) Update(id string, req request.WebPageRequest) (err error) {
	_, err = _i.repo.CheckExist(id)
	if err != nil {
		err = errors.New("web page by id:" + id + ", not found")
		return err
	}

	return _i.repo.Update(id, req.ToDomainUpdate())
}

func (_i *webPageService) Delete(userId, id string) error {
	_, err := _i.repo.CheckExist(id)
	if err != nil {
		err = errors.New("web page by id:" + id + ", not found")
		return err
	}

	req := request.WebPageDeleteRequest{}
	req.DeletedAt = time.Now()
	req.DeletedBy = uuid.Must(uuid.Parse(userId))
	req.IsActive = false

	return _i.repo.Delete(req, id)
}

func (_i *webPageService) FindByUserId(Id string) (res []*schema.WebPage, err error) {
	res, err = _i.repo.FindByUserId(Id)
	if err != nil {
		err = errors.New("web page by user id not found")
		return
	}

	return
}
