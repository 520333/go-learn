package gorm

import (
	"context"
	"testing"
	"time"
)

func TestContextTimeOutCancel(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*20)
	defer cancel()
	ContextTimeOutCancel(ctx)
}
