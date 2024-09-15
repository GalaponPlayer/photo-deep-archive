package entity

type CommonDBAttributes struct {
	CreatedAt  string `json:"created_at"`
	ModifiedAt string `json:"modified_at"`
	DeletedAt  string `json:"deleted_at"`
}
