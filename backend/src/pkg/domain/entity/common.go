package entity

type CommonDBAttributes struct {
	CreatedAt  int64 `json:"created_at"`
	ModifiedAt int64 `json:"modified_at"`
	DeletedAt  int64 `json:"deleted_at"`
}
