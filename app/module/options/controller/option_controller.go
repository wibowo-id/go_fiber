package controller

import (
	"github.com/gofiber/fiber/v2"
	auth_service "go_fiber_wibowo/app/module/auth/service"
	"go_fiber_wibowo/app/module/options/service"
	"go_fiber_wibowo/utils/response"
)

type optionController struct {
	optionService service.OptionService
	authService   auth_service.AuthService
}

type OptionController interface {
	GetWilayah(c *fiber.Ctx) error
	GetRole(c *fiber.Ctx) error
	GetPermission(c *fiber.Ctx) error
}

func NewOptionsController(optionService service.OptionService, authService auth_service.AuthService) OptionController {
	return &optionController{
		optionService: optionService,
		authService:   authService,
	}
}

// GetWilayah get all api/options/wilayah
// @Summary      Get all api/options/wilayah
// @Description  API for getting all api/options/wilayah
// @Tags         Task
// @Security     Bearer
// @Success      200  {object}  response.Response
// @Failure      401  {object}  response.Response
// @Failure      404  {object}  response.Response
// @Failure      422  {object}  response.Response
// @Failure      500  {object}  response.Response
// @Router       /api/options/wilayah [get]
func (_i *optionController) GetWilayah(c *fiber.Ctx) error {
	kode := c.Query("kode")
	area := ""
	if len(kode) == 0 {
		area = "Provinsi"
	} else if len(kode) == 2 {
		area = "Kota / Kabupaten"
	} else if len(kode) == 5 {
		area = "Kecamatan"
	} else if len(kode) == 8 {
		area = "Kelurahan"
	}

	wilayah, err := _i.optionService.GetWilayah(kode)
	if err != nil {
		return err
	}

	return response.Resp(c, response.Response{
		Messages: response.Messages{"Wilayah " + area + " list successfully retrieved"},
		Data:     wilayah,
	})
}

// GetRole get all api/options/wilayah
// @Summary      Get all api/options/wilayah
// @Description  API for getting all api/options/wilayah
// @Tags         Task
// @Security     Bearer
// @Success      200  {object}  response.Response
// @Failure      401  {object}  response.Response
// @Failure      404  {object}  response.Response
// @Failure      422  {object}  response.Response
// @Failure      500  {object}  response.Response
// @Router       /api/options/wilayah [get]
func (_i *optionController) GetRole(c *fiber.Ctx) error {
	data, err := _i.optionService.GetRole()
	if err != nil {
		return err
	}

	return response.Resp(c, response.Response{
		Messages: response.Messages{"Role list successfully retrieved"},
		Data:     data,
	})
}

// GetPermission get all api/options/wilayah
// @Summary      Get all api/options/wilayah
// @Description  API for getting all api/options/wilayah
// @Tags         Task
// @Security     Bearer
// @Success      200  {object}  response.Response
// @Failure      401  {object}  response.Response
// @Failure      404  {object}  response.Response
// @Failure      422  {object}  response.Response
// @Failure      500  {object}  response.Response
// @Router       /api/options/wilayah [get]
func (_i *optionController) GetPermission(c *fiber.Ctx) error {
	data, err := _i.optionService.GetPermission()
	if err != nil {
		return err
	}

	return response.Resp(c, response.Response{
		Messages: response.Messages{"Permission list successfully retrieved"},
		Data:     data,
	})
}
