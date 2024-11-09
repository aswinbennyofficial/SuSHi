package models

import (
	"time"

	"golang.org/x/crypto/ssh"
)

type SSHConnection struct {
	TimeBucketKey time.Time
	Client   *ssh.Client
}