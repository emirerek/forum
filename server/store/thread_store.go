package store

import (
	"context"
	"forum/model"

	"gorm.io/gorm"
)

type ThreadStore struct {
	db *gorm.DB
}

func NewThreadStore(db *gorm.DB) *ThreadStore {
	return &ThreadStore{db}
}

func (store *ThreadStore) SelectThread(ctx context.Context, threadId int) (model.Thread, error) {
	var thread model.Thread
	err := store.db.
		WithContext(ctx).
		Preload("Account").
		Preload("Replies.Account").
		First(&thread, threadId).Error
	if err != nil {
		return thread, err
	}
	return thread, nil
}

func (store *ThreadStore) SelectThreads(ctx context.Context) ([]model.Thread, error) {
	var threads []model.Thread
	err := store.db.
		WithContext(ctx).
		Preload("Account").
		Find(&threads).Error
	if err != nil {
		return nil, err
	}
	return threads, nil
}

func (store *ThreadStore) SelectThreadsBySubforum(ctx context.Context, subforumId int) ([]model.Thread, error) {
	var threads []model.Thread
	err := store.db.
		WithContext(ctx).
		Preload("Account").
		Where("subforum_id = ?", subforumId).
		Find(&threads).Error
	if err != nil {
		return nil, err
	}
	return threads, nil
}

func (store *ThreadStore) InsertThread(ctx context.Context, threadInsert *model.ThreadInsert) error {
	thread := &model.Thread{
		AccountId:  threadInsert.AccountId,
		SubforumId: threadInsert.SubforumId,
		Title:      threadInsert.Title,
		Content:    threadInsert.Content,
	}
	err := store.db.
		WithContext(ctx).
		Create(thread).Error
	if err != nil {
		return err
	}
	return nil
}

func (store *ThreadStore) UpdateThread(ctx context.Context, threadUpdate *model.ThreadUpdate) error {
	thread := &model.Thread{
		Title:   threadUpdate.Title,
		Content: threadUpdate.Content,
	}
	result := store.db.
		WithContext(ctx).
		Model(&thread).
		Where("id = ?", threadUpdate.Id).
		Updates(thread)
	if err := result.Error; err != nil {
		return err
	}
	if result.RowsAffected != 1 {
		return ErrNoRecordsAffected
	}
	return nil
}

func (store *ThreadStore) DeleteThread(ctx context.Context, threadId int) error {
	result := store.db.
		WithContext(ctx).
		Delete(&model.Thread{}, threadId)
	if err := result.Error; err != nil {
		return err
	}
	if result.RowsAffected != 1 {
		return ErrNoRecordsAffected
	}
	return nil
}
