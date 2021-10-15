package service

import (
	"backend-a-antar-jemput/internal/contract"
	"backend-a-antar-jemput/internal/entities"
	"backend-a-antar-jemput/internal/repository"
	"backend-a-antar-jemput/tools/helper"
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type AuthService struct {
	AuthRepository    repository.AuthRepository
	SessionRepository repository.SessionRepository
}

func (auth AuthService) Login(contr contract.LoginRequest) (string, error) {
	ent := entities.Login{Username: contr.Username}
	err := auth.AuthRepository.Get(&ent)
	if err != nil {
		return err.Error(), err
	}

	if ent.Password != contr.Password {
		err = errors.New("wrong password")
		return err.Error(), err
	}
	fmt.Println(os.Getenv("EXPIRED-SESSION"))
	expDay, err := strconv.ParseInt(os.Getenv("EXPIRED-SESSION"), 32, 10)
	if err != nil {
		panic(err)
		// return err.Error(), err
	}

	session := entities.Session{Login: ent, SID: helper.GenerateSessionID(), ExpiredAt: time.Now().AddDate(0, 0, int(expDay))}
	err = auth.SessionRepository.Create(&session)
	if err != nil {
		return err.Error(), err
	}

	LoginContract := contract.LoginResponse{}
	LoginContract.FromEntity(&ent, &session)

	userid := auth.AuthRepository.GetUserID(ent)

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = userid
	claims["nama"] = LoginContract.Name
	claims["role"] = LoginContract.Role
	claims["sid"] = LoginContract.SID
	claims["exp"] = LoginContract.ExpiredAt
	claims["iat"] = LoginContract.IssuedAt
	t, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return err.Error(), err
	}

	return t, nil

}

func (auth AuthService) ValidateSession(sid string) error {
	ent := entities.Session{SID: sid}
	err := auth.SessionRepository.Get(&ent)
	if err != nil {
		return err
	}
	return nil
}

func (auth AuthService) Logout(contr *contract.LoginResponse) error {
	ent := entities.Session{SID: contr.SID}
	err := auth.SessionRepository.Get(&ent)
	if err != nil {
		return err
	}
	err = auth.SessionRepository.Delete(&ent)
	if err != nil {
		return err
	}
	return nil
}
