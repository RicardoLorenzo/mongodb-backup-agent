package net

import (
	"bufio"
	"fmt"
	"net"
	"strconv"

	u "github.com/RicardoLorenzo/mongodb-backup-agent/utils"
)

type NetworkError struct {
	message string
	err     error
}

func (e *NetworkError) Error() string {
	return e.message
}

type BackupClient struct {
	server     string
	port       int
	connection net.Conn
}

func NewBackupClient(server string, port int) BackupClient {
	client := BackupClient{}
	client.server = server
	if port > 0 {
		client.port = port
	} else {
		client.port = 4444
	}
	return client
}

func (client *BackupClient) Connect() error {
	stringUtils := u.StringUtils{}
	connectionString := stringUtils.StringConcat([]string{client.server, ":", strconv.Itoa(client.port)})

	var err error
	client.connection, err = net.Dial("tcp", connectionString)
	if err != nil {
		return &NetworkError{fmt.Sprint(err), err}
	}
	return nil
}

func (client *BackupClient) Close() {
	if client.connection != nil {
		client.connection.Close()
	}
}

func (client *BackupClient) ReadLine() (string, error) {
	message, err := bufio.NewReader(client.connection).ReadString('\n')
	if err != nil {
		return "", &NetworkError{fmt.Sprint(err), err}
	}
	return message, nil
}

func (client *BackupClient) WriteLine(line string) error {
	stringUtils := u.StringUtils{}
	writeLine := stringUtils.StringConcat([]string{line, "\n"})

	_, err := client.connection.Write([]byte(writeLine))
	if err != nil {
		return &NetworkError{fmt.Sprint(err), err}
	}
	return nil
}
