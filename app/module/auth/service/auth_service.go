package service

import (
	"errors"
	"github.com/google/uuid"
	"go_fiber_wibowo/app/database/schema"
	"go_fiber_wibowo/app/module/auth/request"
	"go_fiber_wibowo/app/module/auth/response"
	user_repo "go_fiber_wibowo/app/module/user/repository"
	"go_fiber_wibowo/utils/helpers"
	"go_fiber_wibowo/utils/otp"
	"go_fiber_wibowo/utils/token"
	"golang.org/x/crypto/bcrypt"
	"time"
)

// AuthService
type authService struct {
	userRepo user_repo.UserRepository
}

// define interface of IAuthService
//
//go:generate mockgen -destination=article_service_mock.go -package=service . AuthService
type AuthService interface {
	LoginAdmin(req request.LoginRequest) (res *response.LoginResponse, err error)
	Login(req request.LoginRequest) (res *response.LoginResponse, err error)
	Register(req request.RegisterRequest) (res response.RegisterResponse, err error)
	ForgotPassword(req request.ForgotPasswordRequest) (Phone string, err error)
	ChangePassword(req request.ChangePasswordRequest, UserId string) (res response.RegisterResponse, err error)
	RequestOTP(req request.OtpRequest) (OtpResp response.OtpResponse, err error)
	VerifyOTP(req request.OtpConfirmRequest) (OtpResp response.OtpResponse, err error)
	Profile(Id string) (res *schema.User, err error)
	Logout(Id string) (res *schema.User, err error)
	UpdateProfile(req *request.UpdateProfileRequest, Id uuid.UUID) (res response.UpdateProfileResponse, err error)
	UpdateFirebaseToken(req request.UpdateFirebaseTokenRequest) (res response.UpdateFirebaseResponse, err error)
	UpdateToken(token string, Id uuid.UUID) (res response.UpdateProfileResponse, err error)
	GetUserByToken(token string) (res *schema.User, err error)
	GetUserById(id string) (res *schema.User, err error)
}

// init AuthService
func NewAuthService(userRepo user_repo.UserRepository) AuthService {
	return &authService{
		userRepo: userRepo,
	}
}

func (_i *authService) LoginAdmin(req request.LoginRequest) (res *response.LoginResponse, err error) {
	user, err := _i.userRepo.FindUserByEmail(req.Email)
	if err != nil {
		return
	}

	if user == nil {
		err = errors.New("user not found")
		return
	}

	if !user.ComparePassword(req.Password) {
		err = errors.New("password not match")
		return
	}

	claims, err := token.GenerateToken(user.Id)
	if err != nil {
		return
	}

	extractToken, err := token.ExtractTokenID(claims)
	_i.UpdateToken(claims, user.Id)

	res = extractToken

	return
}

func (_i *authService) Login(req request.LoginRequest) (res *response.LoginResponse, err error) {
	user, err := _i.userRepo.FindUserByEmail(req.Email)
	if err != nil {
		return
	}

	if user == nil {
		err = errors.New("user not found")
		return
	}

	if !user.ComparePassword(req.Password) {
		err = errors.New("password not match")
		return
	}

	claims, err := token.GenerateToken(user.Id)
	if err != nil {
		return
	}

	extractToken, err := token.ExtractTokenID(claims)
	_i.UpdateToken(claims, user.Id)

	res = extractToken

	return
}

func (_i *authService) Register(req request.RegisterRequest) (res response.RegisterResponse, err error) {
	hashPass := helpers.GeneratePassword(req.Password)
	req.Telepon = helpers.VerifyFormatPhone(req.Telepon)

	user, err := _i.userRepo.CheckUserByPhone(req.Telepon)
	if err != nil {
		return
	}

	if user != nil {
		err = errors.New("telepon already exists")
		return
	}

	userEmail, err := _i.userRepo.CheckUserByEmail(req.Email)
	if err != nil {
		return
	}

	if userEmail != nil {
		err = errors.New("email already exists")
		return
	}

	err = helpers.VerifyPassword(req.Password, req.ConfirmPassword)
	if err != nil && errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
		err = errors.New("password not match")
		return
	}

	newUser := &schema.User{
		Id:          uuid.New(),
		NamaLengkap: req.NamaLengkap,
		Telepon:     req.Telepon,
		Email:       req.Email,
		Password:    hashPass,
	}

	user, err = _i.userRepo.CreateUser(newUser)
	if err != nil {
		return
	}

	res.Id = uuid.New()
	res.NamaLengkap = user.NamaLengkap
	res.Telepon = user.Telepon
	res.Email = user.Email

	return
}

func (_i *authService) ForgotPassword(req request.ForgotPasswordRequest) (Phone string, err error) {
	user, err := _i.userRepo.FindUserByEmail(req.Email)
	if err != nil {
		return
	}

	if user == nil {
		err = errors.New("email not found")
		return
	}

	data := &schema.ForgotPassword{}
	data.Id = uuid.New()
	data.Email = user.Email
	data.Key = helpers.RandStringBytes(36)
	data.ExpiredDate = time.Now().Local().Add(time.Hour * time.Duration(1))
	data.CreatedAt = time.Now()
	data.CreatedBy = user.Id

	_, err = _i.userRepo.CreateForgotPassword(data)
	if err != nil {
		return
	}

	Phone = user.Telepon

	return
}

func (_i *authService) ChangePassword(req request.ChangePasswordRequest, UserId string) (res response.RegisterResponse, err error) {
	if req.Password != req.ConfirmPassword {
		err = errors.New("new password not match !")
		return
	}

	hashPass := helpers.GeneratePassword(req.Password)

	_, err = _i.userRepo.ChangePassword(hashPass, UserId)
	if err != nil {
		return
	}

	return
}

func (_i *authService) RequestOTP(req request.OtpRequest) (OtpResp response.OtpResponse, err error) {
	user, err := _i.userRepo.FindUserByPhone(req.Telepon)
	if err != nil {
		return
	}

	if user == nil {
		err = errors.New("telepon not found")
		return
	}

	getOTP := otp.GenerateOtp()
	message := "*APP_NAME OTP*\r\nJangan berikan kode ini kepada orang lain!\r\n[#] Kode OTP Anda adalah *" + getOTP + "*"
	_i.userRepo.SendMessage(req.Telepon, message)
	_i.userRepo.AddOTP(req.Telepon, getOTP)

	return
}

func (_i *authService) VerifyOTP(req request.OtpConfirmRequest) (OtpResp response.OtpResponse, err error) {
	req.Telepon = helpers.VerifyFormatPhone(req.Telepon)

	user, err := _i.userRepo.FindUserByOtp(req.Telepon, req.Otp)
	if err != nil {
		return
	}

	if user == nil {
		err = errors.New("telepon & otp not found")
		return
	}

	return
}

func (_i *authService) Profile(Id string) (res *schema.User, err error) {
	res, err = _i.userRepo.FindUserById(Id)
	if err != nil {
		err = errors.New("user not found")
		return
	}

	return
}

func (_i *authService) Logout(Id string) (res *schema.User, err error) {
	res, err = _i.userRepo.RemoveToken(Id)
	if err != nil {
		err = errors.New("user not found")
		return
	}

	return
}

func (_i *authService) UpdateToken(token string, Id uuid.UUID) (res response.UpdateProfileResponse, err error) {
	userToken := schema.User{}
	userToken.RememberToken = token
	userToken.Id = Id

	user, err := _i.userRepo.UpdateToken(userToken)
	if err != nil {
		return
	}

	if user == nil {
		err = errors.New("update token failed")
		return
	}

	return
}

func (_i *authService) GetUserByToken(token string) (res *schema.User, err error) {
	userToken := schema.User{}
	userToken.RememberToken = token

	res, err = _i.userRepo.FindUserByToken(userToken)
	if err != nil {
		return
	}

	if res == nil {
		err = errors.New("user with token not found")
		return
	}

	return
}

func (_i *authService) GetUserById(id string) (res *schema.User, err error) {
	res, err = _i.userRepo.FindUserById(id)
	if err != nil {
		return
	}

	if res == nil {
		err = errors.New("user by id not found")
		return
	}

	return
}

func (_i *authService) UpdateProfile(req *request.UpdateProfileRequest, Id uuid.UUID) (res response.UpdateProfileResponse, err error) {
	err = _i.userRepo.UpdateProfile(req, Id)
	if err != nil {
		err = errors.New("update profile failed")
		return
	}

	res.Email = req.Email
	res.Telepon = req.Telepon
	res.NamaLengkap = req.NamaLengkap

	return
}

func (_i *authService) UpdateFirebaseToken(req request.UpdateFirebaseTokenRequest) (res response.UpdateFirebaseResponse, err error) {
	return
}
