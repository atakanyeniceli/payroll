package handler

import (
	"net/http"
	"time"
)

func (h *Handler) WebLogin(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	password := r.FormValue("password")

	user, err := h.Service.Login(email, password)

	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		// Güvenlik: İç sistem hatalarını kullanıcıya gösterme.
		w.Write([]byte("E-posta veya şifre hatalı."))
		return
	}

	// Giriş başarılı, enjekte edilen token yöneticisini kullanarak oturum oluştur.
	sessionToken, err := h.TokenManager.CreateSession(user.ID, 24*time.Hour)
	if err != nil {
		// Gerçek bir uygulamada bu hatayı loglamalısınız.
		http.Error(w, "Oturum oluşturulurken bir sunucu hatası oluştu.", http.StatusInternalServerError)
		return
	}

	// Oluşturulan token'ı istemciye cookie olarak gönder.
	http.SetCookie(w, &http.Cookie{
		Name:     "session_token",
		Value:    sessionToken,
		Path:     "/",
		Expires:  time.Now().Add(24 * time.Hour),
		HttpOnly: true,  // JavaScript'in cookie'ye erişimini engeller (XSS koruması)
		Secure:   false, // Üretimde true olmalı (sadece HTTPS üzerinden gönderilir)
		SameSite: http.SameSiteLaxMode,
	})
	w.Header().Set("HX-Redirect", "/dashboard")
}
