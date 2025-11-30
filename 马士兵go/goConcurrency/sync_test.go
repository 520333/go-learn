package goConcurrency

import "testing"

func TestSyncErr(t *testing.T) {
	SyncErr()
}
func TestSyncSyncLock(t *testing.T) {
	SyncLock()
}
func TestSyncMutex(t *testing.T) {
	SyncMutex()
}
func TestSyncLockAndNo(t *testing.T) {
	SyncLockAndNo()
}
func TestSyncRLock(t *testing.T) {
	SyncRLock()
}

func TestSyncMapErr(t *testing.T) {
	SyncMapErr()
}
func TestSyncMap(t *testing.T) {
	SyncMap()
}
func TestSyncMapMethod(t *testing.T) {
	SyncMapMethod()
}

func TestSyncAtomicAdd(t *testing.T) {
	SyncAtomicAdd()
}
func TestSyncAtomicValue(t *testing.T) {
	SyncAtomicValue()
}
func TestSyncPool(t *testing.T) {
	SyncPool()
}
func TestSyncOnce(t *testing.T) {
	SyncOnce()
}
