package interfaces

import (
	"database/sql"
	"main-svc/domain"
	"time"

	"gorm.io/gorm"
)

type DbRepo struct {
	DB *gorm.DB
}

func NewDbRepo(db *gorm.DB) *DbRepo {
	return &DbRepo{DB: db}
}

type UserRepo DbRepo

// Implement domain.UserRepository interface
func (repo UserRepo) FindByID(id int) (*domain.User, error) {
	query := `SELECT u.id, u.name, u.username FROM users u WHERE u.id = @userID`
	var user domain.User

	row := repo.DB.Debug().Raw(query, sql.Named("userID", id)).Row()
	err := row.Scan(&user.ID, &user.Name, &user.Username)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (repo UserRepo) Store(user domain.User) error {
	createUserGorm := UserGorm{
		Name:      user.Name,
		Username:  user.Username,
		Password:  user.Password,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err := repo.DB.Debug().Create(&createUserGorm).Error
	if err != nil {
		return err
	}
	return nil
}
