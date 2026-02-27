package repositories

import (
	"api_golang/models"
	"database/sql"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(u *models.User) (*models.User, error) {
	q := `INSERT INTO users (id, name, age, phone, email) VALUES (?, ?, ?, ?, ?)`
	_, err := r.db.Exec(q, u.ID, u.Name, u.Age, u.Phone, u.Email)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (r *UserRepository) GetByID(id string) (*models.User, bool) {
	q := `SELECT id, name, age, phone, email FROM users WHERE id = ?`
	var u models.User
	err := r.db.QueryRow(q, id).Scan(&u.ID, &u.Name, &u.Age, &u.Phone, &u.Email)
	if err != nil {
		return nil, false
	}
	return &u, true
}

func (r *UserRepository) GetByEmail(email string) (*models.User, bool) {
	q := `SELECT id, name, age, phone, email FROM users WHERE email = ?`
	var u models.User
	err := r.db.QueryRow(q, email).Scan(&u.ID, &u.Name, &u.Age, &u.Phone, &u.Email)
	if err != nil {
		return nil, false
	}
	return &u, true
}

func (r *UserRepository) List() []*models.User {
	q := `SELECT id, name, age, phone, email FROM users`
	rows, err := r.db.Query(q)
	if err != nil {
		return nil
	}
	defer rows.Close()
	var list []*models.User
	for rows.Next() {
		var u models.User
		if err := rows.Scan(&u.ID, &u.Name, &u.Age, &u.Phone, &u.Email); err != nil {
			continue
		}
		list = append(list, &u)
	}
	return list
}

func (r *UserRepository) Update(u *models.User) (*models.User, bool) {
	q := `UPDATE users SET name = ?, age = ?, phone = ?, email = ? WHERE id = ?`
	res, err := r.db.Exec(q, u.Name, u.Age, u.Phone, u.Email, u.ID)
	if err != nil {
		return nil, false
	}
	n, _ := res.RowsAffected()
	if n == 0 {
		return nil, false
	}
	return u, true
}

func (r *UserRepository) Delete(id string) bool {
	q := `DELETE FROM users WHERE id = ?`
	res, err := r.db.Exec(q, id)
	if err != nil {
		return false
	}
	n, _ := res.RowsAffected()
	return n > 0
}
