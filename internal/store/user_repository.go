package store

import "github.com/soldiii/supervisor_app/internal/models"

type UserRepository struct {
	store *Store
}

func (r *UserRepository) Create(u *models.User) (*models.User, error) {
	if err := r.store.db.QueryRow(
		"INSERT INTO users (email, name, surname, patronymic, reg_date_time, password) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id",
		u.Email,
		u.Name,
		u.Surname,
		u.Patronymic,
		u.Reg_date_time,
		u.Password).Scan(&u.Id); err != nil {
		return nil, err
	}
	return u, nil
}

func (r *UserRepository) GetByEmail(email string) (*models.User, error) {
	u := &models.User{}
	if err := r.store.db.QueryRow("SELECT id, email, name, surname, patronymic, reg_date_time, password FROM users WHERE email = $1",
		email).Scan(&u.Id,
		&u.Email,
		&u.Name,
		&u.Surname,
		&u.Patronymic,
		&u.Reg_date_time,
		&u.Password); err != nil {
		return nil, err
	}
	return u, nil
}
