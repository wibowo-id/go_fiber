package controller

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"go_fiber_wibowo/app/database/schema"
	"go_fiber_wibowo/app/module/auth/request"
	"go_fiber_wibowo/app/module/auth/service"
	role_permission_service "go_fiber_wibowo/app/module/role_permission/service"
	user_repo "go_fiber_wibowo/app/module/user/repository"
	wa_request "go_fiber_wibowo/app/module/wa_sender/request"
	wa_service "go_fiber_wibowo/app/module/wa_sender/service"
	"go_fiber_wibowo/utils/helpers"
	"go_fiber_wibowo/utils/response"
	"go_fiber_wibowo/utils/token"
)

type authController struct {
	authService           service.AuthService
	waSenderService       wa_service.WaSenderService
	rolePermissionService role_permission_service.RolePermissionService
	userRepo              user_repo.UserRepository
}

type AuthController interface {
	LoginAdmin(c *fiber.Ctx) error
	Login(c *fiber.Ctx) error
	Register(c *fiber.Ctx) error
	ForgotPassword(c *fiber.Ctx) error
	ChangePassword(c *fiber.Ctx) error
	ChangePasswordPublic(c *fiber.Ctx) error
	RequestOTP(c *fiber.Ctx) error
	VerifyOTP(c *fiber.Ctx) error
	Profile(c *fiber.Ctx) error
	Logout(c *fiber.Ctx) error
	UpdateProfile(c *fiber.Ctx) error
	UpdateFirebaseToken(c *fiber.Ctx) error
	Menu(c *fiber.Ctx) error
	Dashboard(c *fiber.Ctx) error
}

func NewAuthController(authService service.AuthService, rolePermissionService role_permission_service.RolePermissionService, waSenderService wa_service.WaSenderService, userRepo user_repo.UserRepository) AuthController {
	return &authController{
		authService:           authService,
		rolePermissionService: rolePermissionService,
		waSenderService:       waSenderService,
		userRepo:              userRepo,
	}
}

// LoginAdmin
// @Summary      Do login admin
// @Description  API for do login admin
// @Tags         Authentication
// @Security     Bearer
// @Body 	     request.LoginRequest
// @Success      200  {object}  response.Response
// @Failure      401  {object}  response.Response
// @Failure      404  {object}  response.Response
// @Failure      422  {object}  response.Response
// @Failure      500  {object}  response.Response
// @Router       /api/auth/admin/login [post]
func (_i *authController) LoginAdmin(c *fiber.Ctx) error {
	req := new(request.LoginRequest)
	if err := response.ParseAndValidate(c, req); err != nil {
		return err
	}

	res, err := _i.authService.LoginAdmin(*req)
	if err != nil {
		return err
	}

	return response.Resp(c, response.Response{
		Data:     res,
		Messages: response.Messages{"Login success"},
		Code:     fiber.StatusOK,
	})
}

// Login
// @Summary      Do login
// @Description  API for do login
// @Tags         Authentication
// @Security     Bearer
// @Body 	     request.LoginRequest
// @Success      200  {object}  response.Response
// @Failure      401  {object}  response.Response
// @Failure      404  {object}  response.Response
// @Failure      422  {object}  response.Response
// @Failure      500  {object}  response.Response
// @Router       /api/auth/login [post]
func (_i *authController) Login(c *fiber.Ctx) error {
	req := new(request.LoginRequest)
	if err := response.ParseAndValidate(c, req); err != nil {
		return err
	}

	res, err := _i.authService.Login(*req)
	if err != nil {
		return err
	}

	return response.Resp(c, response.Response{
		Data:     res,
		Messages: response.Messages{"Login success"},
		Code:     fiber.StatusOK,
	})
}

// Register
// @Summary      Register
// @Description  API for register
// @Tags         Authentication
// @Body 	     request.RegisterRequest
// @Success      200  {object}  response.Response
// @Failure      401  {object}  response.Response
// @Failure      404  {object}  response.Response
// @Failure      422  {object}  response.Response
// @Failure      500  {object}  response.Response
// @Router       /api/v1/register [post]
func (_i *authController) Register(c *fiber.Ctx) error {
	req := new(request.RegisterRequest)
	if err := response.ParseAndValidate(c, req); err != nil {
		return err
	}

	res, err := _i.authService.Register(*req)
	if err != nil {
		return err
	}

	otpParam := &wa_request.OtpRequest{}
	otpParam.Phone = req.Telepon
	otpParam.Message = `Registrasi berhasil...\n\n`

	_, err = _i.waSenderService.SendOtp(*otpParam)
	if err != nil {
		return err
	}

	return response.Resp(c, response.Response{
		Data:     res,
		Messages: response.Messages{"Register success"},
		Code:     fiber.StatusOK,
	})
}

func (_i *authController) ForgotPassword(c *fiber.Ctx) error {
	req := new(request.ForgotPasswordRequest)
	if err := response.ParseAndValidate(c, req); err != nil {
		return err
	}

	phone, err := _i.authService.ForgotPassword(*req)
	if err != nil {
		return err
	}

	otpParam := &wa_request.OtpRequest{}
	otpParam.Phone = phone
	otpParam.Message = `Forgot password...\n\n`

	_, err = _i.waSenderService.SendOtp(*otpParam)
	if err != nil {
		return err
	}

	return response.Resp(c, response.Response{
		Messages: response.Messages{"Request forgot password success, please check OTP number on your whatsapp"},
		Code:     fiber.StatusOK,
	})
}

func (_i *authController) ChangePassword(c *fiber.Ctx) error {
	GetToken := token.ExtractToken(c)
	user, err := token.ExtractTokenID(GetToken)
	if err != nil {
		return err
	}

	userToken := schema.User{}
	userToken.RememberToken = user.Token

	_, err = _i.userRepo.FindUserByToken(userToken)
	if err != nil {
		err = errors.New("user with token not found, please re-login ")
		return err
	}

	req := new(request.ChangePasswordRequest)
	if err := response.ParseAndValidate(c, req); err != nil {
		return err
	}

	_, err = _i.authService.ChangePassword(*req, user.UserId)
	if err != nil {
		return err
	}

	return response.Resp(c, response.Response{
		Messages: response.Messages{"Change password success"},
		Code:     fiber.StatusOK,
	})
}

func (_i *authController) ChangePasswordPublic(c *fiber.Ctx) error {
	req := new(request.ChangePasswordPublicRequest)
	if err := response.ParseAndValidate(c, req); err != nil {
		return err
	}

	req.Telepon = helpers.VerifyFormatPhone(req.Telepon)
	user, err := _i.userRepo.FindUserByPhone(req.Telepon)
	if err != nil {
		err = errors.New("user by phone not found")
		return err
	}

	newReq := request.ChangePasswordRequest{}
	newReq.Password = req.Password
	newReq.ConfirmPassword = req.ConfirmPassword

	_, err = _i.authService.ChangePassword(newReq, user.Id.String())
	if err != nil {
		return err
	}

	return response.Resp(c, response.Response{
		Messages: response.Messages{"Change password success"},
		Code:     fiber.StatusOK,
	})
}

func (_i *authController) RequestOTP(c *fiber.Ctx) error {
	req := new(request.OtpRequest)
	if err := response.ParseAndValidate(c, req); err != nil {
		return err
	}

	otpParam := &wa_request.OtpRequest{}
	otpParam.Phone = req.Telepon
	otpParam.Message = `Request OTP...\n\n`

	_, err := _i.waSenderService.SendOtp(*otpParam)
	if err != nil {
		return err
	}

	return response.Resp(c, response.Response{
		Messages: response.Messages{"OTP has been sent to whatsapp"},
		Code:     fiber.StatusOK,
	})
}

func (_i *authController) VerifyOTP(c *fiber.Ctx) error {
	req := new(request.OtpConfirmRequest)
	if err := response.ParseAndValidate(c, req); err != nil {
		return err
	}

	_, err := _i.authService.VerifyOTP(*req)
	if err != nil {
		return err
	}

	return response.Resp(c, response.Response{
		Messages: response.Messages{"OTP has been verified"},
		Code:     fiber.StatusOK,
	})
}

func (_i *authController) Profile(c *fiber.Ctx) error {
	GetToken := token.ExtractToken(c)
	user, err := token.ExtractTokenID(GetToken)
	if err != nil {
		return err
	}

	res, err := _i.authService.Profile(user.UserId)
	if err != nil {
		return err
	}

	return response.Resp(c, response.Response{
		Data:     res,
		Messages: response.Messages{"Success"},
		Code:     fiber.StatusOK,
	})
}

func (_i *authController) Logout(c *fiber.Ctx) error {
	GetToken := token.ExtractToken(c)
	user, err := token.ExtractTokenID(GetToken)
	if err != nil {
		return err
	}

	_, err = _i.authService.Logout(user.UserId)
	if err != nil {
		return err
	}

	return response.Resp(c, response.Response{
		Messages: response.Messages{"Success"},
		Code:     fiber.StatusOK,
	})
}

func (_i *authController) UpdateProfile(c *fiber.Ctx) error {
	req := new(request.UpdateProfileRequest)
	if err := response.ParseAndValidate(c, req); err != nil {
		return err
	}

	GetToken := token.ExtractToken(c)
	user, err := token.ExtractTokenID(GetToken)
	if err != nil {
		return err
	}

	Id := uuid.Must(uuid.Parse(user.UserId))
	req.Telepon = helpers.VerifyFormatPhone(req.Telepon)

	res, err := _i.authService.UpdateProfile(req, Id)
	if err != nil {
		return err
	}

	return response.Resp(c, response.Response{
		Data:     res,
		Messages: response.Messages{"Update profile success"},
		Code:     fiber.StatusOK,
	})
}

func (_i *authController) UpdateFirebaseToken(c *fiber.Ctx) error {
	req := new(request.RegisterRequest)
	if err := response.ParseAndValidate(c, req); err != nil {
		return err
	}

	res, err := _i.authService.Register(*req)
	if err != nil {
		return err
	}

	return response.Resp(c, response.Response{
		Data:     res,
		Messages: response.Messages{"Update firebase token success"},
		Code:     fiber.StatusOK,
	})
}

func (_i *authController) Menu(c *fiber.Ctx) error {
	GetToken := token.ExtractToken(c)
	getUser, err := _i.authService.GetUserByToken(GetToken)
	if err != nil {
		return err
	}

	res, err := _i.rolePermissionService.GetMenu(getUser.RoleId)
	if err != nil {
		return err
	}

	return response.Resp(c, response.Response{
		Data:     res,
		Messages: response.Messages{"Success"},
		Code:     fiber.StatusOK,
	})
}

func (_i *authController) Dashboard(c *fiber.Ctx) error {
	res, err := _i.userRepo.GetDashboard()
	if err != nil {
		return err
	}

	return response.Resp(c, response.Response{
		Data:     res,
		Messages: response.Messages{"Dashboard data successfully retrieved"},
		Code:     fiber.StatusOK,
	})
}
