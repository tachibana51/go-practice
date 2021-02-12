package user
import(
	"2fa/domain/model"
	"time"
)
type User struct {
	ID model.UserID
	Email string
	MFAType string
	CreatedAt time.Time
	UpdatedAt time.Time
}
