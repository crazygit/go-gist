package main

import (
	"crypto/md5"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func sayHelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println(r.Form)
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key: ", k)
		fmt.Println("val: ", strings.Join(v, ""))
	}
	fmt.Fprint(w, "Hello astaxie!")
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("request method:", r.Method)
	if r.Method == "GET" {
		curtime := time.Now().Unix()
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(curtime, 10))
		token := fmt.Sprintf("%x", h.Sum(nil))
		t, _ := template.ParseFiles("login.gtpl")
		t.Execute(w, token)

	} else {
		fmt.Println("username: ", r.FormValue("username"))
		fmt.Println("password: ", r.FormValue("password"))
		//		fmt.Fprintln(w, r.FormValue("username"))
		template.HTMLEscape(w, []byte(r.FormValue("username")))
	}

}

func main() {
	http.HandleFunc("/", sayHelloName)
	http.HandleFunc("/login", login)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAnserve: ", err)
	}
}
