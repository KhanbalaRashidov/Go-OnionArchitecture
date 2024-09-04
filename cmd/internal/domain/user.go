package domain

type User struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Username string `json:"username" gorm:"unique"`
	Password string `json:"password"`
	Email    string `json:"email" gorm:"unique"`
}
