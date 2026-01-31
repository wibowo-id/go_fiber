package controller

import (
	"go_fiber_wibowo/app/module/role/request"
	"go_fiber_wibowo/app/module/role/service"
	globalRequest "go_fiber_wibowo/app/request"
	"go_fiber_wibowo/utils/crypto"
	"go_fiber_wibowo/utils/paginator"
	"go_fiber_wibowo/utils/response"
	"go_fiber_wibowo/utils/token"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
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

	req.Id = crypto.EncryptedUUID{UUID: uuid.New()}
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
	id, err := crypto.DecryptParam(c, "id")
	if err != nil {
		return err
	}
	beritaKegiatan, err := _i.roleService.Show(id.String())
	if err != nil {
		return err
	}

	return response.Resp(c, response.Response{
		Messages: response.Messages{"Role successfully retrieved"},
		Data:     beritaKegiatan,
	})
}

func (_i *roleController) Update(c *fiber.Ctx) error {
	id, err := crypto.DecryptParam(c, "id")
	if err != nil {
		return err
	}
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
	err = _i.roleService.Update(id.String(), *req)
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
	id, err := crypto.DecryptParam(c, "id")
	if err != nil {
		return err
	}
	err = _i.roleService.Delete(user.UserId, id.String())
	if err != nil {
		return err
	}

	return response.Resp(c, response.Response{
		Messages: response.Messages{"Role successfully deleted"},
	})
}
