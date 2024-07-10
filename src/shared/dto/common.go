package dto

import (
	authConstants "app/src/shared/auth/constants"
	"app/src/shared/jwt"
	"time"
)

type CurrentUser struct {
	FullName string             `json:"fullName,omitempty"`
	Email    string             `json:"email"`
	Id       string             `json:"id"`
	Role     authConstants.Role `json:"role"`
}

func NewCurrentUser(payload *jwt.JwtPayload) *CurrentUser {

	role, _ := authConstants.RoleFromString(payload.Role)

	if payload != nil {
		return &CurrentUser{
			FullName: payload.FullName,
			Email:    payload.Email,
			Id:       payload.UserId,
			Role:     role,
		}
	}
	return nil
}

type AuditableResDto struct {
	Id        string    `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt time.Time `json:"deletedAt"`
}
