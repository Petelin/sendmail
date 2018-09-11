package main

import "github.com/petelin/sendmail"

func main() {
	sendmail.SendMail("Robot@notify.com","[Robot]", "859598732@qq.com", "this is title", "this is body\nand a new line")
}
