package services

import (
	"context"

	api_proto "www.velocidex.com/golang/velociraptor/api/proto"
	config_proto "www.velocidex.com/golang/velociraptor/config/proto"
	flows_proto "www.velocidex.com/golang/velociraptor/flows/proto"
)

func GetVFSService(config_obj *config_proto.Config) (VFSService, error) {
	org_manager, err := GetOrgManager()
	if err != nil {
		return nil, err
	}

	return org_manager.Services(config_obj.OrgId).VFSService()
}

type VFSService interface {
	// Lists all the directories in the VFS path provided. This is
	// used by the tree widget in the GUI so it only returns
	// directories. For both files and directories see ListFiles()
	// below.
	ListDirectories(
		ctx context.Context,
		config_obj *config_proto.Config,
		client_id string,
		components []string) (*api_proto.VFSListResponse, error)

	StatDirectory(
		config_obj *config_proto.Config,
		client_id string,
		vfs_components []string) (*api_proto.VFSListResponse, error)

	StatDownload(
		config_obj *config_proto.Config,
		client_id string,
		accessor string,
		path_components []string) (*flows_proto.VFSDownloadInfo, error)
}
