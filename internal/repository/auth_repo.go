package repository

import (
	"backend-a-antar-jemput/internal/entities"

	"gorm.io/gorm"
)

type AuthRepository interface {
	Get(ent *entities.Login) error
	GetUserID(ent entities.Login) uint
	Create(ent *entities.Login) error
	Update(ent *entities.Login) error
	Delete(ent *entities.Login) error
}

type SessionRepository interface {
	Get(ent *entities.Session) error
	Create(ent *entities.Session) error
	Delete(ent *entities.Session) error
}

type AuthRepositoryMysql struct {
	Db *gorm.DB
}

func (a AuthRepositoryMysql) GetUserID(ent entities.Login) uint {
	var err *gorm.DB
	var id uint
	if ent.LoginAs == 1 {
		user := entities.Customers{Login: ent}
		err = a.Db.Where(&user).First(&user)
		id = user.ID

	} else {
		user := entities.Agents{Login: ent, LoginID: ent.ID}
		err = a.Db.Where(&user).First(&user)
		id = user.ID
	}
	if err.Error != nil {
		return 0
	}
	return id
}

func (a AuthRepositoryMysql) Get(ent *entities.Login) error {
	err := a.Db.Where(ent).First(&ent)
	if err.Error != nil {
		return err.Error
	}
	return err.Error
}

func (a AuthRepositoryMysql) Create(ent *entities.Login) error {
	err := a.Db.Create(&ent)
	if err.Error != nil {
		return err.Error
	}
	return err.Error
}

func (a AuthRepositoryMysql) Update(ent *entities.Login) error {
	err := a.Db.Save(&ent)
	if err.Error != nil {
		return err.Error
	}
	return err.Error
}

func (a AuthRepositoryMysql) Delete(ent *entities.Login) error {
	err := a.Db.Delete(&ent)
	if err.Error != nil {
		return err.Error
	}
	return err.Error
}

type SessionRepositoryMysql struct {
	Db *gorm.DB
}

func (a SessionRepositoryMysql) Get(ent *entities.Session) error {
	err := a.Db.Where(ent).First(&ent)
	if err.Error != nil {
		return err.Error
	}
	return err.Error
}

func (a SessionRepositoryMysql) Create(ent *entities.Session) error {
	err := a.Db.Create(&ent)
	if err.Error != nil {
		return err.Error
	}
	return err.Error
}

func (a SessionRepositoryMysql) Delete(ent *entities.Session) error {
	err := a.Db.Where(ent).Delete(&ent)
	if err.Error != nil {
		return err.Error
	}
	return err.Error
}
