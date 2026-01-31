package controller

import (
	"go_fiber_wibowo/app/module/version/request"
	"go_fiber_wibowo/app/module/version/service"
	"go_fiber_wibowo/utils/crypto"
	"go_fiber_wibowo/utils/paginator"
	"go_fiber_wibowo/utils/response"

	"github.com/gofiber/fiber/v2"
)

type versionController struct {
	versionService service.VersionService
}

type VersionController interface {
	Index(c *fiber.Ctx) error
	IndexOne(c *fiber.Ctx) error
	Show(c *fiber.Ctx) error
	Store(c *fiber.Ctx) error
	Update(c *fiber.Ctx) error
	Delete(c *fiber.Ctx) error
}

func NewVersionController(versionService service.VersionService) VersionController {
	return &versionController{
		versionService: versionService,
	}
}

// get all versions
// @Summary      Get all versions
// @Description  API for getting all versions
// @Tags         Task
// @Security     Bearer
// @Success      200  {object}  response.Response
// @Failure      401  {object}  response.Response
// @Failure      404  {object}  response.Response
// @Failure      422  {object}  response.Response
// @Failure      500  {object}  response.Response
// @Router       /versions [get]
func (_i *versionController) Index(c *fiber.Ctx) error {
	paginate, err := paginator.Paginate(c)
	if err != nil {
		return err
	}

	var req request.VersionsRequest
	req.Pagination = paginate

	articles, paging, err := _i.versionService.All(req)
	if err != nil {
		return err
	}

	return response.Resp(c, response.Response{
		Messages: response.Messages{"Article list successfully retrieved"},
		Data:     articles,
		Meta:     paging,
	})
}

// get one versions
// @Summary      Get one versions
// @Description  API for getting one versions
// @Tags         Task
// @Security     Bearer
// @Success      200  {object}  response.Response
// @Failure      401  {object}  response.Response
// @Failure      404  {object}  response.Response
// @Failure      422  {object}  response.Response
// @Failure      500  {object}  response.Response
// @Router       /version [get]
func (_i *versionController) IndexOne(c *fiber.Ctx) error {
	var req request.VersionsRequest

	versions, err := _i.versionService.One(req)
	if err != nil {
		return err
	}

	err = c.Status(fiber.StatusOK).JSON(versions)
	if err != nil {
		return err
	}
	return err
}

// get one version
// @Summary      Get one version
// @Description  API for getting one version
// @Tags         Task
// @Security     Bearer
// @Param        id path int true "Version ID"
// @Success      200  {object}  response.Response
// @Failure      401  {object}  response.Response
// @Failure      404  {object}  response.Response
// @Failure      422  {object}  response.Response
// @Failure      500  {object}  response.Response
// @Router       /versions/:id [get]
func (_i *versionController) Show(c *fiber.Ctx) error {
	id, err := crypto.DecryptParam(c, "id")
	if err != nil {
		return err
	}
	versions, err := _i.versionService.Show(id)
	if err != nil {
		return err
	}

	return response.Resp(c, response.Response{
		Messages: response.Messages{"Version successfully retrieved"},
		Data:     versions,
	})
}

// create version
// @Summary      Create version
// @Description  API for create version
// @Tags         Task
// @Security     Bearer
// @Body 	     request.VersionRequest
// @Success      200  {object}  response.Response
// @Failure      401  {object}  response.Response
// @Failure      404  {object}  response.Response
// @Failure      422  {object}  response.Response
// @Failure      500  {object}  response.Response
// @Router       /versions [post]
func (_i *versionController) Store(c *fiber.Ctx) error {
	req := new(request.VersionRequest)
	if err := response.ParseAndValidate(c, req); err != nil {
		return err
	}

	err := _i.versionService.Store(*req)
	if err != nil {
		return err
	}

	return response.Resp(c, response.Response{
		Messages: response.Messages{"Version successfully created"},
	})
}

// update version
// @Summary      update version
// @Description  API for update version
// @Tags         Task
// @Security     Bearer
// @Body 	     request.VersionRequest
// @Param        id path int true "Version ID"
// @Success      200  {object}  response.Response
// @Failure      401  {object}  response.Response
// @Failure      404  {object}  response.Response
// @Failure      422  {object}  response.Response
// @Failure      500  {object}  response.Response
// @Router       /versions/:id [put]
func (_i *versionController) Update(c *fiber.Ctx) error {
	id, err := crypto.DecryptParam(c, "id")
	if err != nil {
		return err
	}
	req := new(request.VersionRequest)
	if err := response.ParseAndValidate(c, req); err != nil {
		return err
	}
	err = _i.versionService.Update(id, *req)
	if err != nil {
		return err
	}

	return response.Resp(c, response.Response{
		Messages: response.Messages{"Version successfully updated"},
	})
}

// delete version
// @Summary      delete version
// @Description  API for delete version
// @Tags         Task
// @Security     Bearer
// @Param        id path int true "Version ID"
// @Success      200  {object}  response.Response
// @Failure      401  {object}  response.Response
// @Failure      404  {object}  response.Response
// @Failure      422  {object}  response.Response
// @Failure      500  {object}  response.Response
// @Router       /versions/:id [delete]
func (_i *versionController) Delete(c *fiber.Ctx) error {
	id, err := crypto.DecryptParam(c, "id")
	if err != nil {
		return err
	}
	err = _i.versionService.Destroy(id)
	if err != nil {
		return err
	}

	return response.Resp(c, response.Response{
		Messages: response.Messages{"Version successfully deleted"},
	})
}
