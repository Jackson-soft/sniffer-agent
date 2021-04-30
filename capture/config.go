package capture

import (
	"math/rand"
	"time"

	log "github.com/sirupsen/logrus"
	sd "github.com/zr-hebo/sniffer-agent/session-dealer"
)

var (
	localIPAddr *string

	sessionPool = make(map[string]sd.ConnSession)
	// sessionPoolLock sync.Mutex
)

func init() {
	ipAddr, err := getLocalIPAddr()
	if err != nil {
		panic(err)
	}

	localIPAddr = &ipAddr
	log.Infof("parsed local ip address:%s", *localIPAddr)

	rand.Seed(time.Now().UnixNano())
}
