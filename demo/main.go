package main

import (
	"fmt"
	"github.com/petelin/sendmail"
)

func main() {
	err := sendmail.SendMail("x@x.com","Robot", "859598732@qq.com", "重新启动TRADE系统", "测试邮件发送系统.")
	fmt.Println(err)
}
