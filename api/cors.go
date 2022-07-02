package main

import (
	"net/http"
)

func (a *App) Cors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		o := r.Header.Get("Origin")
		if o == "" {
			o = "*"
		}

		h := w.Header()
		h.Set("Access-Control-Allow-Origin", o)
		h.Set("Access-Control-Allow-Methods", "*")
		h.Set("Access-Control-Allow-Headers", "*")

		if r.Method == "OPTIONS" {
			w.Write([]byte("OK"))
			return
		}
		next.ServeHTTP(w, r)
	})
}
