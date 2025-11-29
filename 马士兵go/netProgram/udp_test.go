package netProgram

import "testing"

func TestUdpServerBasic(t *testing.T) {
	UdpServerBasic()
}
func TestUdpClientBasic(t *testing.T) {
	UdpClientBasic()
}

func TestUdpServerConnect(t *testing.T) {
	UdpServerConnect()
}
func TestUdpClientConnect(t *testing.T) {
	UdpClientConnect()
}

// 对等连接
func TestUdpServerPeer(t *testing.T) {
	UdpServerPeer()
}
func TestUdpClientPeer(t *testing.T) {
	UdpClientPeer()
}
