package model

// Storage is
type Storage struct {
	IDAutoModel
	UUIDModel
	NameModel
	TypeMdel
	URLModel
	UserID int64 `json:"user_id"`
	TimeAllModel
}
