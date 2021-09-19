package contract

import (
	"backend-a-antar-jemput/internal/entities"
)

type LoginRequest struct {
	Username string
	Password string
}

type LoginResponse struct {
	ID        uint
	Name      string
	Role      string
	SID       string
	ExpiredAt int64
	IssuedAt  int64
}

func (l *LoginResponse) FromEntity(ent *entities.Login, sess *entities.Session) {
	l.ID = ent.ID
	l.Name = ent.Username
	l.SID = sess.SID
	l.ExpiredAt = sess.ExpiredAt.Unix()
	l.IssuedAt = sess.CreatedAt.Unix()
	switch ent.LoginAs {
	case 1:
		l.Role = "Customer"
	case 2:
		l.Role = "Agen"
	}
}
