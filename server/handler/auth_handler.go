package handler

import (
	"errors"
	"forum/model"
	"forum/store"
	"forum/utility"
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type SessionInfo struct {
	AccountID uint
	Username  string
	Email     string
}

type AuthHandler struct {
	store   *store.AccountStore
	session *sessions.CookieStore
}

func NewAuthHandler(store *store.AccountStore, session *sessions.CookieStore) *AuthHandler {
	return &AuthHandler{store, session}
}

func (handler *AuthHandler) Login(c echo.Context) error {
	var login *model.AccountPost
	if err := c.Bind(&login); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(&login); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	credentials, err := handler.store.SelectAccountCredentials(c.Request().Context(), login.Password, login.Email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ErrInvalidCredentials
		}
		return ErrLogin.SetInternal(err)
	}
	success, err := utility.VerifyPassword(credentials.PasswordHash, login.Password)
	if err != nil {
		return ErrLogin.SetInternal(err)
	}
	if !success {
		return ErrInvalidPassword
	}
	session, err := handler.session.Get(c.Request(), "account")
	if err != nil {
		return ErrLogin.SetInternal(err)
	}
	session.Values["authenticated"] = true
	session.Values["accountID"] = credentials.ID
	session.Values["accountUsername"] = credentials.Username
	err = session.Save(c.Request(), c.Response().Writer)
	if err != nil {
		return ErrLogin.SetInternal(err)
	}
	return c.JSON(http.StatusOK, "logged in successfuly")
}

func (handler *AuthHandler) Logout(c echo.Context) error {
	session, err := handler.session.Get(c.Request(), "account")
	if err != nil {
		return ErrLogout.SetInternal(err)
	}
	session.Options.MaxAge = -1 // delete the session
	err = session.Save(c.Request(), c.Response().Writer)
	if err != nil {
		return ErrLogout.SetInternal(err)
	}
	return c.JSON(http.StatusOK, "logged out successfully")
}
