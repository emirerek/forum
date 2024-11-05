package handler

import (
	"forum/model"
	"forum/store"
	"forum/utility"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type AccountHandler struct {
	store *store.AccountStore
}

func NewAccountHandler(store *store.AccountStore) *AccountHandler {
	return &AccountHandler{store}
}

func (handler *AccountHandler) GetAccount(c echo.Context) error {
	accountID, err := strconv.Atoi(c.Param("accountId"))
	if err != nil {
		return ErrPathParam
	}
	account, err := handler.store.SelectAccount(c.Request().Context(), accountID)
	if err != nil {
		return handleDatabaseError(err)
	}
	return c.JSON(http.StatusOK, account)
}

func (handler *AccountHandler) GetAccounts(c echo.Context) error {
	accounts, err := handler.store.SelectAccounts(c.Request().Context())
	if err != nil {
		return handleDatabaseError(err)
	}
	return c.JSON(http.StatusOK, accounts)
}

func (handler *AccountHandler) PostAccount(c echo.Context) error {
	var accountPost *model.AccountPost
	if err := c.Bind(&accountPost); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(&accountPost); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	passwordHash, err := utility.HashPassword(accountPost.Password)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError).SetInternal(err)
	}
	accountInsert := &model.AccountInsert{
		Username:     accountPost.Username,
		Email:        accountPost.Email,
		PasswordHash: passwordHash,
	}
	err = handler.store.InsertAccount(c.Request().Context(), accountInsert)
	if err != nil {
		return handleDatabaseError(err)
	}
	return c.JSON(http.StatusCreated, "account created successfully")
}

func (handler *AccountHandler) DeleteAccount(c echo.Context) error {
	accountID, err := strconv.Atoi(c.Param("accountId"))
	if err != nil {
		return ErrPathParam
	}
	err = handler.store.DeleteAccount(c.Request().Context(), accountID)
	if err != nil {
		return handleDatabaseError(err)
	}
	return c.JSON(http.StatusOK, "account deleted successfully")
}
