package repository

import (
	"context"
)

func (r *Repository) Delete(ctx context.Context, id int, userID int) error {
	// Güvenlik için user_id kontrolü şart!
	query := "DELETE FROM overtimes WHERE id = $1 AND user_id = $2"

	result, err := r.DB.ExecContext(ctx, query, id, userID)
	if err != nil {
		return err
	}

	// İsteğe bağlı: Silinen satır sayısını kontrol edebiliriz
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		// Kayıt bulunamadıysa veya bu kullanıcıya ait değilse hata dönmüyoruz ama
		// istersen burada özel bir hata döndürebilirsin.
	}

	return nil
}
