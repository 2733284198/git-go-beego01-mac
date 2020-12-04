package main

import (
	"fmt"

	"github.com/astaxie/beego/httplib"
)

func main() {
	fmt.Println("\n\n ==> 29 beego httplib ")

	req := httplib.Get("http://www.baidu.com")
	//strdata, err = req.Body().String()
	strdata, err := req.String()

	if nil != err {
		fmt.Println(strdata)
	} else {
		//fmt.Println("==》 没有数据")
		fmt.Println(strdata)
	}
}
