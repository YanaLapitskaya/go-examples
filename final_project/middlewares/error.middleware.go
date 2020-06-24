package middlewares

import (
	"log"
	"net/http"
)

func ErrorHandler(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				error := err.(error).Error()
				log.Println("Recovered in ErrorHandler: ", error)
				http.Error(w, error, http.StatusBadRequest)
				return
			}
		}()
		next.ServeHTTP(w, r)
	})
}
