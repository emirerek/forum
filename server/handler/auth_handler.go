package handler

import (
	"errors"
	"forum/model"
	"forum/store"
	"forum/utility"
	"io"
	"net/http"
	"os"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type SessionInfo struct {
	AccountId uint
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

func (handler *AuthHandler) Register(c echo.Context) error {
	var register model.AccountPost
	if err := c.Bind(&register); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(&register); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	profilePic, err := c.FormFile("profilePic")
	var profilePath string
	if err == nil && profilePic != nil {
		file, err := profilePic.Open()
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "failed to open profile picture")
		}
		defer file.Close()
		profilePath = "uploads/profilepics/" + profilePic.Filename
		out, err := os.Create(profilePath)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "failed to create profile picture file")
		}
		defer out.Close()
		if _, err = io.Copy(out, file); err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "failed to save profile picture")
		}
	}
	hash, err := utility.HashPassword(register.Password)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to hash password")
	}
	account := &model.AccountInsert{
		Username:     register.Username,
		Email:        register.Email,
		ProfilePath:  profilePath,
		PasswordHash: hash,
	}
	if err := handler.store.InsertAccount(c.Request().Context(), account); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, "registered successfully")
}

func (handler *AuthHandler) Login(c echo.Context) error {
	var login model.AccountPostLogin
	if err := c.Bind(&login); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if login.Password == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "password is required")
	}
	if login.Username == "" && login.Email == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "username or email is required")
	}
	credentials, err := handler.store.SelectAccountCredentials(c.Request().Context(), login.Username, login.Email)
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
	session.Values["accountId"] = credentials.Id
	session.Values["accountUsername"] = credentials.Username
	session.Values["isAdmin"] = credentials.IsAdmin
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

func (handler *AuthHandler) Me(c echo.Context) error {
	session, err := handler.session.Get(c.Request(), "account")
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to get session")
	}
	auth, ok := session.Values["authenticated"].(bool)
	if !ok || !auth {
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{"authenticated": false})
	}
	isAdmin := false
	if v, ok := session.Values["isAdmin"].(bool); ok {
		isAdmin = v
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"authenticated": true,
		"accountId":     session.Values["accountId"],
		"username":      session.Values["accountUsername"],
		"isAdmin":       isAdmin,
	})
}
