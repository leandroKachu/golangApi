package middleware

import (
	"api/src/auth"
	"api/src/errorsResponse"
	"fmt"
	"net/http"
)

// logger show in terminal the result from handlefunc
func Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("\n %s %s %s", r.Method, r.RequestURI, r.Host)
		next(w, r)
	}
}

// create a validation for request which one had a required validation
func Authenticator(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("middleware executing here")
		if err := auth.ValidToken(r); err != nil {
			errorsResponse.Error(w, http.StatusUnauthorized, err)
		}
		//if false err, keep going
		next(w, r)
	}
}
