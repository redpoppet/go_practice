package main

import (
	"flag"
	"fmt"
	"module-main/submodule"
	p "mypackage"

	"github.com/link1st/example/stringutil"
)

var (
	str = ""
)

func init() {
	flag.StringVar(&str, "str", str, "输入字符")
	flag.Parse()
}

func main() {
	fmt.Println("hello world")
	p.New()
	submodule.Test()
	fmt.Printf("p.Old(3, 4): %v\n", p.Old(3, 4))
	if str == "" {
		fmt.Println("示例： go run main.go -str hello")
		fmt.Println("str 参数必填")
		flag.Usage()
		return
	}
	str = stringutil.Reversal(str)
	str = stringutil.ToUpper(str)
	fmt.Println(str)


}
