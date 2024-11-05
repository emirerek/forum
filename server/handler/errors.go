package handler

import (
	"errors"
	"forum/store"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

var ErrQueryParam = echo.NewHTTPError(http.StatusBadRequest, "failed to parse query parameters")
var ErrPathParam = echo.NewHTTPError(http.StatusBadRequest, "failed to parse path parameter")
var ErrNotFound = echo.NewHTTPError(http.StatusNotFound, "resource not found")
var ErrNotUnique = echo.NewHTTPError(http.StatusConflict, "resource not unique")
var ErrFKViolation = echo.NewHTTPError(http.StatusUnprocessableEntity, "referenced resource not found")

var ErrInvalidCredentials = echo.NewHTTPError(http.StatusUnauthorized, "invalid credentials")
var ErrInvalidPassword = echo.NewHTTPError(http.StatusUnauthorized, "invalid password")
var ErrLogin = echo.NewHTTPError(http.StatusInternalServerError, "failed to log in")
var ErrLogout = echo.NewHTTPError(http.StatusInternalServerError, "failed to log out")

var ErrDatabase = echo.NewHTTPError(http.StatusInternalServerError, "database error")

func handleDatabaseError(err error) *echo.HTTPError {
	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		return ErrNotFound
	case errors.Is(err, store.ErrNoRecordsAffected):
		return ErrNotFound
	case errors.Is(err, gorm.ErrDuplicatedKey):
		return ErrNotUnique
	case errors.Is(err, gorm.ErrForeignKeyViolated):
		return ErrFKViolation
	default:
		return ErrDatabase.SetInternal(err)
	}
}
