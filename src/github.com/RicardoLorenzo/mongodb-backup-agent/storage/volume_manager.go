package storage

import (
	"errors"
	"fmt"
)

type volumeError struct {
	message string
	err     error
}

func (e *volumeError) Error() string {
	return e.message
}

type Volume struct {
	Name, Path string
}

type Snapshot struct {
	Name   string
	Volume Volume
}

type VolumeManager struct {
	FStype string
}

func (volumeManager *VolumeManager) CreateSnapshot(s Snapshot) (bool, error) {
	switch volumeManager.FStype {
	case "btrs":
		utils := BtrfsUtils{}
		sucess, err := utils.createSnapshot(s)
		if !sucess {
			return false, &volumeError{fmt.Sprint(err), err}
		}
		return true, nil
	}
	return false, nil
}

func (volumeManager *VolumeManager) ListSnapshots(v Volume) ([]Snapshot, error) {
	switch volumeManager.FStype {
	case "btrs":
		utils := BtrfsUtils{}
		snapshots, err := utils.listSnapshots(v)
		if snapshots == nil {
			return nil, &volumeError{fmt.Sprint(err), err}
		}
		// TODO
	}
	return nil, &volumeError{"unsupported filesystem", errors.New("unsuported filesystem")}
}

func (volumeManager *VolumeManager) DeleteSnapshot(s Snapshot) (bool, error) {
	switch volumeManager.FStype {
	case "btrs":
		utils := BtrfsUtils{}
		sucess, err := utils.deleteSnapshot(s)
		if !sucess {
			return false, &volumeError{fmt.Sprint(err), err}
		}
		return true, nil
	}
	return false, nil
}
