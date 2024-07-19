package utils

import (
	"sync"
	"github.com/aswinbennyofficial/SuSHi/models"
	"time"
)

var (
    uuidMap     = make(map[string]*models.SSHConnection)
    timeMap     = make(map[time.Time]map[string]struct{})
    mapMutex    sync.RWMutex
    expirationTime = 15 * time.Minute
)

func StoreSSHConnection(uuid string, conn *models.SSHConnection) {
	mapMutex.Lock()
	defer mapMutex.Unlock()

	uuidMap[uuid] = conn
}

func GetSSHConnection(uuid string) (*models.SSHConnection, bool) {
	mapMutex.RLock()
	defer mapMutex.RUnlock()

	conn, exists := uuidMap[uuid]
	return conn, exists
}

