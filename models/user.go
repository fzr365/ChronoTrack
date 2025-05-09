package models

type User struct {
	ID       uint   `json:"id" gorm:"primaryKey"` // 显式定义 ID 字段
	Username string `json:"username"`
	Password string `json:"password"`
}
