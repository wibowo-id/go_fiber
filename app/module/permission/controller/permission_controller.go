package controller

import (
	"go_fiber_wibowo/app/module/permission/request"
	"go_fiber_wibowo/app/module/permission/service"
	globalRequest "go_fiber_wibowo/app/request"
	"go_fiber_wibowo/utils/crypto"
	"go_fiber_wibowo/utils/paginator"
	"go_fiber_wibowo/utils/response"
	"go_fiber_wibowo/utils/token"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
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

	req.Id = crypto.EncryptedUUID{UUID: uuid.New()}
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
	id, err := crypto.DecryptParam(c, "id")
	if err != nil {
		return err
	}
	beritaKegiatan, err := _i.permissionService.Show(id.String())
	if err != nil {
		return err
	}

	return response.Resp(c, response.Response{
		Messages: response.Messages{"Permission successfully retrieved"},
		Data:     beritaKegiatan,
	})
}

func (_i *permissionController) Update(c *fiber.Ctx) error {
	id, err := crypto.DecryptParam(c, "id")
	if err != nil {
		return err
	}
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
	err = _i.permissionService.Update(id.String(), *req)
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
	id, err := crypto.DecryptParam(c, "id")
	if err != nil {
		return err
	}
	err = _i.permissionService.Delete(user.UserId, id.String())
	if err != nil {
		return err
	}

	return response.Resp(c, response.Response{
		Messages: response.Messages{"Permission successfully deleted"},
	})
}
