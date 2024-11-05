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

func (store *ThreadStore) SelectThread(ctx context.Context, threadID int) (model.Thread, error) {
	var thread model.Thread
	err := store.db.
		WithContext(ctx).
		Preload("Account", func(db *gorm.DB) *gorm.DB {
			return db.Select("id", "username")
		}).
		Preload("Replies.Account").
		First(&thread, threadID).Error
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

func (store *ThreadStore) SelectThreadsBySubforum(ctx context.Context, subforumID int) ([]model.Thread, error) {
	var threads []model.Thread
	err := store.db.
		WithContext(ctx).
		Preload("Account").
		Where("subforum_id = ?", subforumID).
		Find(&threads).Error
	if err != nil {
		return nil, err
	}
	return threads, nil
}

func (store *ThreadStore) InsertThread(ctx context.Context, threadInsert *model.ThreadInsert) error {
	thread := &model.Thread{
		AccountID:  threadInsert.AccountID,
		SubforumID: threadInsert.SubforumID,
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
		Where("id = ?", threadUpdate.ID).
		Updates(thread)
	if err := result.Error; err != nil {
		return err
	}
	if result.RowsAffected != 1 {
		return ErrNoRecordsAffected
	}
	return nil
}

func (store *ThreadStore) DeleteThread(ctx context.Context, threadID int) error {
	result := store.db.
		WithContext(ctx).
		Delete(&model.Thread{}, threadID)
	if err := result.Error; err != nil {
		return err
	}
	if result.RowsAffected != 1 {
		return ErrNoRecordsAffected
	}
	return nil
}
