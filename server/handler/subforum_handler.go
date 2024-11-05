package handler

import (
	"forum/model"
	"forum/store"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type SubforumHandler struct {
	store *store.SubforumStore
}

func NewSubforumHandler(store *store.SubforumStore) *SubforumHandler {
	return &SubforumHandler{store}
}

func (handler *SubforumHandler) GetSubforum(c echo.Context) error {
	subforumID, err := strconv.Atoi(c.Param("subforumId"))
	if err != nil {
		return ErrPathParam
	}
	subforum, err := handler.store.SelectSubforum(c.Request().Context(), subforumID)
	if err != nil {
		return handleDatabaseError(err)
	}
	return c.JSON(http.StatusOK, subforum)
}

func (handler *SubforumHandler) GetSubforums(c echo.Context) error {
	subforums, err := handler.store.SelectSubforums(c.Request().Context())
	if err != nil {
		return handleDatabaseError(err)
	}
	return c.JSON(http.StatusOK, subforums)
}

func (handler *SubforumHandler) PostSubforum(c echo.Context) error {
	var subforumPost model.SubforumPost
	if err := c.Bind(&subforumPost); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(&subforumPost); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	subforumInsert := &model.SubforumInsert{
		Title:       subforumPost.Title,
		Description: subforumPost.Description,
	}
	err := handler.store.InsertSubforum(c.Request().Context(), subforumInsert)
	if err != nil {
		return handleDatabaseError(err)
	}
	return c.JSON(http.StatusCreated, "subforum created successfully")
}

func (handler *SubforumHandler) PatchSubforum(c echo.Context) error {
	var subforumPatch model.SubforumPatch
	subforumID, err := strconv.Atoi(c.Param("subforumId"))
	if err != nil {
		return ErrPathParam
	}
	if err := c.Bind(&subforumPatch); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(&subforumPatch); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	subforumUpdate := &model.SubforumUpdate{
		ID:          subforumID,
		Title:       subforumPatch.Title,
		Description: subforumPatch.Description,
	}
	err = handler.store.UpdateSubforum(c.Request().Context(), subforumUpdate)
	if err != nil {
		return handleDatabaseError(err)
	}
	return c.JSON(http.StatusOK, "subforum updated successfully")
}

func (handler *SubforumHandler) DeleteSubforum(c echo.Context) error {
	subforumID, err := strconv.Atoi(c.Param("subforumId"))
	if err != nil {
		return ErrPathParam
	}
	err = handler.store.DeleteSubforum(c.Request().Context(), subforumID)
	if err != nil {
		return handleDatabaseError(err)
	}
	return c.JSON(http.StatusOK, "subforum deleted successfully")
}
