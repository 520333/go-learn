package gorm

import (
	"testing"
)

func TestMigrate(t *testing.T) {
	Migrate()
}

// 指针类型和非指针类型区别
func TestPointerDiff(t *testing.T) {
	PointerDiff()
}
