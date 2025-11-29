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

// 组播测试
func TestUdpReceiverMultiCast(t *testing.T) {
	UdpReceiverMultiCast()
}
func TestUdpSenderMultiCast(t *testing.T) {
	UdpSenderMultiCast()
}

// 广播测试
func TestUdpReceiverBroadCast(t *testing.T) {
	UdpReceiverBroadCast()
}
func TestUdpSenderBroadCast(t *testing.T) {
	UdpSenderBroadCast()
}
