package entity

type User struct {
	Id        uint   `json:"id"`
	Name      string `json:"name"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	IsDeleted uint   `json:"is_deleted"`
}
