package repository

import (
	"fmt"

	"github.com/atakanyeniceli/payroll/models/user"
)

func (r *Repository) CreateUser(user *user.User) error {

	var newID int

	query := `INSERT INTO users (firstname, lastname, email, password) 
			  VALUES ($1, $2, $3,$4) RETURNING id`

	err := r.DB.QueryRow(query, user.Firstname, user.Lastname, user.Email, user.Password).Scan(&newID)
	if err != nil {
		// Örneğin benzersiz (unique) email hatası burada yakalanabilir.
		return fmt.Errorf("kayıt oluşturulurken DB hatası: %w", err)
	}
	return nil
}
