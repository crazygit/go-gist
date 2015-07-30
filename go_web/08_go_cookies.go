package main

import (
	"fmt"
	"net/http"
	"time"
)

func setCookie(w http.ResponseWriter, r *http.Request) {
	expiration := time.Now()
	expiration = expiration.AddDate(0, 0, 1)
	cookie := http.Cookie{Name: "username", Value: "crazygit", Expires: expiration}
	http.SetCookie(w, &cookie)
	fmt.Fprint(w, "set cookie!")
}

func getCookie(w http.ResponseWriter, r *http.Request) {
	for _, cookie := range r.Cookies() {
		fmt.Fprintf(w, "%v\n", cookie)
	}
}

func main() {
	http.HandleFunc("/set", setCookie)
	http.HandleFunc("/get", getCookie)
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		fmt.Println(err)
	}
}
