package demohdlr

import (
	"context"
	"lambda-go/internal/dispatcher"

	"github.com/go-tk/di"
)

// Functions 返回DI函数
func Functions() (functions []di.Function) {
	functions = append(functions, registerDemoHdlr())
	return
}

func registerDemoHdlr() di.Function {
	var (
		dispatcher *dispatcher.Dispatcher
		callback   func(context.Context) error
	)
	body := func(context.Context) error {
		demoHdlr := NewDemoHdlr()
		callback = func(context.Context) error {
			dispatcher.RegisterHandler("Demo", demoHdlr)
			return nil
		}
		return nil
	}
	return di.Function{
		Tag: di.FullFunctionName(registerDemoHdlr),
		Hooks: []di.Hook{
			{InValueID: "dispatcher", InValuePtr: &dispatcher, CallbackPtr: &callback},
		},
		Body: body,
	}
}
