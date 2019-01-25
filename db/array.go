package db

import (
	"errors"
	"github.com/papvan/hlcup/models"
)

const Version = "array"

const (
	MaxAccounts           = 1500000
	DefaultShardsCount    = 509
)

type DB struct {
	accounts     []models.Account

	lockA *ShardedLock
}

func New() *DB {

	db := new(DB)

	db.accounts = make([]models.Account, MaxAccounts)

	db.lockA = NewShardedLock(DefaultShardsCount)

	return db
}

// TODO: Возможно это нах не надо, проверить позже
var ErrAlreadyExists = errors.New("already exists")

func (db *DB) GetAccount(id uint32) models.Account {
	if id >= MaxAccounts {
		return models.Account{}
	}
	db.lockA.RLock(id)
	defer db.lockA.RUnlock(id)
	return db.accounts[id]
}

func (db *DB) AddUser(v models.Account) error {
	if err := v.Validate(); err != nil {
		return err
	}
	db.lockA.Lock(v.Id)
	if db.accounts[v.Id].IsValid() {
		return ErrAlreadyExists
	}
	db.accounts[v.Id] = v
	db.lockA.Unlock(v.Id)
	return nil
}