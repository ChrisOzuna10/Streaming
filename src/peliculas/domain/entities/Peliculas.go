package entities

import userEntities "Streaming/src/users/domain/entities"

type Pelicula struct {
	ID          uint              `gorm:"primaryKey;column:id_pelicula"`
	Title       string            `gorm:"column:title_pelicula;type:varchar(100);not null" json:"title"`
	Description string            `gorm:"column:description_pelicula;type:text" json:"description"`
	URL         string            `gorm:"column:url_pelicula;type:varchar(500);not null" json:"url"`
	Genero      string            `gorm:"column:genero;type:varchar(50);not null" json:"genero"`
	UserID      uint              `gorm:"column:user_id" json:"user_id"`
	User        userEntities.User `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
}
