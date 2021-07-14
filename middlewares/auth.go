package middlewares

/**
 * VERY VERY BAD AUTH don't ever use this in production
 * considering this server can run arbitrary commands
 */
import (
	"context"
	"net/http"
	"strings"
)

type auth struct {
	secret string
}

const AuthKey MWContext = "auth"

func Auth(secret string) Middleware {
	return func(h http.Handler) http.Handler {

		a := auth{secret: secret}
		return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
			auth := r.Header.Get("Auth")
			auth = strings.TrimPrefix(auth, "Bearer ")

			if a.secret != auth {
				rw.WriteHeader(401)

				rw.Write([]byte(`{ "message": "Not Authorized" }`))
				return
			}

			ctx := context.WithValue(r.Context(), AuthKey, true)

			h.ServeHTTP(rw, r.WithContext(ctx))
		})
	}
}
