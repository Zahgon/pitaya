package groups

import (
	"context"
	"sync"
	"time"

	"github.com/topfreegames/pitaya/v3/pkg/config"
	"go.etcd.io/etcd/api/v3/mvccpb"
	clientv3 "go.etcd.io/etcd/client/v3"
)

var (
	clientInstance     *clientv3.Client
	transactionTimeout time.Duration
	etcdOnce           sync.Once
)

// EtcdGroupService base ETCD struct solution
type EtcdGroupService struct {
	cancelFunc context.CancelFunc
}

// NewEtcdGroupService returns a new group instance
func NewEtcdGroupService(conf config.EtcdGroupServiceConfig, clientOrNil *clientv3.Client) (*EtcdGroupService, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func initClientInstance(config config.EtcdGroupServiceConfig, clientOrNil *clientv3.Client) error {
	_ = "STUB: not implemented"
	return nil
}

func createBaseClient(config config.EtcdGroupServiceConfig) (*clientv3.Client, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func groupKey(groupName string) string { _ = "STUB: not implemented"; return "" }

func memberKey(groupName, uid string) string { _ = "STUB: not implemented"; return "" }

func getGroupKV(ctx context.Context, groupName string) (*mvccpb.KeyValue, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *EtcdGroupService) createGroup(ctx context.Context, groupName string, leaseID clientv3.LeaseID) error {
	_ = "STUB: not implemented"
	return nil
}

// GroupCreate creates a group struct inside ETCD, without TTL
func (c *EtcdGroupService) GroupCreate(ctx context.Context, groupName string) error {
	_ = "STUB: not implemented"
	return nil
}

// GroupCreateWithTTL creates a group struct inside ETCD, with TTL, using leaseID
func (c *EtcdGroupService) GroupCreateWithTTL(ctx context.Context, groupName string, ttlTime time.Duration) error {
	_ = "STUB: not implemented"
	return nil
}

// GroupMembers returns all member's UIDs
func (c *EtcdGroupService) GroupMembers(ctx context.Context, groupName string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GroupContainsMember checks whether a UID is contained in current group or not
func (c *EtcdGroupService) GroupContainsMember(ctx context.Context, groupName, uid string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// GroupAddMember adds UID to group
func (c *EtcdGroupService) GroupAddMember(ctx context.Context, groupName, uid string) error {
	_ = "STUB: not implemented"
	return nil
}

// GroupRemoveMember removes specified UID from group
func (c *EtcdGroupService) GroupRemoveMember(ctx context.Context, groupName, uid string) error {
	_ = "STUB: not implemented"
	return nil
}

// GroupRemoveAll clears all UIDs in the group
func (c *EtcdGroupService) GroupRemoveAll(ctx context.Context, groupName string) error {
	_ = "STUB: not implemented"
	return nil
}

// GroupDelete deletes the whole group, including members and base group
func (c *EtcdGroupService) GroupDelete(ctx context.Context, groupName string) error {
	_ = "STUB: not implemented"
	return nil
}

// GroupCountMembers get current member amount in group
func (c *EtcdGroupService) GroupCountMembers(ctx context.Context, groupName string) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// GroupRenewTTL will renew ETCD lease TTL
func (c *EtcdGroupService) GroupRenewTTL(ctx context.Context, groupName string) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *EtcdGroupService) Close() { _ = "STUB: not implemented"; return }
