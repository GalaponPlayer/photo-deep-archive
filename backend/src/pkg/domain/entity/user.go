package entity

import "app/src/pkg/lib"

type UserID lib.UUIDv4

type User struct {
	ID   UserID `json:"id"`
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
