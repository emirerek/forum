package handler

import (
	"forum/model"
	"forum/store"
	"net/http"
	"strconv"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo/v4"
)

type ThreadHandler struct {
	store   *store.ThreadStore
	session *sessions.CookieStore
}

func NewThreadHandler(store *store.ThreadStore, session *sessions.CookieStore) *ThreadHandler {
	return &ThreadHandler{store, session}
}

func (handler *ThreadHandler) GetThread(c echo.Context) error {
	threadId, err := strconv.Atoi(c.Param("threadId"))
	if err != nil {
		return ErrPathParam
	}
	thread, err := handler.store.SelectThread(c.Request().Context(), threadId)
	if err != nil {
		return handleDatabaseError(err)
	}
	return c.JSON(http.StatusOK, thread)
}

func (handler *ThreadHandler) GetThreads(c echo.Context) error {
	var threads []model.Thread
	var subforumId int
	var err error
	if c.QueryParams().Has("subforumId") {
		subforumId, err = strconv.Atoi(c.QueryParam("subforumId"))
		if err != nil {
			return ErrQueryParam
		}
		threads, err = handler.store.SelectThreadsBySubforum(c.Request().Context(), subforumId)
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
		AccountId:  threadPost.AccountId,
		SubforumId: threadPost.SubforumId,
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
	session, err := handler.session.Get(c.Request(), "account")
	if err != nil || session.Values["authenticated"] != true {
		return c.JSON(http.StatusUnauthorized, "authentication required")
	}
	userId, _ := session.Values["accountId"].(int)
	isAdmin := session.Values["isAdmin"] == true
	threadId, err := strconv.Atoi(c.Param("threadId"))
	if err != nil {
		return ErrPathParam
	}
	thread, err := handler.store.SelectThread(c.Request().Context(), threadId)
	if err != nil {
		return handleDatabaseError(err)
	}
	if !isAdmin && int(thread.AccountId) != userId {
		return c.JSON(http.StatusForbidden, "not allowed to edit this thread")
	}
	var threadPatch model.ThreadPatch
	if err := c.Bind(&threadPatch); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(&threadPatch); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	threadUpdate := &model.ThreadUpdate{
		Id:      threadId,
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
	session, err := handler.session.Get(c.Request(), "account")
	if err != nil || session.Values["authenticated"] != true {
		return c.JSON(http.StatusUnauthorized, "authentication required")
	}
	userId, _ := session.Values["accountId"].(int)
	isAdmin := session.Values["isAdmin"] == true
	threadId, err := strconv.Atoi(c.Param("threadId"))
	if err != nil {
		return ErrPathParam
	}
	thread, err := handler.store.SelectThread(c.Request().Context(), threadId)
	if err != nil {
		return handleDatabaseError(err)
	}
	if !isAdmin && int(thread.AccountId) != userId {
		return c.JSON(http.StatusForbidden, "not allowed to delete this thread")
	}
	err = handler.store.DeleteThread(c.Request().Context(), threadId)
	if err != nil {
		return handleDatabaseError(err)
	}
	return c.JSON(http.StatusOK, "thread deleted successfully")
}
