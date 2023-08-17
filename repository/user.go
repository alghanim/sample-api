package repository

import (
	"alghanim/mediacmsAPI/model"
	"database/sql"
)

type UserRepository interface {
	Get(id int) (*model.User, error)
	Update(id int, user *model.User) (*model.User, error)
	// TODO: Add other operations like Create, Update, Delete
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Get(id int) (*model.User, error) {
	user := &model.User{}
	err := r.db.QueryRow("SELECT id, name, email FROM users WHERE id = ?", id).Scan(&user.ID, &user.Name, &user.Email)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *userRepository) Update(id int, user *model.User) (*model.User, error) {

	_, err := r.db.Exec("udpate users set name=?, email? where id = ?", user.Name, user.Email, user.ID)
	if err != nil {
		return nil, err
	}
	return r.Get(id)
}

// TODO: Implement other operations like Create, Update, Delete
