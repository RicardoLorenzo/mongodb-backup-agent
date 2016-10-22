package storage

import (
	"errors"
	"fmt"
)

type VolumeError struct {
	message string
	err     error
}

func (e *VolumeError) Error() string {
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
			return false, &VolumeError{fmt.Sprint(err), err}
		}
		return true, nil
	}
	return false, nil
}

func (volumeManager *VolumeManager) ListSnapshots(v Volume) ([]Snapshot, error) {
	switch volumeManager.FStype {
	case "btrs":
		btrfs := BtrfsUtils{}
		snapshots, err := btrfs.listSnapshots(v)
		if err != nil {
			return nil, &VolumeError{fmt.Sprint(err), err}
		}
		return snapshots, nil
	}
	return nil, &VolumeError{"unsupported filesystem", errors.New("unsuported filesystem")}
}

func (volumeManager *VolumeManager) DeleteSnapshot(s Snapshot) (bool, error) {
	switch volumeManager.FStype {
	case "btrs":
		utils := BtrfsUtils{}
		sucess, err := utils.deleteSnapshot(s)
		if !sucess {
			return false, &VolumeError{fmt.Sprint(err), err}
		}
		return true, nil
	}
	return false, nil
}
