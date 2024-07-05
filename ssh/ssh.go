package ssh

import (
	"net/http"
	"sync"

	"github.com/aswinbennyofficial/SuSHi/models"
	"github.com/gorilla/websocket"
	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/ssh"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}


func HandleSSHConnection(config models.Config,w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("Error upgrading to WebSocket: %v", err)
		return
	}
	defer conn.Close()

	// Hardcoded SSH access values (replace with your actual values)
	sshHost := config.SSHConfig.SSHHost
	sshPort := config.SSHConfig.SSHPort
	sshUser := config.SSHConfig.SSHUser
	privateKey := config.SSHConfig.PrivateKey

	// Parse the private key
	signer, err := ssh.ParsePrivateKey([]byte(privateKey))
	if err != nil {
		log.Fatal().Msgf("Failed to parse private key: %v", err)
	}

	sshconfig := &ssh.ClientConfig{
		User: sshUser,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(signer),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	// Connect to SSH server
	client, err := ssh.Dial("tcp", sshHost+":"+sshPort, sshconfig)
	if err != nil {
		log.Fatal().Msgf("Failed to dial SSH server: %v", err)
	}
	defer client.Close()

	// Open a session
	session, err := client.NewSession()
	if err != nil {
		log.Fatal().Msgf("Failed to create session: %v", err)
	}
	defer session.Close()

	// Allocate a pseudo terminal
	modes := ssh.TerminalModes{
		ssh.ECHO:          1,
		ssh.TTY_OP_ISPEED: 14400,
		ssh.TTY_OP_OSPEED: 14400,
	}

	// Request a pseudo terminal
	err = session.RequestPty("xterm", 80, 40, modes)
	if err != nil {
		log.Fatal().Msgf("Failed to request pseudo terminal: %v", err)
	}

	// Set up pipes for stdin, stdout, and stderr
	stdin, err := session.StdinPipe()
	if err != nil {
		log.Fatal().Msgf("Failed to create stdin pipe: %v", err)
	}
	stdout, err := session.StdoutPipe()
	if err != nil {
		log.Fatal().Msgf("Failed to create stdout pipe: %v", err)
	}
	// stderr, err := session.StderrPipe()
	if err != nil {
		log.Fatal().Msgf("Failed to create stderr pipe: %v", err)
	}

	// Start remote shell
	err = session.Shell()
	if err != nil {
		log.Fatal().Msgf("Failed to start shell: %v", err)
	}

	// Wait group for managing goroutines
	var wg sync.WaitGroup
	wg.Add(2)

	// Goroutine to read from WebSocket and write to SSH stdin
	go func() {
		defer wg.Done()
		for {
			_, msg, err := conn.ReadMessage()
			if err != nil {
				log.Printf("Error reading WebSocket message: %v", err)
				return
			}
			_, err = stdin.Write(msg)
			if err != nil {
				log.Printf("Error writing to SSH stdin: %v", err)
				return
			}
		}
	}()

	// Goroutine to read from SSH stdout and write to WebSocket
	go func() {
		defer wg.Done()
		buf := make([]byte, 1024)
		for {
			n, err := stdout.Read(buf)
			if err != nil {
				log.Printf("Error reading SSH stdout: %v", err)
				return
			}
			err = conn.WriteMessage(websocket.TextMessage, buf[:n])
			if err != nil {
				log.Printf("Error writing WebSocket message: %v", err)
				return
			}
		}
	}()

	// Wait for goroutines to finish
	wg.Wait()
}