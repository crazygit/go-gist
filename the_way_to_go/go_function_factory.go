package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	addJpeg := MakeAddSuffix(".jpeg")
	addBmp := MakeAddSuffix(".bmp")
	fmt.Println(addJpeg("file"))
	fmt.Println(addBmp("file"))
	end := time.Now()
	delta := end.Skkkub(start)
	fmt.Printf("use %s ", delta)

}

func MakeAddSuffix(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}
