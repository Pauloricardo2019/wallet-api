package user

import (
	"errors"
	"time"
	"wallet-api/adapter/database/sql"
	"wallet-api/internal/model"

	"gorm.io/gorm"
)

type ListAllFn func(pagination *model.Pagination) ([]model.User, error)

func ListAll(pagination *model.Pagination) ([]model.User, error) {
	db, err := sql.GetGormDB()
	if err != nil {
		return nil, err
	}

	offset := (pagination.Page - 1) * pagination.Limit

	queryBuider := db.Limit(pagination.Limit).Offset(offset).Order(pagination.Sort)

	users := make([]model.User, 0)

	if err = queryBuider.Find(&users).
		Error; err != nil {
		return nil, err
	}

	return users, nil
}

type GetByIDFn func(ID uint64) (bool, *model.User, error)

func GetByID(ID uint64) (bool, *model.User, error) {
	db, err := sql.GetGormDB()
	if err != nil {
		return false, nil, err
	}

	result := &model.User{}
	if err = db.First(&result, "id = ?", ID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, result, nil
		}
		return false, nil, err
	}

	return true, result, nil
}

type GetByLoginFn func(login string) (bool, *model.User, error)

func GetByLogin(username string) (bool, *model.User, error) {
	db, err := sql.GetGormDB()
	if err != nil {
		return false, nil, err
	}

	result := &model.User{}
	if err = db.First(&result, "username = ?", username).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, result, nil
		}
		return false, nil, err
	}

	return true, result, err

}

type GetByEmailFn func(email string) (bool, *model.User, error)

func GetByEmail(email string) (found bool, user *model.User, err error) {
	db, err := sql.GetGormDB()
	if err != nil {
		return false, nil, err
	}

	result := &model.User{}
	if err = db.First(&result, "email = ?", email).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, result, nil
		}
		return false, nil, err
	}
	return true, result, err
}

type CreateFn func(user *model.User) (*model.User, error)

func Create(user *model.User) (*model.User, error) {
	db, err := sql.GetGormDB()
	if err != nil {
		return nil, err
	}

	user.CreatedAt = time.Now()

	if err = db.Create(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

type UpdateFn func(user *model.User) error

func Update(user *model.User) error {
	db, err := sql.GetGormDB()
	if err != nil {
		return err
	}

	if err = db.Model(&user).
		Select("*").
		Omit("HashedPassword, Password, Username").
		Save(&user).
		Error; err != nil {
		return err
	}

	return nil
}

type DeleteFn func(ID uint64) error

func Delete(ID uint64) error {
	db, err := sql.GetGormDB()
	if err != nil {
		return err
	}

	user := &model.User{}
	if err = db.Where("ID = ?", ID).Delete(user).Error; err != nil {
		return err
	}

	return nil
}
