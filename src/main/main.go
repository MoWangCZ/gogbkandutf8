package main

import (
	"fmt"
	_ "hello/imp"
	"github.com/axgle/mahonia"
)
func main(){
	fmt.Println("main-->")
	gbk := mahonia.NewEncoder("gbk").ConvertString("hello,世界")
	gbkbytes:= []byte(gbk)
	fmt.Println(gbk)
	fmt.Printf("%x",gbkbytes)
}
