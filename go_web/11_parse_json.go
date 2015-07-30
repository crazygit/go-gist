package main

import (
	"encoding/json"
	"fmt"
)

type Server struct {
	ServerName string
	ServerIP   string
}

type Serverslice struct {
	Servers []Server
}

func main() {
	var s Serverslice
	a := `{"servers":[{"serverName":"Shanghai_VPN","serverIP":"127.0.0.1"},{"serverName":"Beijing_VPN","serverIP":"127.0.0.2"}]}`
	json.Unmarshal([]byte(a), &s)
	fmt.Println(s)

	b := []byte(`{"Name":"Wednesday","Age":6,"Parents":["Gomez","Morticia"]}`)
	var f interface{}
	err := json.Unmarshal(b, &f)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(f)
	m, ok := f.(map[string]interface{})
	fmt.Println(ok)
	for k, v := range m {
		switch vv := v.(type) {
		case string:
			fmt.Println(k, "is string", vv)
		}
	}
}
