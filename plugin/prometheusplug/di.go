package prometheusplug

import (
	"context"
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/go-tk/configset"
	"github.com/go-tk/di"
	"github.com/prometheus/client_golang/prometheus/push"
)

// Functions 返回DI函数
func Functions() (functions []di.Function) {
	var config config
	configset.MustReadValue("prometheusplug", &config)
	if err := validator.New().Struct(config); err != nil {
		panic(fmt.Sprintf("validate config: %v", err))
	}

	for i := range config.Pushers {
		pusherConfig := &config.Pushers[i]
		functions = append(functions, providePusher(pusherConfig))
	}
	return
}

type config struct {
	Pushers []pusherConfig `validate:"dive"`
}

type pusherConfig struct {
	ID             string `validate:"required"`
	URL            string `validate:"required,url"`
	Job            string `validate:"required"`
	GroupingLabels map[string]string
}

func providePusher(pusherConfig *pusherConfig) di.Function {
	var (
		pusher *push.Pusher
	)
	body := func(context.Context) error {
		pusher = push.New(pusherConfig.URL, pusherConfig.Job)
		for key, value := range pusherConfig.GroupingLabels {
			pusher.Grouping(key, value)
		}
		return nil
	}
	return di.Function{
		Tag: di.FullFunctionName(providePusher),
		Results: []di.Result{
			{OutValueID: pusherConfig.ID, OutValuePtr: &pusher},
		},
		Body: body,
	}
}
