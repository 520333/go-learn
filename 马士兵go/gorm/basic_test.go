package gorm

import "testing"

func TestBasicUsage(t *testing.T) {
	BasicUsage()
}

func TestCreate(t *testing.T) {
	Create()
}

func TestRetrieve(t *testing.T) {
	Retrieve(2)
}

func TestUpdate(t *testing.T) {
	Update()
}

func TestDelete(t *testing.T) {
	Delete(1)
}
func TestDebug(t *testing.T) {
	Debug()
}

func TestLog(t *testing.T) {
	Log()
}
