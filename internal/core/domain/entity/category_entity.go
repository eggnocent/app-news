package entity

type CategoryEntity struct {
	ID    int64
	Title string
	Slug  string
	User  UserEntity
}
