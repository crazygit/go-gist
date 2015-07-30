package main

import (
	"html/template"
	"os"
)

type Friend struct {
	Fname string
}
type Person struct {
	UserName string
	Emails   []string
	Friends  []*Friend
}

func main() {
	t := template.New("fieldname example")
	t, _ = t.Parse("hello {{.UserName}}! {{.email}}")
	p := Person{UserName: "world"}
	t.Execute(os.Stdout, p)

	f1 := Friend{Fname: "xiaowang"}
	f2 := Friend{Fname: "xiaogang"}
	t = template.New("filedname example")
	t, _ = t.Parse(`hello {{.UserName}}!
	{{range .Emails}}
		an email {{.}}
	{{end}}
	{{with .Friends}}
	{{range .}}
		my friend name is {{.Fname}}
	{{end}}
	{{end}}
	`)
	p = Person{UserName: "crazygit",
		Emails:  []string{"crazygit@foxmail.com", "crazygit@gmail.com"},
		Friends: []*Friend{&f1, &f2}}
	t.Execute(os.Stdout, p)

	tEmpty := template.New("template test")
	tEmpty = template.Must(tEmpty.Parse("空 pipeline if demo: {{if ``}} 不会输出. {{end}}\n"))
	tEmpty.Execute(os.Stdout, nil)

	tWithValue := template.New("template test")
	tWithValue = template.Must(tWithValue.Parse("不为空的 pipeline if demo: {{if `anything`}} 我有内容，我会输出. {{end}}\n"))
	tWithValue.Execute(os.Stdout, nil)

	tIfElse := template.New("template test")
	tIfElse = template.Must(tIfElse.Parse("if-else demo: {{if `anything`}} if部分 {{else}} else部分.{{end}}\n"))
	tIfElse.Execute(os.Stdout, nil)
}
