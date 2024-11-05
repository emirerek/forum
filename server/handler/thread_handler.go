package handler

import (
	"forum/model"
	"forum/store"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ThreadHandler struct {
	store *store.ThreadStore
}

func NewThreadHandler(store *store.ThreadStore) *ThreadHandler {
	return &ThreadHandler{store}
}

func (handler *ThreadHandler) GetThread(c echo.Context) error {
	threadID, err := strconv.Atoi(c.Param("threadId"))
	if err != nil {
		return ErrPathParam
	}
	thread, err := handler.store.SelectThread(c.Request().Context(), threadID)
	if err != nil {
		return handleDatabaseError(err)
	}
	return c.JSON(http.StatusOK, thread)
}

func (handler *ThreadHandler) GetThreads(c echo.Context) error {
	var threads []model.Thread
	var subforumID int
	var err error
	if c.QueryParams().Has("subforumID") {
		subforumID, err = strconv.Atoi(c.QueryParam("subforumID"))
		if err != nil {
			return ErrQueryParam
		}
		threads, err = handler.store.SelectThreadsBySubforum(c.Request().Context(), subforumID)
	} else {
		threads, err = handler.store.SelectThreads(c.Request().Context())
	}
	if err != nil {
		return handleDatabaseError(err)
	}
	return c.JSON(http.StatusOK, threads)
}

func (handler *ThreadHandler) PostThread(c echo.Context) error {
	var threadPost model.ThreadPost
	if err := c.Bind(&threadPost); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(&threadPost); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	threadInsert := &model.ThreadInsert{
		AccountID:  threadPost.AccountID,
		SubforumID: threadPost.SubforumID,
		Title:      threadPost.Title,
		Content:    threadPost.Content,
	}
	err := handler.store.InsertThread(c.Request().Context(), threadInsert)
	if err != nil {
		return handleDatabaseError(err)
	}
	return c.JSON(http.StatusCreated, "thread created successfully")
}

func (handler *ThreadHandler) PatchThread(c echo.Context) error {
	var threadPatch model.ThreadPatch
	threadID, err := strconv.Atoi(c.Param("threadId"))
	if err != nil {
		return ErrPathParam
	}
	if err := c.Bind(&threadPatch); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(&threadPatch); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	threadUpdate := &model.ThreadUpdate{
		ID:      threadID,
		Title:   threadPatch.Title,
		Content: threadPatch.Content,
	}
	err = handler.store.UpdateThread(c.Request().Context(), threadUpdate)
	if err != nil {
		return handleDatabaseError(err)
	}
	return c.JSON(http.StatusOK, "thread updated successfully")
}

func (handler *ThreadHandler) DeleteThread(c echo.Context) error {
	threadID, err := strconv.Atoi(c.Param("threadId"))
	if err != nil {
		return ErrPathParam
	}
	err = handler.store.DeleteThread(c.Request().Context(), threadID)
	if err != nil {
		return handleDatabaseError(err)
	}
	return c.JSON(http.StatusOK, "thread deleted successfully")
}
