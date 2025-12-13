package repository

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/atakanyeniceli/payroll/models/user"
)

func (r *Repository) GetUserByEmail(email string) (*user.User, error) {
	user := &user.User{}
	query := "SELECT id,firstname,lastname,password FROM users WHERE email = $1"

	err := r.DB.QueryRow(query, email).Scan(&user.ID, &user.Firstname, &user.Lastname, &user.Password)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			// Kullanıcı bulunamadı, bu bir hata değil, beklenen bir durum.
			// Servis katmanının bu durumu yönetebilmesi için sql.ErrNoRows'u döndürüyoruz.
			return nil, sql.ErrNoRows
		}
		// Diğer tüm veritabanı hataları için.
		return nil, fmt.Errorf("kullanıcı sorgulanırken veritabanı hatası: %w", err)
	}

	user.Email = email
	return user, nil
}
