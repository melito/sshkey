// +build !windows

package sshkey

import (
	"fmt"
	"github.com/howeyc/gopass"
)

func unixReadPassword(prompt string) (password string, err error) {
	fmt.Printf("%s", prompt)
	binPass := gopass.GetPasswd()
	password = string(binPass)
	return password, nil
}

func init() {
	PasswordPrompt = unixReadPassword
}
