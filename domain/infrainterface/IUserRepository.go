package infrainterface
import (
	"2fa/domain/model"
	"2fa/domain/model/user"
)
type IUserRepository interface {
	ActivateAndCreate(user user.User) error
	FindById(id model.UserID) (user.User, error)
}