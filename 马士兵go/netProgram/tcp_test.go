package netProgram

import "testing"

func TestTcpServer(t *testing.T) {
	TcpServer()
}
func TestTcpClient(t *testing.T) {
	TcpClient()
}
func TestTcpTimeoutClient(t *testing.T) {
	TcpTimeoutClient()
}

func TestTcpBacklogServer(t *testing.T) {
	TcpBacklogServer()
}
func TestTcpBacklogClient(t *testing.T) {
	TcpBacklogClient()
}

func TestTcpServerRW(t *testing.T) {
	TcpServerRW()
}

func TestTcpClientRW(t *testing.T) {
	TcpClientRW()
}

func TestTcpW(t *testing.T) {
	TcpW()
}

func TestTcpServerRWConcurrency(t *testing.T) {
	TcpServerRWConcurrency()
}
func TestTcpClientRWConcurrency(t *testing.T) {
	TcpClientRWConcurrency()
}
