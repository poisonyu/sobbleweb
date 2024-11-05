package main

import (
	"fmt"
	"reflect"
	"strings"
	"time"
)

func Print(x interface{}) {
	v := reflect.ValueOf(x)
	t := v.Type()
	fmt.Printf("type %s\n", t)
	for i := 0; i < v.NumMethod(); i++ {
		// v.Method返回一个方法值(reflect.Value)，可以使用reflect.Value.Call调用该方法
		methType := v.Method(i)
		name := t.Method(i).Name
		// if name == "Round" {
		// 	res := methType.Call([]reflect.Value{reflect.ValueOf("1s")})
		// }
		// t.Method描述了这个方法的名称和类型
		fmt.Printf("func (%s) %s%s\n", t, name, methType.String())
	}
}

func main() {
	Print(time.Hour)
	fmt.Println()
	Print(new(strings.Replacer))
}
