package handler

import (
	"forum/model"
	"forum/store"
	"net/http"
	"strconv"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo/v4"
)

type ReplyHandler struct {
	store   *store.ReplyStore
	session *sessions.CookieStore
}

func NewReplyHandler(store *store.ReplyStore, session *sessions.CookieStore) *ReplyHandler {
	return &ReplyHandler{store, session}
}

func (handler *ReplyHandler) GetReply(c echo.Context) error {
	replyId, err := strconv.Atoi(c.Param("replyId"))
	if err != nil {
		return ErrPathParam
	}
	reply, err := handler.store.SelectReply(c.Request().Context(), replyId)
	if err != nil {
		return handleDatabaseError(err)
	}
	return c.JSON(http.StatusOK, reply)
}

func (handler *ReplyHandler) GetReplies(c echo.Context) error {
	replies, err := handler.store.SelectReplies(c.Request().Context())
	if err != nil {
		return handleDatabaseError(err)
	}
	return c.JSON(http.StatusOK, replies)
}

func (handler *ReplyHandler) PostReply(c echo.Context) error {
	session, err := handler.session.Get(c.Request(), "account")
	if err != nil || session.Values["authenticated"] != true {
		return c.JSON(http.StatusUnauthorized, "authentication required")
	}
	userId, _ := session.Values["accountId"].(int)
	var replyPost model.ReplyPost
	if err := c.Bind(&replyPost); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	replyInsert := &model.ReplyInsert{
		ThreadId:  replyPost.ThreadId,
		Content:   replyPost.Content,
		AccountId: userId,
	}
	if handler.store.InsertReply(c.Request().Context(), replyInsert) != nil {
		return handleDatabaseError(nil)
	}
	return c.JSON(http.StatusOK, "reply created successfully")
}

func (handler *ReplyHandler) PatchReply(c echo.Context) error {
	session, err := handler.session.Get(c.Request(), "account")
	if err != nil || session.Values["authenticated"] != true {
		return c.JSON(http.StatusUnauthorized, "authentication required")
	}
	userId, _ := session.Values["accountId"].(int)
	isAdmin := session.Values["isAdmin"] == true
	var replyPatch model.ReplyPatch
	replyId, err := strconv.Atoi(c.Param("replyId"))
	if err != nil {
		return ErrPathParam
	}
	reply, err := handler.store.SelectReply(c.Request().Context(), replyId)
	if err != nil {
		return handleDatabaseError(err)
	}
	if !isAdmin && int(reply.AccountId) != userId {
		return c.JSON(http.StatusForbidden, "not allowed to edit this reply")
	}
	if err := c.Bind(&replyPatch); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(&replyPatch); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	replyUpdate := &model.ReplyUpdate{
		Id:      replyId,
		Content: replyPatch.Content,
	}
	err = handler.store.UpdateReply(c.Request().Context(), replyUpdate)
	if err != nil {
		return handleDatabaseError(err)
	}
	return c.JSON(http.StatusOK, "reply updated successfully")
}

func (handler *ReplyHandler) DeleteReply(c echo.Context) error {
	session, err := handler.session.Get(c.Request(), "account")
	if err != nil || session.Values["authenticated"] != true {
		return c.JSON(http.StatusUnauthorized, "authentication required")
	}
	userId, _ := session.Values["accountId"].(int)
	isAdmin := session.Values["isAdmin"] == true
	replyId, err := strconv.Atoi(c.Param("replyId"))
	if err != nil {
		return ErrPathParam
	}
	reply, err := handler.store.SelectReply(c.Request().Context(), replyId)
	if err != nil {
		return handleDatabaseError(err)
	}
	if !isAdmin && int(reply.AccountId) != userId {
		return c.JSON(http.StatusForbidden, "not allowed to delete this reply")
	}
	err = handler.store.DeleteReply(c.Request().Context(), replyId)
	if err != nil {
		return handleDatabaseError(err)
	}
	return c.JSON(http.StatusOK, "reply deleted successfully")
}
