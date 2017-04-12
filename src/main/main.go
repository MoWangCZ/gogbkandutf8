package main

import (
	"fmt"
	_ "hello/imp"	//这里的下划线表示只是调用hello/imp里面的init函数
	"github.com/axgle/mahonia"
	"common"
)
func main(){
	// 1:utf8->gbk,因为go编辑器默认为UTF8，所以汉字都是UTF8编码的
	gbk := mahonia.NewEncoder("gbk").ConvertString("hello,世界")
	fmt.Println(gbk)
	common.BytesToHexSpace([]byte(gbk))
	fmt.Printf("GBK BYTES:%x\n",[]byte(gbk))//打印出来就是GBK正确编码
	fmt.Println("=============")
	// 2:gbk->utf8
	utf82:= mahonia.NewDecoder("gbk").ConvertString(gbk)
	fmt.Println(utf82)
	common.BytesToHexSpace([]byte(utf82))
}


