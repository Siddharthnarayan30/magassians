package middleware

import (
	"net/http"

	"org.Magassians/util"
)

func Recovery(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		defer func() {
			if rvr := recover(); rvr != nil && rvr != http.ErrAbortHandler {
				util.Log.Print(rvr)

				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusInternalServerError)
			}

		}()

		next.ServeHTTP(w, r)

	})
}
