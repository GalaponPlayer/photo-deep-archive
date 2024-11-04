package entity

type UserID uint

func (id UserID) Value() uint {
	return uint(id)
}

type User struct {
	ID   UserID `json:"id" gorm:"primaryKey;default:auto_random()"`
	Name string `json:"name"`
	CommonDBAttributes
}

func NewUser(id UserID, name string, ts int64) *User {
	return &User{
		ID:   id,
		Name: name,
		CommonDBAttributes: CommonDBAttributes{
			CreatedAt:  ts,
			ModifiedAt: ts,
			DeletedAt:  0,
		},
	}
}
