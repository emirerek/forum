package store

import (
	"context"
	"forum/model"

	"gorm.io/gorm"
)

type ReplyStore struct {
	db *gorm.DB
}

func NewReplyStore(db *gorm.DB) *ReplyStore {
	return &ReplyStore{db}
}

func (store *ReplyStore) SelectReply(ctx context.Context, replyId int) (model.Reply, error) {
	var reply model.Reply
	err := store.db.
		WithContext(ctx).
		Preload("Account", func(db *gorm.DB) *gorm.DB {
			return db.Select("id", "username")
		}).
		First(&reply, replyId).Error
	if err != nil {
		return reply, err
	}
	return reply, nil
}

func (store *ReplyStore) SelectReplies(ctx context.Context) ([]model.Reply, error) {
	var replies []model.Reply
	err := store.db.
		WithContext(ctx).
		Preload("Account", func(db *gorm.DB) *gorm.DB {
			return db.Select("id", "username")
		}).
		Find(&replies).Error
	if err != nil {
		return nil, err
	}
	return replies, nil
}

func (store *ReplyStore) InsertReply(ctx context.Context, replyInsert *model.ReplyInsert) error {
	reply := &model.Reply{
		AccountID: replyInsert.AccountID,
		ThreadID:  replyInsert.ThreadID,
		Content:   replyInsert.Content,
	}
	err := store.db.
		WithContext(ctx).
		Create(reply).Error
	if err != nil {
		return err
	}
	return nil
}

func (store *ReplyStore) UpdateReply(ctx context.Context, replyUpdate *model.ReplyUpdate) error {
	reply := &model.Reply{
		Content: replyUpdate.Content,
	}
	result := store.db.
		WithContext(ctx).
		Model(&reply).
		Where("id = ?", replyUpdate.ID).
		Updates(reply)
	if err := result.Error; err != nil {
		return err
	}
	if rows := result.RowsAffected; rows != 1 {
		return ErrNoRecordsAffected
	}
	return nil
}

func (store *ReplyStore) DeleteReply(ctx context.Context, replyId int) error {
	result := store.db.
		WithContext(ctx).
		Delete(&model.Reply{}, replyId)
	if err := result.Error; err != nil {
		return err
	}
	if rows := result.RowsAffected; rows != 1 {
		return ErrNoRecordsAffected
	}
	return nil
}
