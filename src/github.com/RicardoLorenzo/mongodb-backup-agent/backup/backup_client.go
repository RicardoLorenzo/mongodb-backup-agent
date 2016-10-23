package backup

import (
	"bufio"
	"fmt"
	"strings"

	n "github.com/RicardoLorenzo/mongodb-backup-agent/net"
	u "github.com/RicardoLorenzo/mongodb-backup-agent/utils"
)

type BackupError struct {
	Message string
	Err     error
}

func (e *BackupError) Error() string {
	return e.Message
}

type BackupCommand struct {
	Name, Path, Snapshot string
}

type BackupClient struct {
	agentId   string
	netClient n.NetworkClient
}

func NewBackupClient(server string, port int, agentId string) (BackupClient, error) {
	client := n.NewNetworkClient(server, port)
	backupClient := BackupClient{agentId: agentId, netClient: client}
	return backupClient, backupClient.netClient.Connect()
}

func (client *BackupClient) Register() error {
	stringUtils := u.StringUtils{}
	registrationLine := stringUtils.StringConcat([]string{"register agent ", client.agentId})
	err := client.netClient.WriteLine(registrationLine)
	if err != nil {
		return &BackupError{fmt.Sprint(err), err}
	}

	reply, err := client.netClient.ReadLine()
	if err != nil {
		return &BackupError{fmt.Sprint(err), err}
	}
	if reply != "registered" {
		message := stringUtils.StringConcat([]string{"cannot register backup client [", reply, "]"})
		return &BackupError{message, err}
	}
	_, err = client.verifyNetworkReply("registered")
	return err
}

func (client *BackupClient) Disconnect() {
	client.netClient.Disconnect()
}

func (client *BackupClient) GetCommands() ([]BackupCommand, error) {
	var backupCommands []BackupCommand
	err := client.netClient.WriteLine("get commands")
	if err != nil {
		return nil, &BackupError{fmt.Sprint(err), err}
	}

	var reply string
	reply, err = client.verifyNetworkReply("")
	scanner := bufio.NewScanner(strings.NewReader(reply))
	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())

		var path, snapshot string
		if len(fields) > 1 {
			path = fields[1]
		}
		if len(fields) > 2 {
			snapshot = fields[2]
		}

		backupCommands = append(backupCommands, BackupCommand{Name: fields[0], Path: path, Snapshot: snapshot})
	}
	return backupCommands, nil
}

func (client *BackupClient) verifyNetworkReply(message string) (string, error) {
	stringUtils := u.StringUtils{}
	reply, err := client.netClient.ReadLine()
	if err != nil {
		return "", &BackupError{fmt.Sprint(err), err}
	}
	if len(message) > 0 && reply != message {
		m := stringUtils.StringConcat([]string{"cannot register backup client [", reply, "]"})
		return reply, &BackupError{m, err}
	}
	return reply, nil
}
