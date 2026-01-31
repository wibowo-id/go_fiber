package controller

import (
	"go_fiber_wibowo/app/module/user/request"
	"go_fiber_wibowo/app/module/user/service"
	globalRequest "go_fiber_wibowo/app/request"
	"go_fiber_wibowo/utils/crypto"
	"go_fiber_wibowo/utils/helpers"
	"go_fiber_wibowo/utils/paginator"
	"go_fiber_wibowo/utils/response"
	"go_fiber_wibowo/utils/token"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type userController struct {
	userService service.UserService
}

type UserController interface {
	Index(c *fiber.Ctx) error
	Show(c *fiber.Ctx) error
	Store(c *fiber.Ctx) error
	Update(c *fiber.Ctx) error
	Delete(c *fiber.Ctx) error
}

func NewUserController(userService service.UserService) UserController {
	return &userController{
		userService: userService,
	}
}

// Index get all api/user
// @Summary      Get all api/user
// @Description  API for getting all api/user
// @Tags         Task
// @Security     Bearer
// @Success      200  {object}  response.Response
// @Failure      401  {object}  response.Response
// @Failure      404  {object}  response.Response
// @Failure      422  {object}  response.Response
// @Failure      500  {object}  response.Response
// @Router       /api/user [get]
func (_i *userController) Index(c *fiber.Ctx) error {
	paginate, err := paginator.Paginate(c)
	if err != nil {
		return err
	}

	var req globalRequest.Pagination
	req.Pagination = paginate

	user, paging, err := _i.userService.All(req, c)
	if err != nil {
		return err
	}

	return response.Resp(c, response.Response{
		Messages: response.Messages{"User list successfully retrieved"},
		Data:     user,
		Meta:     paging,
	})
}

// Show get one user
// @Summary      Get one user
// @Description  API for getting one user
// @Tags         Task
// @Security     Bearer
// @Param        id path int true "User Id"
// @Success      200  {object}  response.Response
// @Failure      401  {object}  response.Response
// @Failure      404  {object}  response.Response
// @Failure      422  {object}  response.Response
// @Failure      500  {object}  response.Response
// @Router       /api/user/:id [get]
func (_i *userController) Show(c *fiber.Ctx) error {
	id, err := crypto.DecryptParam(c, "id")
	if err != nil {
		return err
	}
	user, err := _i.userService.Show(id.String())
	if err != nil {
		return err
	}

	return response.Resp(c, response.Response{
		Messages: response.Messages{"User successfully retrieved"},
		Data:     user,
	})
}

// Store create user
// @Summary      Create user
// @Description  API for create user
// @Tags         Task
// @Security     Bearer
// @Body 	     request.UserRequest
// @Success      200  {object}  response.Response
// @Failure      401  {object}  response.Response
// @Failure      404  {object}  response.Response
// @Failure      422  {object}  response.Response
// @Failure      500  {object}  response.Response
// @Router       /api/user [post]
func (_i *userController) Store(c *fiber.Ctx) error {
	req := new(request.UserRequest)
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
	req.Password = helpers.GeneratePassword("password123")
	req.CreatedAt = time.Now()
	req.CreatedBy = uuid.Must(uuid.Parse(user.UserId))

	err = _i.userService.Store(*req)
	if err != nil {
		return err
	}

	return response.Resp(c, response.Response{
		Messages: response.Messages{"User successfully created"},
	})
}

// Update update user
// @Summary      update user
// @Description  API for update user
// @Tags         Task
// @Security     Bearer
// @Body 	     request.UserRequest
// @Param        id path int true "User Id"
// @Success      200  {object}  response.Response
// @Failure      401  {object}  response.Response
// @Failure      404  {object}  response.Response
// @Failure      422  {object}  response.Response
// @Failure      500  {object}  response.Response
// @Router       /api/user/:id [put]
func (_i *userController) Update(c *fiber.Ctx) error {
	id, err := crypto.DecryptParam(c, "id")
	if err != nil {
		return err
	}
	req := new(request.UserRequest)
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
	err = _i.userService.Update(id.String(), *req)
	if err != nil {
		return err
	}

	return response.Resp(c, response.Response{
		Messages: response.Messages{"User successfully updated"},
	})
}

// Delete delete user
// @Summary      delete user
// @Description  API for delete user
// @Tags         Task
// @Security     Bearer
// @Param        id path int true "User Id"
// @Success      200  {object}  response.Response
// @Failure      401  {object}  response.Response
// @Failure      404  {object}  response.Response
// @Failure      422  {object}  response.Response
// @Failure      500  {object}  response.Response
// @Router       /api/user/:id [delete]
func (_i *userController) Delete(c *fiber.Ctx) error {
	GetToken := token.ExtractToken(c)
	user, err := token.ExtractTokenID(GetToken)
	if err != nil {
		return err
	}
	id, err := crypto.DecryptParam(c, "id")
	if err != nil {
		return err
	}
	err = _i.userService.Delete(user.UserId, id.String())
	if err != nil {
		return err
	}

	return response.Resp(c, response.Response{
		Messages: response.Messages{"User successfully deleted"},
	})
}
