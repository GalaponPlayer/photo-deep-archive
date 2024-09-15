package entity

type UserID string
type User struct {
	ID   UserID `json:"id"`
	Name string `json:"name"`
	CommonDBAttributes
}

func NewUser(id UserID, name string) *User {
	return &User{
		ID:   id,
		Name: name,
	}
}
