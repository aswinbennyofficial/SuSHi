package models

import (
	"golang.org/x/crypto/ssh"
)

type SSHConnection struct {
	Username string
	Client   *ssh.Client
}