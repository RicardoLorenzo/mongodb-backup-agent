package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	b "github.com/RicardoLorenzo/mongodb-backup-agent/backup"
)

func main() {
	c := make(chan os.Signal, 2)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		os.Exit(1)
	}()

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

	commands := NewBackupCommands(config)
	bclient, err := b.NewBackupClient(backupServer, backupPort, agentId)
	bclient.Register()

	for {
		commands := bclient.GetCommands()
		for _, c := range commands {
			commands.RunCommand(c)
		}
		time.Sleep(1 * time.Second)
	}
}
