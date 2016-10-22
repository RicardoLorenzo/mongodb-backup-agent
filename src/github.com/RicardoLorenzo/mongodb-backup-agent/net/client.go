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

type NetworkClient struct {
	server     string
	port       int
	connection net.Conn
}

func NewNetworkClient(server string, port int) NetworkClient {
	client := NetworkClient{}
	client.server = server
	if port > 0 {
		client.port = port
	} else {
		client.port = 4444
	}
	return client
}

func (client *NetworkClient) Connect() error {
	stringUtils := u.StringUtils{}
	connectionString := stringUtils.StringConcat([]string{client.server, ":", strconv.Itoa(client.port)})

	var err error
	client.connection, err = net.Dial("tcp", connectionString)
	if err != nil {
		return &NetworkError{fmt.Sprint(err), err}
	}
	return nil
}

func (client *NetworkClient) Disconnect() {
	if client.connection != nil {
		client.connection.Close()
	}
}

func (client *NetworkClient) ReadLine() (string, error) {
	message, err := bufio.NewReader(client.connection).ReadString('\n')
	if err != nil {
		return "", &NetworkError{fmt.Sprint(err), err}
	}
	return message, nil
}

func (client *NetworkClient) WriteLine(line string) error {
	stringUtils := u.StringUtils{}
	writeLine := stringUtils.StringConcat([]string{line, "\n"})

	_, err := client.connection.Write([]byte(writeLine))
	if err != nil {
		return &NetworkError{fmt.Sprint(err), err}
	}
	return nil
}
