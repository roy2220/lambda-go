package dispatcher

import (
	"context"
	"fmt"

	"github.com/prometheus/client_golang/prometheus/push"
)

// Dispatcher 为事件分发器。
type Dispatcher struct {
	pusher            *push.Pusher
	eventType2Handler map[EventType]Handler
}

// New 创建一个事件分发器。
func New(pusher *push.Pusher) *Dispatcher {
	var d Dispatcher
	d.pusher = pusher
	d.eventType2Handler = make(map[EventType]Handler)
	return &d
}

// RegisterHandler 把事件处理器注册到分发器。如果 eventType 对应的事件处理器已经存在，返回错误。
func (d *Dispatcher) RegisterHandler(eventType EventType, handler Handler) error {
	if _, ok := d.eventType2Handler[eventType]; ok {
		return fmt.Errorf("dispatcher: handler already exists; eventType=%q", eventType)
	}
	d.eventType2Handler[eventType] = handler
	return nil
}

// HandleEvent 把事件分发到对应的处理器。如果找不到 eventType 对应的事件处理器，返回错误。
func (d *Dispatcher) HandleEvent(ctx context.Context, event Event) error {
	handler, ok := d.eventType2Handler[event.Type]
	if !ok {
		return fmt.Errorf("dispatcher: handler not found; eventType=%q", event.Type)
	}
	err := handler.HandleEvent(ctx, event)
	// TODO : pusher
	return err
}

// Handler 为事件处理器的接口类型。
type Handler interface {
	HandleEvent(ctx context.Context, event Event) (err error)
}

// Event 为事件信息。
type Event struct {
	Type EventType
	Data EventData
}

// EventType 为事件类型。
type EventType = string

// EventData 为事件数据。
type EventData = string
