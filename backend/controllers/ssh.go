package controllers

import (
	"encoding/json"
	"io"
	"net/http"
	"time"

	// "os/user"
	"sync"

	"github.com/aswinbennyofficial/SuSHi/models"
	"github.com/aswinbennyofficial/SuSHi/utils"

	// "github.com/google/uuid"
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

	uuid := r.URL.Query().Get("uuid")
	log.Debug().Msgf("UUID: %s", uuid)
	sshConnection,exists := utils.GetSSHConnection(uuid)
	if !exists {
		log.Printf("No SSH connection found for UUID: %s", uuid)
		// utils.ResponseHelper(w, http.StatusNotFound, "No SSH connection found for UUID", nil)
		return
	}


	client := sshConnection.Client
	

	
	// Open a session
	session, err := client.NewSession()
	if err != nil {
		log.Error().Msgf("Failed to create session: %v", err)
	}
	defer session.Close()

	allocatePseudoTerminal(session)

	// Set up pipes for stdin, stdout, and stderr
	stdin,stdout,err := getSessionPipe(session)
	if err != nil {
		log.Error().Msgf("Failed to set up pipes: %v", err)
	}
	

	// Start remote shell
	err = session.Shell()
	if err != nil {
		log.Error().Msgf("Failed to start shell: %v", err)
	}

	// Wait group for managing goroutines
	var wg sync.WaitGroup
	wg.Add(2)

	// Goroutine to read from WebSocket and write to SSH stdin
	go readFromWebSocket(conn, stdin,&wg, uuid)

	// Goroutine to read from SSH stdout and write to WebSocket
	go writeToWebSocket(conn, stdout,&wg)

	// Wait for goroutines to finish
	wg.Wait()
}

func getSessionPipe(session *ssh.Session)(io.WriteCloser, io.Reader, error){
	stdin, err := session.StdinPipe()
	if err != nil {
		return nil,nil,err
	}
	stdout, err := session.StdoutPipe()
	if err != nil {
		return nil,nil,err
	}

	return stdin,stdout,nil
	
}

func allocatePseudoTerminal(session *ssh.Session)(error){
	// Allocate a pseudo terminal
	modes := ssh.TerminalModes{
		ssh.ECHO:          1,
		ssh.TTY_OP_ISPEED: 14400,
		ssh.TTY_OP_OSPEED: 14400,
	}

	// Request a pseudo terminal
	err := session.RequestPty("xterm", 600, 800, modes)
	if err != nil {
		return err
	}

	return nil
}

func readFromWebSocket(conn *websocket.Conn, stdin io.WriteCloser, wg *sync.WaitGroup, uuid string) {
    defer wg.Done()
    for {
        _, msg, err := conn.ReadMessage()
        if err != nil {
            log.Error().Err(err).Msg("Error reading WebSocket message")
            return
        }

        var message models.Message
        err = json.Unmarshal(msg, &message)
        if err != nil {
            log.Error().Err(err).Msg("Error unmarshaling message")
            continue
        }

        switch message.Type {
        case "heartbeat":
            now := utils.RoundToNearestMinute(time.Now())
            utils.UpdateTimeBucket(now, uuid)
            log.Debug().Str("UUID", uuid).Time("Timestamp", now).Msg("Heartbeat received, time bucket updated")
        default:
            _, err = stdin.Write([]byte(message.Data))
            if err != nil {
                log.Error().Err(err).Msg("Error writing to SSH stdin")
                return
            }
        }
    }
}

func writeToWebSocket(conn *websocket.Conn, stdout io.Reader,wg *sync.WaitGroup) {
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
}