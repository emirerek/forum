package handler

import (
	"forum/model"
	"forum/store"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ReplyHandler struct {
	store *store.ReplyStore
}

func NewReplyHandler(store *store.ReplyStore) *ReplyHandler {
	return &ReplyHandler{store}
}

func (handler *ReplyHandler) GetReply(c echo.Context) error {
	replyID, err := strconv.Atoi(c.Param("replyId"))
	if err != nil {
		return ErrPathParam
	}
	reply, err := handler.store.SelectReply(c.Request().Context(), replyID)
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
	var replyPost model.ReplyPost
	if err := c.Bind(&replyPost); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	replyInsert := &model.ReplyInsert{
		AccountID: replyPost.AccountID,
		ThreadID:  replyPost.ThreadID,
		Content:   replyPost.Content,
	}
	err := handler.store.InsertReply(c.Request().Context(), replyInsert)
	if err != nil {
		return handleDatabaseError(err)
	}
	return c.JSON(http.StatusOK, "reply created successfully")
}

func (handler *ReplyHandler) PatchReply(c echo.Context) error {
	var replyPatch model.ReplyPatch
	replyID, err := strconv.Atoi(c.Param("replyId"))
	if err != nil {
		return ErrPathParam
	}
	if err := c.Bind(&replyPatch); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(&replyPatch); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	replyUpdate := &model.ReplyUpdate{
		ID:      replyID,
		Content: replyPatch.Content,
	}
	err = handler.store.UpdateReply(c.Request().Context(), replyUpdate)
	if err != nil {
		return handleDatabaseError(err)
	}
	return c.JSON(http.StatusOK, "reply updated successfully")
}

func (handler *ReplyHandler) DeleteReply(c echo.Context) error {
	replyID, err := strconv.Atoi(c.Param("replyId"))
	if err != nil {
		return ErrPathParam
	}
	err = handler.store.DeleteReply(c.Request().Context(), replyID)
	if err != nil {
		return handleDatabaseError(err)
	}
	return c.JSON(http.StatusOK, "reply deleted successfully")
}
