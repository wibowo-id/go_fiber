package repository

import (
	"bytes"
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"go_fiber_wibowo/app/database/schema"
	authRequest "go_fiber_wibowo/app/module/auth/request"
	authResponse "go_fiber_wibowo/app/module/auth/response"
	"go_fiber_wibowo/app/module/user/request"
	"go_fiber_wibowo/app/module/user/response"
	paginationRequest "go_fiber_wibowo/app/request"
	"go_fiber_wibowo/internal/bootstrap/database"
	"go_fiber_wibowo/utils/config"
	"go_fiber_wibowo/utils/paginator"
	"io"
	"log"
	"net/http"
	"strings"
)

type userRepository struct {
	DB                *database.Database
	Cfg               *config.Config
	OtpConfirmRequest authRequest.OtpConfirmRequest
}

//go:generate mockgen -destination=user_repository_mock.go -package=repository . UserRepository
type UserRepository interface {
	All(req paginationRequest.Pagination, c *fiber.Ctx) (user []*response.UserResponse, paging paginator.Pagination, err error)
	FindOne(id string) (user *response.UserResponse, err error)
	CheckExist(id string) (res *schema.User, err error)
	Create(user *schema.User) (err error)
	Update(id string, user *schema.User) (err error)
	Delete(req request.UserDeleteRequest, id string) (err error)

	FindUserByEmail(email string) (user *schema.User, err error)
	FindUserByPhone(telepon string) (user *schema.User, err error)
	CheckUserByEmail(email string) (user *schema.User, err error)
	CheckUserByPhone(telepon string) (user *schema.User, err error)
	FindUserByOtp(telepon, otp string) (user *schema.User, err error)
	FindUserById(Id string) (user *schema.User, err error)
	RemoveToken(Id string) (user *schema.User, err error)
	CreateUser(user *schema.User) (res *schema.User, err error)
	UpdateToken(user schema.User) (res *schema.User, err error)
	ChangePassword(newPassword, UserId string) (res schema.User, err error)
	UpdateProfile(user *authRequest.UpdateProfileRequest, Id uuid.UUID) error
	FindUserByToken(user schema.User) (res *schema.User, err error)
	AddOTP(phone, Otp string) error
	SendMessage(phone, message string) error
	CreateForgotPassword(req *schema.ForgotPassword) (res *schema.ForgotPassword, err error)
	GetDashboard() (res *authResponse.Dashboard, err error)
}

func NewUserRepository(db *database.Database) UserRepository {
	return &userRepository{
		DB: db,
	}
}

func (_i *userRepository) All(req paginationRequest.Pagination, c *fiber.Ctx) (user []*response.UserResponse, paging paginator.Pagination, err error) {
	isClient := c.Query("is_client")
	var count int64

	query := _i.DB.DB.Model(&schema.User{}).
		Select("users.*", "roles.role as nama_role").
		Joins("left join roles on roles.id=users.role_id")

	if isClient == "true" {
		query.Where("users.role_id", nil).Or("users.role_id", "00000000-0000-0000-0000-000000000000")
	} else {
		query.Where("users.role_id IS NOT NULL").Where("users.role_id != 00000000-0000-0000-0000-000000000000")
	}

	query.Count(&count)

	req.Pagination.Count = count
	req.Pagination = paginator.Paging(req.Pagination)

	err = query.Offset(req.Pagination.Offset).Limit(req.Pagination.Limit).Find(&user).Error
	if err != nil {
		return
	}

	paging = *req.Pagination

	return
}

func (_i *userRepository) FindOne(id string) (user *response.UserResponse, err error) {
	if err := _i.DB.DB.First(&user, uuid.Must(uuid.Parse(id))).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (_i *userRepository) CheckExist(id string) (res *schema.User, err error) {
	if err := _i.DB.DB.Select("id").First(&res, uuid.Must(uuid.Parse(id))).Error; err != nil {
		return nil, err
	}

	return res, nil
}

func (_i *userRepository) Create(user *schema.User) (err error) {
	return _i.DB.DB.Create(user).Error
}

func (_i *userRepository) Update(id string, user *schema.User) (err error) {
	return _i.DB.DB.Model(&schema.User{}).
		Where(&schema.User{Id: uuid.Must(uuid.Parse(id))}).
		Updates(user).Error
}

func (_i *userRepository) Delete(req request.UserDeleteRequest, id string) error {
	if err := _i.DB.DB.Table("users").Where("id", id).Updates(map[string]interface{}{"is_active": req.IsActive, "deleted_at": req.DeletedAt, "deleted_by": req.DeletedBy}).Error; err != nil {
		return err
	}

	return nil
}

func (_i *userRepository) FindUserByEmail(email string) (user *schema.User, err error) {
	qState := "SELECT * FROM users "
	qState = qState + "WHERE email = '" + email + "' "
	qState = qState + "AND is_active = true "
	err = _i.DB.DB.Raw(qState).Scan(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (_i *userRepository) FindUserByPhone(telepon string) (user *schema.User, err error) {
	qState := "SELECT * FROM users "
	qState = qState + "WHERE telepon = '" + telepon + "' "
	qState = qState + "AND is_active = true "
	err = _i.DB.DB.Raw(qState).Scan(&user).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (_i *userRepository) CheckUserByEmail(email string) (user *schema.User, err error) {
	qState := "SELECT * FROM users "
	qState = qState + "WHERE email = '" + email + "' "
	err = _i.DB.DB.Raw(qState).Scan(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (_i *userRepository) CheckUserByPhone(telepon string) (user *schema.User, err error) {
	qState := "SELECT * FROM users "
	qState = qState + "WHERE telepon = '" + telepon + "' "
	err = _i.DB.DB.Raw(qState).Scan(&user).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (_i *userRepository) FindUserByOtp(telepon, otp string) (user *schema.User, err error) {
	qState := "SELECT * FROM users "
	qState = qState + "WHERE telepon = '" + telepon + "' "
	qState = qState + "AND otp = '" + otp + "' "

	err = _i.DB.DB.Raw(qState).Scan(&user).Error
	if err != nil {
		return nil, err
	}

	qStateUpdate := "UPDATE users "
	qStateUpdate = qStateUpdate + "SET is_active = true "
	qStateUpdate = qStateUpdate + "WHERE telepon = '" + telepon + "' "
	qStateUpdate = qStateUpdate + "AND otp = '" + otp + "' "

	err = _i.DB.DB.Raw(qStateUpdate).Scan(&user).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (_i *userRepository) FindUserById(Id string) (user *schema.User, err error) {
	qState := "SELECT * FROM users "
	qState = qState + "WHERE id = '" + Id + "' "
	qState = qState + "AND is_active = true "

	err = _i.DB.DB.Raw(qState).Scan(&user).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (_i *userRepository) RemoveToken(Id string) (user *schema.User, err error) {
	qState := "UPDATE users "
	qState = qState + "SET remember_token = null "
	qState = qState + "WHERE id = '" + Id + "' "

	err = _i.DB.DB.Raw(qState).Scan(&user).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (_i *userRepository) CreateForgotPassword(req *schema.ForgotPassword) (res *schema.ForgotPassword, err error) {
	err = _i.DB.DB.Create(req).Error
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (_i *userRepository) UpdateToken(user schema.User) (res *schema.User, err error) {
	if err := _i.DB.DB.Table("users").Where("id", user.Id).Update("remember_token", user.RememberToken).Error; err != nil {
		return res, err
	}

	return res, nil
}

func (_i *userRepository) ChangePassword(newPassword, UserId string) (res schema.User, err error) {
	if err := _i.DB.DB.Table("users").Where("id", UserId).Updates(map[string]interface{}{"password": newPassword, "remember_token": nil}).Error; err != nil {
		return res, err
	}

	return res, nil
}

func (_i *userRepository) FindUserByToken(user schema.User) (res *schema.User, err error) {
	qState := "SELECT * FROM users "
	qState = qState + "WHERE remember_token = '" + user.RememberToken + "' "

	err = _i.DB.DB.Raw(qState).Scan(&user).Error
	if err != nil {
		return nil, err
	}

	res = &user

	return res, nil
}

func (_i *userRepository) CreateUser(user *schema.User) (res *schema.User, err error) {
	if err := _i.DB.DB.Create(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (_i *userRepository) UpdateProfile(user *authRequest.UpdateProfileRequest, Id uuid.UUID) error {
	if err := _i.DB.DB.Table("users").Where("id = ?", Id).Updates(&user).Error; err != nil {
		return err
	}

	return nil
}

func (_i *userRepository) AddOTP(phone, Otp string) error {
	var err error
	u := authRequest.OtpConfirmRequest{}
	u.Otp = Otp

	err = _i.DB.DB.Table("users").Where("telepon = ?", phone).Updates(&u).Error
	if err != nil {
		return err
	}
	return err
}

func (_i *userRepository) SendMessage(recipient, message string) error {
	apiKey := _i.DB.Cfg.App.Wa.ApiKey
	sender := _i.DB.Cfg.App.Wa.Phone
	url := _i.DB.Cfg.App.Wa.UrlText

	payload := authRequest.PayloadSendMessage{}
	payload.ApiKey = apiKey
	payload.Sender = sender
	payload.Number = recipient
	payload.Message = strings.Replace(message, "APP_NAME", _i.DB.Cfg.App.Name, 1)

	buf, err := json.Marshal(payload)
	if err != nil {
		log.Fatal(err)
	}
	reader := bytes.NewReader(buf)

	req, _ := http.NewRequest("POST", url, reader)

	req.Header.Add("Content-Type", "application/json")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	_, _ = io.ReadAll(res.Body)

	return err
}

func (_i *userRepository) GetDashboard() (res *authResponse.Dashboard, err error) {
	dashboard := &authResponse.Dashboard{}

	qUserClient := "SELECT COUNT(id) as user_client FROM users "
	qUserClient = qUserClient + "WHERE role_id = '00000000-0000-0000-0000-000000000000' "
	qUserClient = qUserClient + "AND deleted_at IS NULL "

	if err := _i.DB.DB.Raw(qUserClient).Scan(&dashboard.UserClient).Error; err != nil {
		return nil, err
	}

	qUserAdmin := "SELECT COUNT(id) as user_client FROM users "
	qUserAdmin = qUserAdmin + "WHERE role_id != '00000000-0000-0000-0000-000000000000' "
	qUserAdmin = qUserAdmin + "AND deleted_at IS NULL "

	if err := _i.DB.DB.Raw(qUserAdmin).Scan(&dashboard.UserAdmin).Error; err != nil {
		return nil, err
	}

	return dashboard, nil
}
