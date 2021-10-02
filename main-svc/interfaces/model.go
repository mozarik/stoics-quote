package interfaces

import "time"

type UserGorm struct {
	ID        uint      `gorm:"id"`
	Name      string    `gorm:"name"`
	Username  string    `gorm:"username"`
	Password  string    `gorm:"password"`
	CreatedAt time.Time `gorm:"created_at"`
	UpdatedAt time.Time `gorm:"updated_at"`
}

func (UserGorm) TableName() string {
	return "users"
}
