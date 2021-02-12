package infrainterface

import (
	"2fa/domain/model/authorization"
	"2fa/domain/model/user"
)

type IJWTManager interface {
	Generate(u user.User) (string, error)
	ParseAuthorizationString(s string) (authorization.Authorization, error)
	IsValid(s string) (bool)
}