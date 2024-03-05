package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"go_fiber_wibowo/app/module/web_page/request"
	"go_fiber_wibowo/app/module/web_page/service"
	paginationRequest "go_fiber_wibowo/app/request"
	"go_fiber_wibowo/utils/paginator"
	"go_fiber_wibowo/utils/response"
	"go_fiber_wibowo/utils/token"
	"time"
)

type webPageController struct {
	webPageService service.WebPageService
}

type WebPageController interface {
	Index(c *fiber.Ctx) error
	Show(c *fiber.Ctx) error
	Store(c *fiber.Ctx) error
	Update(c *fiber.Ctx) error
	Delete(c *fiber.Ctx) error
}

func NewWebPageController(webPageService service.WebPageService) WebPageController {
	return &webPageController{
		webPageService: webPageService,
	}
}

// Index get all webPages
// @Summary      Get all webPages
// @Description  API for getting all webPages
// @Tags         Task
// @Security     Bearer
// @Success      200  {object}  response.Response
// @Failure      401  {object}  response.Response
// @Failure      404  {object}  response.Response
// @Failure      422  {object}  response.Response
// @Failure      500  {object}  response.Response
// @Router       /webPages [get]
func (_i *webPageController) Index(c *fiber.Ctx) error {
	paginate, err := paginator.Paginate(c)
	if err != nil {
		return err
	}

	var req paginationRequest.Pagination
	req.Pagination = paginate

	webPages, paging, err := _i.webPageService.All(req)
	if err != nil {
		return err
	}

	return response.Resp(c, response.Response{
		Messages: response.Messages{"Web Page list successfully retrieved"},
		Data:     webPages,
		Meta:     paging,
	})
}

// Show get one webPage
// @Summary      Get one webPage
// @Description  API for getting one webPage
// @Tags         Task
// @Security     Bearer
// @Param        id path int true "Web Page ID"
// @Success      200  {object}  response.Response
// @Failure      401  {object}  response.Response
// @Failure      404  {object}  response.Response
// @Failure      422  {object}  response.Response
// @Failure      500  {object}  response.Response
// @Router       /webPages/:id [get]
func (_i *webPageController) Show(c *fiber.Ctx) error {
	id := c.Params("id")

	webPages, err := _i.webPageService.Show(id)
	if err != nil {
		return err
	}

	return response.Resp(c, response.Response{
		Messages: response.Messages{"Web Page successfully retrieved"},
		Data:     webPages,
	})
}

// Store create webPage
// @Summary      Create webPage
// @Description  API for create webPage
// @Tags         Task
// @Security     Bearer
// @Body 	     request.WebPageRequest
// @Success      200  {object}  response.Response
// @Failure      401  {object}  response.Response
// @Failure      404  {object}  response.Response
// @Failure      422  {object}  response.Response
// @Failure      500  {object}  response.Response
// @Router       /webPages [post]
func (_i *webPageController) Store(c *fiber.Ctx) error {
	req := new(request.WebPageRequest)
	if err := response.ParseAndValidate(c, req); err != nil {
		return err
	}

	GetToken := token.ExtractToken(c)
	user, err := token.ExtractTokenID(GetToken)
	if err != nil {
		return err
	}

	req.Id = uuid.New()
	req.CreatedAt = time.Now()
	req.CreatedBy = uuid.Must(uuid.Parse(user.UserId))

	err = _i.webPageService.Store(*req)
	if err != nil {
		return err
	}

	return response.Resp(c, response.Response{
		Messages: response.Messages{"Web Page successfully created"},
	})
}

// Update webPage
// @Summary      update webPage
// @Description  API for update webPage
// @Tags         Task
// @Security     Bearer
// @Body 	     request.WebPageRequest
// @Param        id path int true "Web Page ID"
// @Success      200  {object}  response.Response
// @Failure      401  {object}  response.Response
// @Failure      404  {object}  response.Response
// @Failure      422  {object}  response.Response
// @Failure      500  {object}  response.Response
// @Router       /webPages/:id [put]
func (_i *webPageController) Update(c *fiber.Ctx) error {
	id := c.Params("id")

	req := new(request.WebPageRequest)
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

	err = _i.webPageService.Update(id, *req)
	if err != nil {
		return err
	}

	return response.Resp(c, response.Response{
		Messages: response.Messages{"Web Page successfully updated"},
	})
}

// Delete webPage
// @Summary      delete webPage
// @Description  API for delete webPage
// @Tags         Task
// @Security     Bearer
// @Param        id path int true "Web Page ID"
// @Success      200  {object}  response.Response
// @Failure      401  {object}  response.Response
// @Failure      404  {object}  response.Response
// @Failure      422  {object}  response.Response
// @Failure      500  {object}  response.Response
// @Router       /webPages/:id [delete]
func (_i *webPageController) Delete(c *fiber.Ctx) error {
	GetToken := token.ExtractToken(c)
	user, err := token.ExtractTokenID(GetToken)
	if err != nil {
		return err
	}

	id := c.Params("id")

	err = _i.webPageService.Delete(user.UserId, id)
	if err != nil {
		return err
	}

	return response.Resp(c, response.Response{
		Messages: response.Messages{"Web Page successfully deleted"},
	})
}
