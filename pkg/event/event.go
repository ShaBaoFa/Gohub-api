package event

import (
	event "github.com/gookit/event"
	"gohub/pkg/logger"
	"sync"
)

// EventManager 负责管理事件和监听器的注册
type EventService struct {
	Manager *event.Manager
}

var once sync.Once
var Event *EventService

func New(name string) *EventService {
	// 初始化Event
	eventS := &EventService{}
	eventS.Manager = event.NewM(name)
	return eventS
}

func Setup(name string) {
	once.Do(func() {
		Event = New(name)
	})
}

// SetupListeners 设置多个事件监听器
func SetupListeners(listeners map[string]event.Listener) {
	// 注册监听器
	for name, item := range listeners {
		Event.Manager.On(name, item)
	}
}

// FireC 触发事件
func FireC(eventName string, data event.M) {
	// 触发事件x
	Event.Manager.FireC(eventName, data)
}

func CloseWait() {
	err := Event.Manager.CloseWait()
	logger.LogIf(err)
}
