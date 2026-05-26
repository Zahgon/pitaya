package groups

import (
	"context"
	"sync"
	"time"

	"github.com/topfreegames/pitaya/v3/pkg/config"
)

var (
	memoryGroupsMu sync.RWMutex
	memoryGroups   map[string]*MemoryGroup
	memoryOnce     sync.Once
	globalCtx      context.Context
	globalCancel   context.CancelFunc
	cleanupWG      sync.WaitGroup
	cleanupOnce    sync.Once
)

// MemoryGroupService base in server memory solution
type MemoryGroupService struct {
	cancelFunc context.CancelFunc
}

// MemoryGroup is the struct stored in each group key(which is the name of the group)
type MemoryGroup struct {
	Uids        []string
	LastRefresh int64
	TTL         int64
}

// NewMemoryGroupService returns a new group instance
func NewMemoryGroupService(config config.MemoryGroupConfig) *MemoryGroupService {
	_ = "STUB: not implemented"
	return nil
}

// All services share the same cancel function

func groupTTLCleanup(ctx context.Context, interval time.Duration) {
	_ = "STUB: not implemented"
	return
}

// Check if context is cancelled before processing

// GroupCreate creates a group without TTL
func (c *MemoryGroupService) GroupCreate(ctx context.Context, groupName string) error {
	_ = "STUB: not implemented"
	return nil
}

// GroupCreateWithTTL creates a group with TTL, which the go routine will clean later
func (c *MemoryGroupService) GroupCreateWithTTL(ctx context.Context, groupName string, ttlTime time.Duration) error {
	_ = "STUB: not implemented"
	return nil
}

// GroupMembers returns all member's UID in given group
func (c *MemoryGroupService) GroupMembers(ctx context.Context, groupName string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GroupContainsMember check whether an UID is contained in given group or not
func (c *MemoryGroupService) GroupContainsMember(ctx context.Context, groupName, uid string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// GroupAddMember adds UID to group
func (c *MemoryGroupService) GroupAddMember(ctx context.Context, groupName, uid string) error {
	_ = "STUB: not implemented"
	return nil
}

// GroupRemoveMember removes specific UID from group
func (c *MemoryGroupService) GroupRemoveMember(ctx context.Context, groupName, uid string) error {
	_ = "STUB: not implemented"
	return nil
}

// GroupRemoveAll clears all UIDs from group
func (c *MemoryGroupService) GroupRemoveAll(ctx context.Context, groupName string) error {
	_ = "STUB: not implemented"
	return nil
}

// GroupDelete deletes the whole group, including members and base group
func (c *MemoryGroupService) GroupDelete(ctx context.Context, groupName string) error {
	_ = "STUB: not implemented"
	return nil
}

// GroupCountMembers get current member amount in group
func (c *MemoryGroupService) GroupCountMembers(ctx context.Context, groupName string) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// GroupRenewTTL will renew lease TTL
func (c *MemoryGroupService) GroupRenewTTL(ctx context.Context, groupName string) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *MemoryGroupService) Close() {
	_ = "STUB: not implemented"
	// Only cancel once, even if Close() is called multiple times
	return
}

// Wait for the goroutine to exit
// Use a channel with timeout to prevent indefinite blocking

// Goroutine exited successfully

// Timeout - this should not happen in normal operation,
// but we continue to prevent tests from hanging
