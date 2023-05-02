package auth

import (
	. "../../services/auth"
	"net/http"
)

func RequireTokenAuthentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/login" {
			cookie, err := r.Cookie("token")
			if err != nil || !IsAuth(cookie) {
				// Если куки не найдены, возвращаем ошибку
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				http.Redirect(w, r, "/login", http.StatusUnauthorized)
				return
			}
			next.ServeHTTP(w, r)
		}
	})
}
