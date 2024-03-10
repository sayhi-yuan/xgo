package core

type ListenerInterface interface {
	Handle(ctx *Context, data interface{})
}

type EventInterface interface {
	Dispatch(ctx *Context, data interface{})
	SyncDispatch(ctx *Context, data interface{})
	RegisterListener(listeners ...ListenerInterface)
}

type BaseEvent struct {
	Name      string `desc:"事件名称"`
	Listeners []ListenerInterface
}

// 注册监听者
func (event *BaseEvent) RegisterListener(listeners ...ListenerInterface) {
	event.Listeners = append(event.Listeners, listeners...)
}

// 并发事件
func (event BaseEvent) Dispatch(ctx *Context, data interface{}) {
	for _, listener := range event.Listeners {
		go listener.Handle(ctx, data)
	}
}

// 同步事件
func (event BaseEvent) SyncDispatch(ctx *Context, data interface{}) {
	for _, listener := range event.Listeners {
		listener.Handle(ctx, data)
	}
}
