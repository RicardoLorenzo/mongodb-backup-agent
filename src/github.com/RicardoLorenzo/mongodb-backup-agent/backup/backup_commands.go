package backup

import c "github.com/RicardoLorenzo/mongodb-backup-agent/conf"
import s "github.com/RicardoLorenzo/mongodb-backup-agent/storage"

type BackupCommandReply struct {
	Message   string
	Volumes   []s.Volume
	Snapshots []s.Snapshot
}

type BackupCommands struct {
	vm s.VolumeManager
}

func NewBackupCommands(config c.Config) BackupCommands {
	fsType := config.GetProperty("storage.type")
	if len(fsType) == 0 {
		fsType = "btrfs"
	}

	vm := s.VolumeManager{FStype: fsType}
	return BackupCommands{vm: vm}
}

func (commands *BackupCommands) RunCommand(command BackupCommand) (BackupCommandReply, error) {
	switch command.Name {
	case "create_snapshot":
		volume := s.Volume{Name: command.Name, Path: command.Path}
		snapshot := s.Snapshot{Name: command.Snapshot, Volume: volume}
		err := commands.vm.CreateSnapshot(snapshot)
		if err != nil {
			return BackupCommandReply{Message: ""}, err
		}
		return BackupCommandReply{Message: "snapshot created"}, nil
	case "delete_snapshot":
		volume := s.Volume{Name: command.Name, Path: command.Path}
		snapshot := s.Snapshot{Name: command.Snapshot, Volume: volume}
		err := commands.vm.DeleteSnapshot(snapshot)
		if err != nil {
			return BackupCommandReply{Message: ""}, err
		}
		return BackupCommandReply{Message: "snapshot deleted"}, nil
	case "list_snapshots":
		volume := s.Volume{Name: command.Name, Path: command.Path}
		snapshots, err := commands.vm.ListSnapshots(volume)
		if err != nil {
			return BackupCommandReply{Message: ""}, err
		}
		return BackupCommandReply{Snapshots: snapshots}, nil
	}
	return BackupCommandReply{Message: ""}, &BackupError{Message: "invalid method"}
}
