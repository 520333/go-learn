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

func TestTcpServerFormat(t *testing.T) {
	TcpServerFormat()
}
func TestTcpClientFormat(t *testing.T) {
	TcpClientFormat()
}

func TestTcpServerShort(t *testing.T) {
	TcpServerShort()
}
func TestTcpClientShort(t *testing.T) {
	TcpClientShort()
}

func TestTcpServerHB(t *testing.T) {
	TcpServerHB()
}
func TestTcpClientHB(t *testing.T) {
	TcpClientHB()
}

// 连接池测试
func TestTcpServerPool(t *testing.T) {
	TcpServerPool()
}
func TestTcpClientPool(t *testing.T) {
	TcpClientPool()
}

// 粘包现象测试
func TestTcpServerSticky(t *testing.T) {
	TcpServerSticky()
}
func TestTcpClientSticky(t *testing.T) {
	TcpClientSticky()
}
