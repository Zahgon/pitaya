// Copyright (c) nano Author and TFG Co. All Rights Reserved.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package pitaya

import (
	"context"
	"time"
)

// Group represents an agglomeration of UIDs which is used to manage
// users. Data sent to the group will be sent to all users in it.

// GroupCreate creates a group
func (app *App) GroupCreate(ctx context.Context, groupName string) error {
	_ = "STUB: not implemented"
	return nil
}

// GroupCreateWithTTL creates a group with given TTL
func (app *App) GroupCreateWithTTL(ctx context.Context, groupName string, ttlTime time.Duration) error {
	_ = "STUB: not implemented"
	return nil
}

// GroupMembers returns all member's UIDs
func (app *App) GroupMembers(ctx context.Context, groupName string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GroupBroadcast pushes the message to all members inside group
func (app *App) GroupBroadcast(ctx context.Context, frontendType, groupName, route string, v interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (app *App) sendDataToMembers(uids []string, frontendType, route string, v interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

// GroupContainsMember checks whether an UID is contained in group or not
func (app *App) GroupContainsMember(ctx context.Context, groupName, uid string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// GroupAddMember adds UID to group
func (app *App) GroupAddMember(ctx context.Context, groupName, uid string) error {
	_ = "STUB: not implemented"
	return nil
}

// GroupRemoveMember removes specified UID from group
func (app *App) GroupRemoveMember(ctx context.Context, groupName, uid string) error {
	_ = "STUB: not implemented"
	return nil
}

// GroupRemoveAll clears all UIDs
func (app *App) GroupRemoveAll(ctx context.Context, groupName string) error {
	_ = "STUB: not implemented"
	return nil
}

// GroupCountMembers get current member amount in group
func (app *App) GroupCountMembers(ctx context.Context, groupName string) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// GroupRenewTTL renews group with the initial TTL
func (app *App) GroupRenewTTL(ctx context.Context, groupName string) error {
	_ = "STUB: not implemented"
	return nil
}

// GroupDelete deletes whole group, including UIDs and base group
func (app *App) GroupDelete(ctx context.Context, groupName string) error {
	_ = "STUB: not implemented"
	return nil
}
