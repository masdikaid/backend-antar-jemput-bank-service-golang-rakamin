package repository

import (
	"backend-a-antar-jemput/internal/databases"
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
}

func (AuthRepositoryMysql) GetUserID(ent entities.Login) uint {
	var err *gorm.DB
	var id uint
	if ent.LoginAs == 1 {
		user := entities.Customers{Login: ent}
		err = databases.DBCon.Where(&user).First(&user)
		id = user.ID

	} else {
		user := entities.Agents{Login: ent, LoginID: ent.ID}
		err = databases.DBCon.Where(&user).First(&user)
		id = user.ID
	}
	if err.Error != nil {
		return 0
	}
	return id
}

func (AuthRepositoryMysql) Get(ent *entities.Login) error {
	err := databases.DBCon.Where(ent).First(&ent)
	if err.Error != nil {
		return err.Error
	}
	return err.Error
}

func (AuthRepositoryMysql) Create(ent *entities.Login) error {
	err := databases.DBCon.Create(&ent)
	if err.Error != nil {
		return err.Error
	}
	return err.Error
}

func (AuthRepositoryMysql) Update(ent *entities.Login) error {
	err := databases.DBCon.Save(&ent)
	if err.Error != nil {
		return err.Error
	}
	return err.Error
}

func (AuthRepositoryMysql) Delete(ent *entities.Login) error {
	err := databases.DBCon.Delete(&ent)
	if err.Error != nil {
		return err.Error
	}
	return err.Error
}

type SessionRepositoryMysql struct {
}

func (SessionRepositoryMysql) Get(ent *entities.Session) error {
	err := databases.DBCon.Where(ent).First(&ent)
	if err.Error != nil {
		return err.Error
	}
	return err.Error
}

func (SessionRepositoryMysql) Create(ent *entities.Session) error {
	err := databases.DBCon.Create(&ent)
	if err.Error != nil {
		return err.Error
	}
	return err.Error
}

func (SessionRepositoryMysql) Delete(ent *entities.Session) error {
	err := databases.DBCon.Where(ent).Delete(&ent)
	if err.Error != nil {
		return err.Error
	}
	return err.Error
}
