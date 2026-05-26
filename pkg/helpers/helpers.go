package helpers

import (
	"bufio"
	"os/exec"
	"testing"
	"time"

	"github.com/nats-io/nats-server/v2/server"
	clientv3 "go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/tests/v3/integration"
)

// GetFreePort returns a free port
func GetFreePort(t testing.TB) int { _ = "STUB: not implemented"; return 0 }

// GetMapKeys returns a string slice with the map keys
func GetMapKeys(t *testing.T, m interface{}) []string { _ = "STUB: not implemented"; return nil }

// GetTestNatsServer gets a test nats server
func GetTestNatsServer(t *testing.T) *server.Server { _ = "STUB: not implemented"; return nil }

// GetTestEtcd gets a test in memory etcd server
func GetTestEtcd(t *testing.T) (*integration.ClusterV3, *clientv3.Client) {
	_ = "STUB: not implemented"
	return nil, nil
}

// WriteFile test helper
func WriteFile(t *testing.T, filepath string, bytes []byte) { _ = "STUB: not implemented"; return }

// ReadFile test helper
func ReadFile(t *testing.T, filepath string) []byte { _ = "STUB: not implemented"; return nil }

// StartProcess starts a process
func StartProcess(t testing.TB, program string, args ...string) *exec.Cmd {
	_ = "STUB: not implemented"
	return nil
}

func waitForServerToBeReady(t testing.TB, out *bufio.Reader) { _ = "STUB: not implemented"; return }

// StartServer starts a server
func StartServer(
	t testing.TB,
	frontend, debug bool,
	svType string,
	port int,
	sdPrefix string,
	grpc, lazyConnection bool,
) func() {
	_ = "STUB: not implemented"
	return nil
}

// always use a random port for prometheus, to avoid e2e conflicts

// FixtureGoldenFileName returns the golden file name on fixtures path
func FixtureGoldenFileName(t *testing.T, name string) string { _ = "STUB: not implemented"; return "" }

func vetExtras(extras []interface{}) (bool, string) { _ = "STUB: not implemented"; return false, "" }

func pollFuncReturn(f interface{}) (interface{}, error) { _ = "STUB: not implemented"; return nil, nil }

// ShouldEventuallyReceive should asserts that eventually channel c receives a value
func ShouldEventuallyReceive(t testing.TB, c interface{}, timeouts ...time.Duration) interface{} {
	_ = "STUB: not implemented"
	return nil
}

// ShouldAlwaysReturn asserts that the return of f should always be v, timeouts: 0 - evaluation interval, 1 - timeout
func ShouldAlwaysReturn(t testing.TB, f interface{}, v interface{}, timeouts ...time.Duration) {
	_ = "STUB: not implemented"
	return
}

// ShouldEventuallyReturn asserts that eventually the return of f should be v, timeouts: 0 - evaluation interval, 1 - timeout
func ShouldEventuallyReturn(t testing.TB, f interface{}, v interface{}, timeouts ...time.Duration) {
	_ = "STUB: not implemented"
	return
}
