package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
)

func IsIP(ip string) (b bool) {
	if m, _ := regexp.MatchString("^[0-9]{1,3}\\.[0-9]{1,3}\\.[0-9]{1,3}\\.[0-9]{1,3}$", ip); !m {
		return false
	}
	return true
}

func main() {
	ip := "127.0.0.1"
	fmt.Printf("%s is IP : %v", ip, IsIP(ip))
	resp, err := http.Get("http://www.baidu.com")
	if err != nil {
		fmt.Println("http get error.")
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("http read error.")
		return
	}
	src := string(body)
	re, _ := regexp.Compile("\\<[\\S\\s]+?\\>")
	src = re.ReplaceAllStringFunc(src, strings.ToLower)

	//去除STYLE
	re, _ = regexp.Compile("\\<style[\\S\\s]+?\\</style\\>")
	src = re.ReplaceAllString(src, "")

	//去除SCRIPT
	re, _ = regexp.Compile("\\<script[\\S\\s]+?\\</script\\>")
	src = re.ReplaceAllString(src, "")

	//去除所有尖括号内的HTML代码，并换成换行符
	re, _ = regexp.Compile("\\<[\\S\\s]+?\\>")
	src = re.ReplaceAllString(src, "\n")
	//去除连续的换行符
	re, _ = regexp.Compile("\\s{2,}")
	src = re.ReplaceAllString(src, "\n")

	fmt.Println(src)
}
