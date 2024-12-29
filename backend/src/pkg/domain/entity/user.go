package entity

type UserID uint

func (id UserID) Value() uint {
	return uint(id)
}

type User struct {
	ID    UserID `json:"id" gorm:"primaryKey;default:auto_random()"`
	Name  string `json:"name"`
	Email string `json:"email" gorm:"type:text;uniqueIndex:idx_email,length:191"`
	CommonDBAttributes
}

func NewUser(id UserID, name string, email string) *User {
	return &User{
		ID:                 id,
		Name:               name,
		Email:              email,
		CommonDBAttributes: CommonDBAttributes{},
	}
}
