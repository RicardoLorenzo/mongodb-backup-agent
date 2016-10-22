package storage

import (
	"bufio"
	"fmt"
	"strings"

	u "github.com/RicardoLorenzo/mongodb-backup-agent/storage/utils"
)

type BtrfsUtils struct{}

func (utils *BtrfsUtils) createSnapshot(s Snapshot) (bool, error) {
	stringUtils := u.StringUtils{}
	snapPath := stringUtils.StringConcat([]string{s.Volume.Path, "/", s.Name})

	commandUtils := u.CommandUtils{}
	c := u.Command{Binary: "btrfs", Args: []string{"subvolume", "snapshot", s.Volume.Path, snapPath}}

	success, err := commandUtils.RunCommand(c)
	if !success {
		return false, &volumeError{fmt.Sprint(err), err}
	}
	return true, nil
}

func (utils *BtrfsUtils) deleteSnapshot(s Snapshot) (bool, error) {
	stringUtils := u.StringUtils{}
	snapPath := stringUtils.StringConcat([]string{s.Volume.Path, "/", s.Name})

	commandUtils := u.CommandUtils{}
	c := u.Command{Binary: "btrfs", Args: []string{"subvolume", "delete", snapPath}}

	success, err := commandUtils.RunCommand(c)
	if !success {
		return false, &volumeError{fmt.Sprint(err), err}
	}
	return true, nil
}

func (utils *BtrfsUtils) listSnapshots(v Volume) ([]Snapshot, error) {
	commandUtils := u.CommandUtils{}
	c := u.Command{Binary: "btrfs", Args: []string{"subvolume", "list", v.Path}}

	success, err := commandUtils.RunCommand(c)
	if !success {
		return nil, &volumeError{fmt.Sprint(err), err}
	}

	output := commandUtils.CommandOutput()
	scanner := bufio.NewScanner(strings.NewReader(output))
	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())

		// snapshot name?
		fmt.Println(fields[0])
	}

	return nil, nil
}
