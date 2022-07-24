package demohdlr

import (
	"context"
	"lambda-go/internal/dispatcher"
)

type DemoHdlr struct {
}

func NewDemoHdlr() *DemoHdlr {
	var dh DemoHdlr
	return &dh
}

var _ dispatcher.Handler = (*DemoHdlr)(nil)

func (dh *DemoHdlr) HandleEvent(ctx context.Context, event dispatcher.Event) error {
	return nil
}
