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
		to,
	}
	cmd := exec.Command("sendmail", args...)

	baseTitle := "=?UTF-8?B?" + base64.StdEncoding.EncodeToString([]byte(title)) +"?="
	input := fmt.Sprintf("Subject:%s\nContent-Type:text/html\n\n%s", baseTitle, fmt.Sprintf("<html><body>%s</body></html>", body))
	fmt.Println(input)
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
