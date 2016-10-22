package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	b "github.com/RicardoLorenzo/mongodb-backup-agent/backup"
	c "github.com/RicardoLorenzo/mongodb-backup-agent/conf"
)

func main() {
	log.SetOutput(os.Stdout)
	config, err := c.NewConfig()
	if err != nil {
		log.Panic(fmt.Sprint(err), err)
	}

	var backupPort int
	agentId := config.GetProperty("backup.id")
	if len(agentId) == 0 {
		log.Panic("backup ID isn't configured")
	}
	backupServer := config.GetProperty("backup.server")
	if len(backupServer) == 0 {
		log.Panic("backup server isn't configured")
	}
	p := config.GetProperty("backup.port")
	if len(p) > 0 {
		backupPort, err = strconv.Atoi(p)
	} else {
		backupPort = 4444
	}
	if len(backupServer) == 0 {
		log.Panic("backup server isn't configured")
	}

	fsType := config.GetProperty("storage.type")
	if len(fsType) == 0 {
		fsType = "btrfs"
	}

	bclient, err := b.NewBackupClient(backupServer, backupPort, agentId)
	bclient.Register()

	//vm := s.VolumeManager{FStype: fsType}
	//v := s.Volume{name: "test", path: "/usr/test"}
	//s := s.Snapshot{Name: "snap_test", Volume: s.Volume{Name: "test", Path: "/usr/test"}}
	//vm.CreateSnapshot(s)
}
