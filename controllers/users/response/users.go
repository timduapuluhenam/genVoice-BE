package response

import (
	"genVoice/business/users"
	"time"
)

type UserRegisterResponse struct {
	Message   string    `json:"message"`
	ID        int       `json:"id"`
	Username  string    `json:"username"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func FromDomainRegister(domain users.Domain) UserRegisterResponse {
	return UserRegisterResponse{
		Message:   "Registration Success",
		ID:        domain.ID,
		Username:  domain.Username,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}

type UserLoginResponse struct {
	Message  string `json:"message"`
	Username string `json:"username"`
	ID       int    `json:"id"`
	Token    string `json:"token"`
}

func FromDomainLogin(domain users.Domain) UserLoginResponse {
	return UserLoginResponse{
		Message:  "Login Success",
		Username: domain.Username,
		ID:       domain.ID,
		Token:    domain.Token,
	}
}
