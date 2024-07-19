package utils

import (
	"golang.org/x/crypto/ssh"
	"github.com/aswinbennyofficial/SuSHi/models"
	"github.com/rs/zerolog/log"
	"fmt"
	
)

// PublicKeyFile returns an ssh.AuthMethod that uses the private key file
func PublicKeyFile(key string, passphrase string) ssh.AuthMethod {
	var signer ssh.Signer
	var err error

	if passphrase == "" {
		// If there is no passphrase, use the private key directly
		signer, err = ssh.ParsePrivateKey([]byte(key))
	} else {
		// If there is a passphrase, decrypt the private key using the passphrase
		signer, err = ssh.ParsePrivateKeyWithPassphrase([]byte(key), []byte(passphrase))
	}

	if err != nil {
		log.Error().Msgf("Failed to parse private key: %v", err)
		return nil
	}

	return ssh.PublicKeys(signer)
}

func ConnectToMachine(machine models.Machine) (*ssh.Client, error) {
	authMethod := PublicKeyFile(machine.PrivateKey, machine.Passphrase)
	if authMethod == nil {
		return nil, fmt.Errorf("failed to create SSH authentication method")
	}

	sshConfig := &ssh.ClientConfig{
		User: machine.Username,
		Auth: []ssh.AuthMethod{
			authMethod,
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	address := fmt.Sprintf("%s:%s", machine.Hostname, machine.Port)
	client, err := ssh.Dial("tcp", address, sshConfig)
	if err != nil {
		log.Error().Msgf("Failed to connect to machine: %v", err)
		return nil, err
	}

	return client, nil
}



