package controller

import (
	"dlh_oss_be/app/module/permission/request"
	"dlh_oss_be/app/module/permission/service"
	globalRequest "dlh_oss_be/app/request"
	"dlh_oss_be/utils/paginator"
	"dlh_oss_be/utils/response"
	"dlh_oss_be/utils/token"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"time"
)

type permissionController struct {
	permissionService service.PermissionService
}

type PermissionController interface {
	Index(c *fiber.Ctx) error
	Store(c *fiber.Ctx) error
	Show(c *fiber.Ctx) error
	Update(c *fiber.Ctx) error
	Delete(c *fiber.Ctx) error
}

func NewPermissionController(permissionService service.PermissionService) PermissionController {
	return &permissionController{
		permissionService: permissionService,
	}
}

func (_i *permissionController) Index(c *fiber.Ctx) error {
	paginate, err := paginator.Paginate(c)
	if err != nil {
		return err
	}

	var req globalRequest.Pagination
	req.Pagination = paginate

	beritaKegiatan, paging, err := _i.permissionService.All(req)
	if err != nil {
		return err
	}

	return response.Resp(c, response.Response{
		Messages: response.Messages{"Permission list successfully retrieved"},
		Data:     beritaKegiatan,
		Meta:     paging,
	})
}

func (_i *permissionController) Store(c *fiber.Ctx) error {
	req := new(request.PermissionRequest)
	if err := response.ParseAndValidate(c, req); err != nil {
		return err
	}

	GetToken := token.ExtractToken(c)
	user, err := token.ExtractTokenID(GetToken)
	if err != nil {
		return err
	}

	req.Id = uuid.New()
	req.IsActive = true
	req.CreatedAt = time.Now()
	req.CreatedBy = uuid.Must(uuid.Parse(user.UserId))

	err = _i.permissionService.Store(*req)
	if err != nil {
		return err
	}

	return response.Resp(c, response.Response{
		Messages: response.Messages{"Permission successfully created"},
	})
}

func (_i *permissionController) Show(c *fiber.Ctx) error {
	id := c.Params("id")

	beritaKegiatan, err := _i.permissionService.Show(id)
	if err != nil {
		return err
	}

	return response.Resp(c, response.Response{
		Messages: response.Messages{"Permission successfully retrieved"},
		Data:     beritaKegiatan,
	})
}

func (_i *permissionController) Update(c *fiber.Ctx) error {
	id := c.Params("id")

	req := new(request.PermissionRequest)
	if err := response.ParseAndValidate(c, req); err != nil {
		return err
	}

	GetToken := token.ExtractToken(c)
	user, err := token.ExtractTokenID(GetToken)
	if err != nil {
		return err
	}

	req.UpdatedAt = time.Now()
	req.UpdatedBy = uuid.Must(uuid.Parse(user.UserId))

	err = _i.permissionService.Update(id, *req)
	if err != nil {
		return err
	}

	return response.Resp(c, response.Response{
		Messages: response.Messages{"Permission successfully updated"},
	})
}

func (_i *permissionController) Delete(c *fiber.Ctx) error {
	GetToken := token.ExtractToken(c)
	user, err := token.ExtractTokenID(GetToken)
	if err != nil {
		return err
	}

	id := c.Params("id")

	err = _i.permissionService.Delete(user.UserId, id)
	if err != nil {
		return err
	}

	return response.Resp(c, response.Response{
		Messages: response.Messages{"Permission successfully deleted"},
	})
}
