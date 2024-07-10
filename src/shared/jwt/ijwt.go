package jwt

import (
	"app/src/apis/user/models"
)

type Manager interface {
	CreateToken(user *user.User) (string, *JwtPayload, error)

	VerifyToken(token string) (*JwtPayload, error)
}
