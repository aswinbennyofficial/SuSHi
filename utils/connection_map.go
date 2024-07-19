package utils

import (
	"fmt"
	"sync"
	"time"

	"github.com/aswinbennyofficial/SuSHi/models"
	"github.com/rs/zerolog/log"
)

var (
    uuidMap     = make(map[string]*models.SSHConnection)
    timeMap     = make(map[time.Time][]string)
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


func StoreInTimeBucket(key time.Time, uuid string) {
	mapMutex.Lock()
	defer mapMutex.Unlock()

	if _, exists := timeMap[key]; !exists {
		timeMap[key] = make([]string, 0)
	}
	timeMap[key] = append(timeMap[key], uuid)
}


func UpdateTimeBucket(key time.Time, uuid string) {
	mapMutex.Lock()
	defer mapMutex.Unlock()

	log.Debug().Msgf("Updating time bucket for %s with key %s", uuid, key)

	// fectch existing timeKey from uuidMap
	timeBucketKey :=uuidMap[uuid].TimeBucketKey

	// remove uuid from old time bucket
	if _, exists := timeMap[timeBucketKey]; exists {
		for i, id := range timeMap[timeBucketKey] {
			if id == uuid {
				timeMap[timeBucketKey] = append(timeMap[timeBucketKey][:i], timeMap[timeBucketKey][i+1:]...)
				break
			}
		}
	}

	// add uuid to new time bucket
	if _, exists := timeMap[key]; !exists {
		timeMap[key] = make([]string, 0)
	}
	timeMap[key] = append(timeMap[key], uuid)

	// update timeBucketKey in uuidMap
	uuidMap[uuid].TimeBucketKey = key

	log.Debug().Msgf("Updated time bucket for %s with key %s", uuid, key)
}

func CheckExpiredBuckets(){
	// start a ticker every 1 minute to check for expired buckets
	ticker := time.NewTicker(1 * time.Minute)

	for range ticker.C {
		removeExpiredBuckets()
	}
}

func removeExpiredBuckets() {
	mapMutex.Lock()
	defer mapMutex.Unlock()
	numDeleted := 0
	// Get the current time
	timeNow := RoundToNearestMinute(time.Now())

	// Get the time X minutes ago
	timeXMinutesAgo := timeNow.Add(-expirationTime)

	// Iterate over the timeMap and delete all entries older than X minutes
	for key, uuids := range timeMap {
		if key.Before(timeXMinutesAgo) {
			for _, uuid := range uuids {
				numDeleted++
				// close ssh connection
				uuidMap[uuid].Client.Close()
				// delete from uuidMap
				delete(uuidMap, uuid)
			}
			delete(timeMap, key)
		}
	}
	log.Debug().Msg("Deleted expired connections: "+fmt.Sprint(numDeleted))
}