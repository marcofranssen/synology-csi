// Copyright 2021 Synology Inc.

package models

import (
	"fmt"
	"github.com/SynologyOpenSource/synology-csi/pkg/dsm/webapi"
	"github.com/container-storage-interface/spec/lib/go/csi"
)

type CreateK8sVolumeSpec struct {
	// 16-byte fields (strings)
	DsmIp            string
	K8sVolumeName    string
	LunName          string
	LunDescription   string
	ShareName        string
	Location         string
	Type             string
	TargetName       string
	SourceSnapshotId string
	SourceVolumeId   string
	Protocol         string
	NfsVersion       string
	// 8-byte fields (int64, maps)
	Size       int64
	DevAttribs map[string]bool
	// 1-byte fields (bool)
	ThinProvisioning bool
	MultipleSession  bool
}

type K8sVolumeRespSpec struct {
	// Structs (sizes vary, keep together at top)
	Lun    webapi.LunInfo
	Target webapi.TargetInfo
	Share  webapi.ShareInfo
	// 16-byte fields (strings)
	DsmIp    string
	VolumeId string
	Location string
	Name     string
	Source   string
	Protocol string
	BaseDir  string
	// 8-byte fields (int64)
	SizeInBytes int64
}

type K8sSnapshotRespSpec struct {
	// 16-byte fields (strings)
	DsmIp      string
	Name       string
	Uuid       string
	ParentName string
	ParentUuid string
	Status     string
	Time       string // only for share snapshot delete
	RootPath   string
	Protocol   string
	// 8-byte fields (int64)
	SizeInBytes int64
	CreateTime  int64
}

type CreateK8sVolumeSnapshotSpec struct {
	// 16-byte fields (strings)
	K8sVolumeId  string
	SnapshotName string
	Description  string
	TakenBy      string
	// 1-byte fields (bool)
	IsLocked bool
}

type NodeStageVolumeSpec struct {
	// 16-byte fields (strings)
	VolumeId          string
	StagingTargetPath string
	Dsm               string
	Source            string
	FormatOptions     string
	// 8-byte fields (pointers)
	VolumeCapability *csi.VolumeCapability
}

type ByVolumeId []*K8sVolumeRespSpec

func (a ByVolumeId) Len() int           { return len(a) }
func (a ByVolumeId) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByVolumeId) Less(i, j int) bool { return a[i].VolumeId < a[j].VolumeId }

type BySnapshotAndParentUuid []*K8sSnapshotRespSpec

func (a BySnapshotAndParentUuid) Len() int      { return len(a) }
func (a BySnapshotAndParentUuid) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a BySnapshotAndParentUuid) Less(i, j int) bool {
	return fmt.Sprintf("%s/%s", a[i].ParentUuid, a[i].Uuid) < fmt.Sprintf("%s/%s", a[j].ParentUuid, a[j].Uuid)
}
