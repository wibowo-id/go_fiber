package repository

import (
	"github.com/google/uuid"
	"go_fiber_wibowo/app/database/schema"
	"go_fiber_wibowo/app/module/web_page/request"
	paginationRequest "go_fiber_wibowo/app/request"
	"go_fiber_wibowo/internal/bootstrap/database"
	"go_fiber_wibowo/utils/config"
	"go_fiber_wibowo/utils/paginator"
)

type webPageRepository struct {
	DB  *database.Database
	Cfg *config.Config
}

//go:generate mockgen -destination=web_page_repository_mock.go -package=repository . WebPageRepository
type WebPageRepository interface {
	GetAll(req paginationRequest.Pagination) (webPages []*schema.WebPage, paging paginator.Pagination, err error)
	FindOne(id string) (webPage *schema.WebPage, err error)
	CheckExist(id string) (res *schema.WebPage, err error)
	Create(webPage *schema.WebPage) (err error)
	Update(id string, webPage *schema.WebPage) (err error)
	Delete(req request.WebPageDeleteRequest, id string) (err error)
	FindByUserId(Id string) (user []*schema.WebPage, err error)
}

func NewWebPageRepository(db *database.Database) WebPageRepository {
	return &webPageRepository{
		DB: db,
	}
}

// implement interface of IArticleRepository
func (_i *webPageRepository) GetAll(req paginationRequest.Pagination) (webPages []*schema.WebPage, paging paginator.Pagination, err error) {
	var count int64

	query := _i.DB.DB.Model(&schema.WebPage{})
	query.Count(&count)

	req.Pagination.Count = count
	req.Pagination = paginator.Paging(req.Pagination)

	err = query.Offset(req.Pagination.Offset).Limit(req.Pagination.Limit).Find(&webPages).Error
	if err != nil {
		return
	}

	paging = *req.Pagination

	return
}

func (_i *webPageRepository) FindOne(id string) (webPage *schema.WebPage, err error) {
	if err := _i.DB.DB.First(&webPage, uuid.Must(uuid.Parse(id))).Error; err != nil {
		return nil, err
	}

	return webPage, nil
}

func (_i *webPageRepository) CheckExist(id string) (res *schema.WebPage, err error) {
	qState := "SELECT id FROM web_pages "
	qState = qState + "WHERE id = '" + id + "' "
	if err := _i.DB.DB.Raw(qState).Error; err != nil {
		return nil, err
	}

	return res, nil
}

func (_i *webPageRepository) Create(webPage *schema.WebPage) (err error) {
	return _i.DB.DB.Create(webPage).Error
}

func (_i *webPageRepository) Update(id string, webPage *schema.WebPage) (err error) {
	err = _i.DB.DB.Model(&schema.WebPage{}).
		Where(&schema.WebPage{Id: uuid.Must(uuid.Parse(id))}).
		Updates(webPage).Error
	if err != nil {
		return err
	}

	return nil
}

func (_i *webPageRepository) Delete(req request.WebPageDeleteRequest, id string) error {
	if err := _i.DB.DB.Table("web_pages").Where("id", id).Updates(map[string]interface{}{"is_active": req.IsActive, "deleted_at": req.DeletedAt, "deleted_by": req.DeletedBy}).Error; err != nil {
		return err
	}

	return nil
}

func (_i *webPageRepository) FindByUserId(Id string) (res []*schema.WebPage, err error) {
	qState := "SELECT * FROM web_pages "
	qState = qState + "WHERE created_by = '" + Id + "' "

	err = _i.DB.DB.Raw(qState).Scan(&res).Error
	if err != nil {
		return nil, err
	}

	return res, nil
}
