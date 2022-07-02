package main

import (
	"net/http"
)

func Authmiddleware(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		username, password, ok := r.BasicAuth()

		if !ok || !CheckUsernameAndPassword(username, password) {
			w.Header().Set("WWW-Authenticate", `Basic realm="Please enter your username and password for this site"`)
			w.WriteHeader(401)
			w.Write([]byte("Unauthorised.\n"))
			return
		}
		handler(w, r)
	}
}

func CheckUsernameAndPassword(username, password string) bool {
	usrname := ConfigsAuth()[0]
	pass := ConfigsAuth()[1]
	return username == usrname && password == pass
}
