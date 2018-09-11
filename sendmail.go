package sendmail

import (
	"os"
	"os/exec"
	"fmt"
	"strings"
	"encoding/base64"
)

func SendMail(from, name, to, title, body string) error{
	args := []string{
		"-f",
		from,
		"-F",
		name,
		"-o",
		"message-content-type=html",
		to,
	}
	cmd := exec.Command("sendmail", args...)

	baseTitle := "=?UTF-8?B?" + base64.StdEncoding.EncodeToString([]byte(title)) +"?="
	input := fmt.Sprintf("Subject:%s\n\n%s", baseTitle, body)
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
	return nil
}
