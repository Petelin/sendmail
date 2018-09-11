package main

import (
	"github.com/petelin/sendmail"
	"fmt"
)

func main() {
	err := sendmail.SendMail("x@x.com","Robot", "859598732@qq.com", "this is title中文测试", "this is body\n<h1>内容</h1>")
	fmt.Println(err)
}
