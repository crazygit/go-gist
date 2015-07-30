package main

import (
	"fmt"
	"github.com/drone/routes"
	"net/http"
)

func getuser(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	uid := params.Get(":uid")
	fmt.Fprintf(w, "your are get user %s", uid)
}

func main() {
	mux := routes.New()
	mux.Get("/user/:uid", getuser)
	http.Handle("/", mux)
	http.ListenAndServe(":8080", nil)
}
