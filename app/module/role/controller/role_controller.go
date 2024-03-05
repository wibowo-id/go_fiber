package controller

import (
	"dlh_oss_be/app/module/role/request"
	"dlh_oss_be/app/module/role/service"
	globalRequest "dlh_oss_be/app/request"
	"dlh_oss_be/utils/paginator"
	"dlh_oss_be/utils/response"
	"dlh_oss_be/utils/token"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"time"
)

type roleController struct {
	roleService service.RoleService
}

type RoleController interface {
	Index(c *fiber.Ctx) error
	Store(c *fiber.Ctx) error
	Show(c *fiber.Ctx) error
	Update(c *fiber.Ctx) error
	Delete(c *fiber.Ctx) error
}

func NewRoleController(roleService service.RoleService) RoleController {
	return &roleController{
		roleService: roleService,
	}
}

func (_i *roleController) Index(c *fiber.Ctx) error {
	paginate, err := paginator.Paginate(c)
	if err != nil {
		return err
	}

	var req globalRequest.Pagination
	req.Pagination = paginate

	beritaKegiatan, paging, err := _i.roleService.All(req)
	if err != nil {
		return err
	}

	return response.Resp(c, response.Response{
		Messages: response.Messages{"Role list successfully retrieved"},
		Data:     beritaKegiatan,
		Meta:     paging,
	})
}

func (_i *roleController) Store(c *fiber.Ctx) error {
	req := new(request.RoleRequest)
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

	err = _i.roleService.Store(*req)
	if err != nil {
		return err
	}

	return response.Resp(c, response.Response{
		Messages: response.Messages{"Role successfully created"},
	})
}

func (_i *roleController) Show(c *fiber.Ctx) error {
	id := c.Params("id")

	beritaKegiatan, err := _i.roleService.Show(id)
	if err != nil {
		return err
	}

	return response.Resp(c, response.Response{
		Messages: response.Messages{"Role successfully retrieved"},
		Data:     beritaKegiatan,
	})
}

func (_i *roleController) Update(c *fiber.Ctx) error {
	id := c.Params("id")

	req := new(request.RoleRequest)
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

	err = _i.roleService.Update(id, *req)
	if err != nil {
		return err
	}

	return response.Resp(c, response.Response{
		Messages: response.Messages{"Role successfully updated"},
	})
}

func (_i *roleController) Delete(c *fiber.Ctx) error {
	GetToken := token.ExtractToken(c)
	user, err := token.ExtractTokenID(GetToken)
	if err != nil {
		return err
	}

	id := c.Params("id")

	err = _i.roleService.Delete(user.UserId, id)
	if err != nil {
		return err
	}

	return response.Resp(c, response.Response{
		Messages: response.Messages{"Role successfully deleted"},
	})
}
