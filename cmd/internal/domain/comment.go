package domain

type Comment struct {
	ID     uint   `json:"id" gorm:"primaryKey"`
	PostID uint   `json:"post_id"`
	UserID uint   `json:"user_id"`
	Text   string `json:"text"`
}
