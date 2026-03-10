package store

import (
	"context"
	"forum/model"

	"gorm.io/gorm"
)

type AccountStore struct {
	db *gorm.DB
}

func NewAccountStore(db *gorm.DB) *AccountStore {
	return &AccountStore{db}
}

func (store *AccountStore) SelectAccount(ctx context.Context, accountId int) (model.Account, error) {
	var account model.Account
	err := store.db.
		WithContext(ctx).
		Preload("Threads.Account").
		Preload("Replies.Account").
		Preload("Threads").
		Preload("Replies").
		First(&account, accountId).Error
	if err != nil {
		return account, err
	}
	return account, nil
}

func (store *AccountStore) SelectAccountCredentials(ctx context.Context, username, email string) (model.AccountCredentials, error) {
	var accountCredentials model.AccountCredentials
	err := store.db.
		WithContext(ctx).
		Where("username = ? OR email = ?", username, email).
		Select("id, username, email, password_hash, is_admin").
		First(&accountCredentials).Error
	if err != nil {
		return accountCredentials, err
	}
	return accountCredentials, nil
}

func (store *AccountStore) SelectAccounts(ctx context.Context) ([]model.Account, error) {
	var accounts []model.Account
	err := store.db.
		WithContext(ctx).
		Find(&accounts).Error
	if err != nil {
		return nil, err
	}
	return accounts, nil
}

func (store *AccountStore) InsertAccount(ctx context.Context, accountInsert *model.AccountInsert) error {
	account := &model.Account{
		Username:     accountInsert.Username,
		Email:        accountInsert.Email,
		PasswordHash: accountInsert.PasswordHash,
		ProfilePath:  accountInsert.ProfilePath,
	}
	err := store.db.
		WithContext(ctx).
		Create(&account).Error
	if err != nil {
		return err
	}
	return nil
}

func (store *AccountStore) UpdateAccountCredentials(ctx context.Context, accountCreds *model.AccountUpdateCredentials) error {
	account := &model.Account{
		Username: accountCreds.Username,
		Email:    accountCreds.Email,
	}
	result := store.db.WithContext(ctx).
		Model(&model.Account{}).
		Where("id = ?", accountCreds.Id).
		Updates(&account)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return ErrNoRecordsAffected
	}
	return nil
}

func (store *AccountStore) DeleteAccount(ctx context.Context, accountId int) error {
	result := store.db.
		WithContext(ctx).
		Delete(&model.Account{}, accountId)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return ErrNoRecordsAffected
	}
	return nil
}
