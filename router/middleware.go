package router

import (
	"context"
	"net/http"

	"github.com/atakanyeniceli/payroll/tools/token"
)

// userContextKey, context içinde kullanıcı ID'sini saklamak için özel bir türdür.
// Bu, farklı paketlerdeki context anahtarlarının çakışmasını önler.
type userContextKey string

const UserIDKey userContextKey = "userID"

// WebAuthMiddleware, web arayüzü için gelen isteklerde session token'ını (cookie) kontrol eden bir ara katmandır.
func WebAuthMiddleware(tokenManager *token.Manager) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// 1. Tarayıcıdan cookie'yi oku
			cookie, err := r.Cookie("session_token")
			if err != nil {
				// Cookie yoksa, kullanıcı giriş yapmamıştır. Login sayfasına yönlendir.
				// Eğer istek HTMX ile yapılmışsa, istemci tarafında tam sayfa yönlendirmesi yap.
				if r.Header.Get("HX-Request") == "true" {
					w.Header().Set("HX-Redirect", "/")
					w.WriteHeader(http.StatusOK)
					return
				}
				http.Redirect(w, r, "/", http.StatusFound)
				return
			}

			// 2. Cookie'deki token değerini al
			tokenValue := cookie.Value

			// 3. Token'ın geçerli olup olmadığını kontrol et
			sessionData, ok := tokenManager.GetSessionData(tokenValue)
			if !ok {
				// Token geçerli değilse (süresi dolmuş veya sahte), login sayfasına yönlendir.
				// Güvenlik için geçersiz cookie'yi de silebiliriz.
				http.SetCookie(w, &http.Cookie{Name: "session_token", MaxAge: -1})

				if r.Header.Get("HX-Request") == "true" {
					w.Header().Set("HX-Redirect", "/")
					w.WriteHeader(http.StatusOK)
					return
				}
				http.Redirect(w, r, "/", http.StatusFound)
				return
			}

			// 4. Token geçerliyse, kullanıcı ID'sini request context'ine ekle.
			// Bu sayede sonraki handler bu bilgiye erişebilir.
			ctx := context.WithValue(r.Context(), UserIDKey, sessionData.UserID)

			// 5. İsteği, güncellenmiş context ile bir sonraki handler'a ilet.
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
