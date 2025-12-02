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

func TestCustomType(t *testing.T) {
	CustomType()
}

func TestIAndCreate(t *testing.T) {
	IAndCreate()
}

func TestServiceCRUD(t *testing.T) {
	ServiceCRUD()
}
func TestPaperCRUD(t *testing.T) {
	PaperCRUD()
}

func TestCustomSerializer(t *testing.T) {
	CustomSerializer()
}
