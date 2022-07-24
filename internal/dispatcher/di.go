package dispatcher

import (
	"context"

	"github.com/go-tk/di"
	"github.com/prometheus/client_golang/prometheus/push"
)

// Functions 返回DI函数
func Functions() (functions []di.Function) {
	functions = append(functions, provideDispatcher())
	return
}

func provideDispatcher() di.Function {
	var (
		pusher *push.Pusher

		dispatcher *Dispatcher
	)
	body := func(ctx context.Context) error {
		dispatcher = New(pusher)
		return nil
	}
	return di.Function{
		Tag: di.FullFunctionName(provideDispatcher),
		Arguments: []di.Argument{
			{InValueID: "prometheus-pusher", InValuePtr: &pusher},
		},
		Results: []di.Result{
			{OutValueID: "dispatcher", OutValuePtr: &dispatcher},
		},
		Body: body,
	}
}
