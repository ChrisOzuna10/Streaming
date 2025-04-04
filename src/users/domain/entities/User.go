package entities

type User struct {
	ID       uint   `gorm:"primaryKey;column:id_user"`
	Name     string `gorm:"column:name_user;type:varchar(100);not null"`
	Email    string `gorm:"column:email_user;type:varchar(100);unique;not null"`
	Password string `gorm:"column:password_user;type:varchar(255);not null"`
}
