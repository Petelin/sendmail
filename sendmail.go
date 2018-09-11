package sendmail

import (
	"os"
	"os/exec"
	"log"
	"fmt"
	"strings"
)

func SendMail(from, name, to, title, body string) error{
	args := []string{
		"-f",
		from,
		"-F",
		name,
		to,
	}
	cmd := exec.Command("sendmail", args...)

	input := fmt.Sprintf("Subject:%s\n\n%s", title, body)
	cmd.Stdin = strings.NewReader(input)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	if err := cmd.Start(); err != nil {
		return err
	}
	err := cmd.Wait()
	if err != nil{
		return err
	}
}
