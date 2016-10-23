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
	c "github.com/RicardoLorenzo/mongodb-backup-agent/conf"
)

func main() {
	channel := make(chan os.Signal, 2)
	signal.Notify(channel, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-channel
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

	bcommands := b.NewBackupCommands(config)
	bclient, err := b.NewBackupClient(backupServer, backupPort, agentId)
	bclient.Register()

	for {
		commands, err := bclient.GetCommands()
		if err != nil {
			log.Print(err)
		}
		for _, command := range commands {
			bcommands.RunCommand(command)
		}
		time.Sleep(1 * time.Second)
	}
}
