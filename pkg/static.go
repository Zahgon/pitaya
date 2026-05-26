// Copyright (c) TFG Co. All Rights Reserved.
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

	"github.com/golang/protobuf/proto"
	"github.com/spf13/viper"
	"github.com/topfreegames/pitaya/v3/pkg/cluster"
	"github.com/topfreegames/pitaya/v3/pkg/component"
	"github.com/topfreegames/pitaya/v3/pkg/config"
	"github.com/topfreegames/pitaya/v3/pkg/interfaces"
	"github.com/topfreegames/pitaya/v3/pkg/metrics"
	"github.com/topfreegames/pitaya/v3/pkg/router"
	"github.com/topfreegames/pitaya/v3/pkg/session"
	"github.com/topfreegames/pitaya/v3/pkg/worker"
)

var DefaultApp Pitaya

// Configure configures the app
func Configure(
	isFrontend bool,
	serverType string,
	serverMode ServerMode,
	serverMetadata map[string]string,
	cfgs ...*viper.Viper,
) {
	_ = "STUB: not implemented"
	return
}

func GetDieChan() chan bool { _ = "STUB: not implemented"; return nil }

func SetDebug(debug bool) { _ = "STUB: not implemented"; return }

func SetHeartbeatTime(interval time.Duration) { _ = "STUB: not implemented"; return }

func GetServerID() string { _ = "STUB: not implemented"; return "" }

func GetMetricsReporters() []metrics.Reporter { _ = "STUB: not implemented"; return nil }

func GetServer() *cluster.Server { _ = "STUB: not implemented"; return nil }

func GetServerByID(id string) (*cluster.Server, error) { _ = "STUB: not implemented"; return nil, nil }

func GetServersByType(t string) (map[string]*cluster.Server, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func GetServers() []*cluster.Server { _ = "STUB: not implemented"; return nil }

func GetSessionFromCtx(ctx context.Context) session.Session {
	_ = "STUB: not implemented"
	return *new(session.Session)
}

func Start() { _ = "STUB: not implemented"; return }

func SetDictionary(dict map[string]uint16) error { _ = "STUB: not implemented"; return nil }

func AddRoute(serverType string, routingFunction router.RoutingFunc) error {
	_ = "STUB: not implemented"
	return nil
}

func Shutdown() { _ = "STUB: not implemented"; return }

func StartWorker() { _ = "STUB: not implemented"; return }

func RegisterRPCJob(rpcJob worker.RPCJob) error { _ = "STUB: not implemented"; return nil }

func Documentation(getPtrNames bool) (map[string]interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func IsRunning() bool { _ = "STUB: not implemented"; return false }

func RPC(ctx context.Context, routeStr string, reply proto.Message, arg proto.Message) error {
	_ = "STUB: not implemented"
	return nil
}

func RPCTo(ctx context.Context, serverID, routeStr string, reply proto.Message, arg proto.Message) error {
	_ = "STUB: not implemented"
	return nil
}

func ReliableRPC(routeStr string, metadata map[string]interface{}, reply, arg proto.Message) (jid string, err error) {
	_ = "STUB: not implemented"
	return "", nil
}

func ReliableRPCWithOptions(routeStr string, metadata map[string]interface{}, reply, arg proto.Message, opts *config.EnqueueOpts) (jid string, err error) {
	_ = "STUB: not implemented"
	return "", nil
}

func SendPushToUsers(route string, v interface{}, uids []string, frontendType string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func SendKickToUsers(uids []string, frontendType string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func GroupCreate(ctx context.Context, groupName string) error {
	_ = "STUB: not implemented"
	return nil
}

func GroupCreateWithTTL(ctx context.Context, groupName string, ttlTime time.Duration) error {
	_ = "STUB: not implemented"
	return nil
}

func GroupMembers(ctx context.Context, groupName string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func GroupBroadcast(ctx context.Context, frontendType, groupName, route string, v interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func GroupContainsMember(ctx context.Context, groupName, uid string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func GroupAddMember(ctx context.Context, groupName, uid string) error {
	_ = "STUB: not implemented"
	return nil
}

func GroupRemoveMember(ctx context.Context, groupName, uid string) error {
	_ = "STUB: not implemented"
	return nil
}

func GroupRemoveAll(ctx context.Context, groupName string) error {
	_ = "STUB: not implemented"
	return nil
}

func GroupCountMembers(ctx context.Context, groupName string) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func GroupRenewTTL(ctx context.Context, groupName string) error {
	_ = "STUB: not implemented"
	return nil
}

func GroupDelete(ctx context.Context, groupName string) error {
	_ = "STUB: not implemented"
	return nil
}

func Register(c component.Component, options ...component.Option) {
	_ = "STUB: not implemented"
	return
}

func RegisterRemote(c component.Component, options ...component.Option) {
	_ = "STUB: not implemented"
	return
}

func RegisterModule(module interfaces.Module, name string) error {
	_ = "STUB: not implemented"
	return nil
}

func RegisterModuleAfter(module interfaces.Module, name string) error {
	_ = "STUB: not implemented"
	return nil
}

func RegisterModuleBefore(module interfaces.Module, name string) error {
	_ = "STUB: not implemented"
	return nil
}

func GetModule(name string) (interfaces.Module, error) {
	_ = "STUB: not implemented"
	return *new(interfaces.Module), nil
}

func GetNumberOfConnectedClients() int64 { _ = "STUB: not implemented"; return 0 }

func IsReady(ctx context.Context) bool { _ = "STUB: not implemented"; return false }
