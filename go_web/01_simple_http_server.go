package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

func sayHelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println(r.Form)
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_log"])
	for k, v := range r.Form {
		fmt.Println("key: ", k)
		fmt.Println("val: ", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello astaxie!") //这个写入到w的是输出到客户端的
}

func main() {
	ip := "0.0.0.0"
	if envip := os.Getenv("IP"); envip != "" {
		fmt.Printf("Use IP:%s in ENV\n", envip)
		ip = envip
	}
	port := "8080"
	if envport := os.Getenv("PORT"); envport != "" {
		fmt.Printf("Use PORT:%s in ENV\n", envport)
		port = envport
	}
	http.HandleFunc("/", sayHelloName)
	err := http.ListenAndServe(ip+":"+port, nil)
	if err != nil {
		log.Fatal("ListenAndserver: ", err)
	}
}
