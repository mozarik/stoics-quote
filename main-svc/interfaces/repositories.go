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

func NewUserRepo(db *gorm.DB) *UserRepo {
	return &UserRepo{DB: db}
}

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

func (repo UserRepo) Store(user domain.User) (*domain.User, error) {
	createUserGorm := UserGorm{
		Name:      user.Name,
		Username:  user.Username,
		Password:  user.Password,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	row := repo.DB.Debug().Create(&createUserGorm).Row()
	err := row.Scan(&createUserGorm.ID, &createUserGorm.Name, &createUserGorm.Username, &createUserGorm.Password, &createUserGorm.CreatedAt, &createUserGorm.UpdatedAt)
	if err != nil {
		return nil, err
	}

	user.ID = int(createUserGorm.ID)
	return &user, nil
}

func (repo UserRepo) FindByUsername(username string) (*domain.User, error) {
	query := `SELECT u.id, u.name, u.username FROM users u WHERE u.username = @username LIMIT 1`
	var user domain.User

	row := repo.DB.Debug().Raw(query, sql.Named("username", "johndoe")).Row()
	err := row.Scan(&user.ID, &user.Name, &user.Username)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (repo UserRepo) Delete(id int) error {
	query := `DELETE FROM users WHERE id = @id`
	err := repo.DB.Debug().Exec(query, sql.Named("id", id)).Error
	if err != nil {
		return err
	}
	return nil
}
