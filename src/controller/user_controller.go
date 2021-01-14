package controller

import (
	"log"
	"regexp"
	"service-user-management/src/entity"
	"service-user-management/src/payload"
	"service-user-management/src/repository"
	"service-user-management/src/utils"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	ur repository.UserRepository
}

func ProvideUserController(ur repository.UserRepository) UserController {
	return UserController{ur: ur}
}

func (uc *UserController) AddUser(c echo.Context) error {
	var request payload.AddUser
	if err := c.Bind(&request); err != nil {
		return c.JSON(400, payload.ResponseWithMessage(400, err.Error()))
	}

	if request.BirthDate == "" || request.BirthPlace == "" || request.Password == "" || request.UserName == "" {
		return c.JSON(400, payload.ResponseWithMessage(400, "Missing params"))
	}

	s := regexp.MustCompile(`\s+`)
	if s.MatchString(request.UserName) {
		return c.JSON(400, payload.ResponseWithMessage(400, "user_name cannot contain spaces"))
	}

	birthDate, err := time.Parse("2006-01-02", request.BirthDate)
	if err != nil {
		return c.JSON(400, payload.ResponseWithMessage(400, "Date format must 'YYYY-MM-DD'"))
	}

	u := entity.User{
		BirthDate:  birthDate,
		UserName:   request.UserName,
		Password:   utils.HSHA256Handler(request.Password),
		BirthPlace: request.BirthPlace,
	}

	if err := uc.ur.InsertUser(&u); err != nil {
		return c.JSON(400, payload.ResponseWithMessage(400, err.Error()))
	}

	return c.JSON(200, payload.ResponseWithData(u))
}

func (uc *UserController) SearchUser(c echo.Context) error {
	search := c.QueryParam("search")
	search = "%" + strings.ToLower(search) + "%"
	users := uc.ur.FindByUsernameOrEmployeeIDOrBirthPlace(search)
	return c.JSON(200, payload.ResponseWithData(users))
}

func (uc *UserController) Login(c echo.Context) error {
	var request payload.LoginRequest
	if err := c.Bind(&request); err != nil {
		return c.JSON(400, payload.ResponseWithMessage(400, err.Error()))
	}

	if request.Password == "" || request.UserName == "" {
		return c.JSON(400, payload.ResponseWithMessage(400, "Missing params"))
	}

	u, err := uc.ur.FindByUsername(request.UserName)
	if err != nil {
		log.Println(err)
		return c.JSON(404, payload.ResponseWithMessage(404, "User not found"))
	}

	if u.Password != utils.HSHA256Handler(request.Password) {
		return c.JSON(401, payload.ResponseWithMessage(401, "Wrong password"))
	}

	tokens, err2 := uc.generateToken(u.UserName)
	if err2 != nil {
		return c.JSON(400, payload.ResponseWithMessage(400, err.Error()))
	}
	return c.JSON(200, payload.ResponseWithData(tokens))
}

func (uc *UserController) generateToken(subject string) (payload.LoginResponse, error) {
	at, err := utils.TokenGenerator(subject, 15)
	if err != nil {
		return payload.LoginResponse{}, err
	}

	rt, _ := utils.TokenGenerator(subject, 60)
	return payload.LoginResponse{AccessToken: at, RefreshToken: rt}, nil
}
