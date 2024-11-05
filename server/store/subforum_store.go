package store

import (
	"context"
	"forum/model"

	"gorm.io/gorm"
)

type SubforumStore struct {
	db *gorm.DB
}

func NewSubforumStore(db *gorm.DB) *SubforumStore {
	return &SubforumStore{db}
}

func (store *SubforumStore) SelectSubforum(ctx context.Context, subforumID int) (model.Subforum, error) {
	var subforum model.Subforum
	err := store.db.
		WithContext(ctx).
		Where("id = ?", subforumID).
		First(&subforum).Error
	if err != nil {
		return subforum, err
	}
	return subforum, nil
}

func (store *SubforumStore) SelectSubforums(ctx context.Context) ([]model.Subforum, error) {
	var subforums []model.Subforum
	latestThreads := store.db.
		Preload("Account").
		Table("thread").
		Select("id").
		Where("created_at = (SELECT MAX(created_at) FROM thread t WHERE t.subforum_id = thread.subforum_id)")
	err := store.db.
		Preload("Threads", func(db *gorm.DB) *gorm.DB {
			return db.Where("id IN (?)", latestThreads).Preload("Account")
		}).
		Find(&subforums).Error
	if err != nil {
		return nil, err
	}
	return subforums, nil
}

func (store *SubforumStore) InsertSubforum(ctx context.Context, subforumInsert *model.SubforumInsert) error {
	subforum := model.Subforum{
		Title:       subforumInsert.Title,
		Description: subforumInsert.Description,
	}
	err := store.db.
		WithContext(ctx).
		Create(&subforum).Error
	if err != nil {
		return err
	}
	return nil
}

func (store *SubforumStore) UpdateSubforum(ctx context.Context, subforumUpdate *model.SubforumUpdate) error {
	subforum := model.Subforum{
		Title:       subforumUpdate.Title,
		Description: subforumUpdate.Description,
	}
	result := store.db.
		WithContext(ctx).
		Model(&subforum).
		Where("id = ?", subforumUpdate.ID).
		Updates(&subforum)
	if err := result.Error; err != nil {
		return err
	}
	if result.RowsAffected != 1 {
		return ErrNoRecordsAffected
	}
	return nil
}

func (store *SubforumStore) DeleteSubforum(ctx context.Context, subforumID int) error {
	result := store.db.
		WithContext(ctx).
		Where("id = ?", subforumID).
		Delete(&model.Subforum{})
	if err := result.Error; err != nil {
		return err
	}
	if result.RowsAffected != 1 {
		return ErrNoRecordsAffected
	}
	return nil
}
