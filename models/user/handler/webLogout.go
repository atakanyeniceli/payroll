package handler

import (
	"net/http"
	"time"
)

// WebLogout, web kullanıcıları için güvenli çıkış işlemini yönetir.
func (h *Handler) WebLogout(w http.ResponseWriter, r *http.Request) {
	// 1. Tarayıcıdan cookie'yi oku
	cookie, err := r.Cookie("session_token")
	if err == nil {
		// 2. Token varsa, sunucu belleğinden sil (Session'ı öldür)
		h.TokenManager.RevokeSession(cookie.Value)
	}

	// 3. Tarayıcıdaki cookie'yi geçersiz kıl (Sil)
	// Aynı isimde, süresi geçmiş ve boş değerli bir cookie göndererek tarayıcının silmesini sağlarız.
	http.SetCookie(w, &http.Cookie{
		Name:     "session_token",
		Value:    "",
		Path:     "/",
		Expires:  time.Unix(0, 0),
		MaxAge:   -1,
		HttpOnly: true,
	})

	// 4. Giriş sayfasına yönlendir
	http.Redirect(w, r, "/", http.StatusFound)
}
