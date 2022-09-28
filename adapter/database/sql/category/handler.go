package category

import (
	"errors"
	"gorm.io/gorm"
	"wallet-api/adapter/database/sql"
	"wallet-api/internal/model"
)

type GetByIDFn func(ID uint64) (bool, *model.Category, error)

func GetByID(ID uint64) (bool, *model.Category, error) {
	db, err := sql.GetGormDB()
	if err != nil {
		return false, nil, err
	}

	result := &model.Category{}

	if err = db.First(&result, "id = ?", ID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, result, nil
		}
		return false, nil, err
	}

	return true, result, nil
}

type CreateFn func(category *model.Category) (*model.Category, error)

func Create(category *model.Category) (*model.Category, error) {
	db, err := sql.GetGormDB()
	if err != nil {
		return nil, err
	}
	
	if err = db.Create(category).Error; err != nil {
		return nil, err
	}

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
