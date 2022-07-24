package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"lambda-go/internal/dispatcher"
	"lambda-go/internal/handler/demohdlr"
	"lambda-go/plugin/prometheusplug"

	"github.com/go-tk/configset"
	"github.com/go-tk/di"
)

func main() {
	// 加载配置
	configset.MustLoad("./etc")

	// 添加DI函数
	var program di.Program
	program.AddFunctions(prometheusplug.Functions()...)
	program.AddFunctions(dispatcher.Functions()...)
	program.AddFunctions(demohdlr.Functions()...)
	program.AddFunctions(handleEvent())

	// 运行DI函数
	defer program.Clean()
	program.MustRun(context.Background())
}

func handleEvent() di.Function {
	var (
		dispatcher1 *dispatcher.Dispatcher
	)
	body := func(ctx context.Context) error {
		ctx, cancel := context.WithTimeout(ctx, 20*time.Minute)
		defer cancel()

		args := os.Args[1:]
		if len(args) != 2 {
			return fmt.Errorf("lambda-go: invalid arguments: args=%#v", args)
		}
		event := dispatcher.Event{
			Type: args[0],
			Data: args[1],
		}

		if err := dispatcher1.HandleEvent(ctx, event); err != nil {
			return fmt.Errorf("handle event; event=%#v: %w", event, err)
		}

		return nil
	}
	return di.Function{
		Tag: di.FullFunctionName(handleEvent),
		Arguments: []di.Argument{
			{InValueID: "dispatcher", InValuePtr: &dispatcher1},
		},
		Body: body,
	}
}
