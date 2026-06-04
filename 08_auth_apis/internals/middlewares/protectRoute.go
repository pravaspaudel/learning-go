package middlewares

import (
	"context"
	"fmt"
	"net/http"

	"github.com/pravaspaudel/10_Auth_Apis/internals/ctx"
	"github.com/pravaspaudel/10_Auth_Apis/internals/utils"
)

func ProtectMiddleware(next http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("token")

		if err != nil {
			utils.ErrorResponse(w, http.StatusUnauthorized, "please authorize to continue")
		}

		fmt.Println(cookie.Value)

		claims, err := utils.VerifyJWT(cookie.Value)

		if err != nil {
			utils.ErrorResponse(w, http.StatusUnauthorized, "invalid token")
			return
		}

		userId, ok := claims["user_id"].(string)
		if !ok {
			utils.ErrorResponse(w, http.StatusUnauthorized, "invalid token claims")
			return
		}

		contxt := context.WithValue(r.Context(), ctx.UserKey, userId)

		next(w, r.WithContext(contxt))
	}
}
