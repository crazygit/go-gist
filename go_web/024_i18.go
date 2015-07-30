package main

import (
	"fmt"
	"time"
)

var locales map[string]map[string]string

func main() {
	locales = make(map[string]map[string]string, 2)
	en := make(map[string]string, 10)

	en["pea"] = "pea"
	en["bean"] = "bean"
	en["how old"] = "I am %d years old"
	en["time_zone"] = "America/Chicago"
	en["date_format"] = "%Y-%m-%d %H:%M:%S"
	locales["en"] = en

	cn := make(map[string]string, 10)
	cn["pea"] = "豌豆"
	cn["bean"] = "毛豆"
	cn["how old"] = "我今年%d岁了"
	cn["time_zone"] = "Asia/Shanghai"
	cn["date_format"] = "%Y年%m月%d日 %H时%M分%S秒"
	locales["zh-CN"] = cn

	lang := "zh-CN"
	fmt.Println(msg(lang, "pea"))
	fmt.Println(msg(lang, "bean"))
	fmt.Printf(msg(lang, "how old")+"\n", 30)

	loc, _ := time.LoadLocation(msg(lang, "time_zone"))
	t := time.Now()
	t = t.In(loc)
	fmt.Println(t.Format(time.RFC3339))

}

func msg(locale, key string) string {
	if v, ok := locales[locale]; ok {
		if v2, ok := v[key]; ok {
			return v2
		}
	}
	return ""
}
