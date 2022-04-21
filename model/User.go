package user

import (
	"errors"

	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Nama     string
	Email    string
	Password string
}

type UserModel struct {
	Db *gorm.DB
}

func New(db *gorm.DB) *UserModel {
	return &UserModel{
		Db: db,
	}
}

func (ur *UserModel) Insert(newUser User) (User, error) {
	if err := ur.Db.Create(&newUser).Error; err != nil {
		log.Warn(err)
		return User{}, errors.New("tidak bisa insert data")
	}
	log.Info()
	return newUser, nil
}

func (ur *UserModel) GetAll() ([]User, error) {
	arrUser := []User{}

	if err := ur.Db.Find(&arrUser).Error; err != nil {
		log.Warn(err)
		return nil, errors.New("tidak bisa select data")
	}

	if len(arrUser) == 0 {
		log.Warn("tidak ada data")
		return nil, errors.New("tidak ada data")
	}

	log.Info()
	return arrUser, nil
}

func (ur *UserModel) GetUser() ([]User, error) {
	arrUser := []User{}
	id := ur.Db.Select("id")
	if err := ur.Db.Find(&arrUser, id).Error; err != nil {
		log.Warn(err)
		return nil, errors.New("tidak bisa select data")
	}

	if len(arrUser) == 0 {
		log.Warn("tidak ada data")
		return nil, errors.New("tidak ada data")
	}

	log.Info()
	return arrUser, nil
}
func (ur *UserModel) Update(updateUser User) (User, error) {
	id := ur.Db.Select("id")
	if err := ur.Db.Select(&updateUser, id).Error; err != nil {
		log.Warn(err)
		return User{}, errors.New("tidak bisa insert data")
	}
	log.Info()
	return updateUser, nil
}
