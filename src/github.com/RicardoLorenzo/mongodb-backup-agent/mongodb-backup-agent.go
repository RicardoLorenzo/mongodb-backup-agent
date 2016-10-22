package main

import s "github.com/RicardoLorenzo/mongodb-backup-agent/storage"

func main() {
	vm := s.VolumeManager{FStype: "btrfs"}
	//v := s.Volume{name: "test", path: "/usr/test"}
	s := s.Snapshot{Name: "snap_test", Volume: s.Volume{Name: "test", Path: "/usr/test"}}

	vm.CreateSnapshot(s)
}
