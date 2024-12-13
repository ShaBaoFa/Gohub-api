package bootstrap

import (
	"github.com/gookit/event"
	"gohub/listener"
	"gohub/pkg/config"
	eventpkg "gohub/pkg/event"
)

func SetupEventBus() {
	eventpkg.Setup(config.Get("app.name"))
	// 创建一个 map，用于存储事件名称到监听器的映射
	listeners := make(map[string]event.Listener)
	listeners["user.login"] = &listener.LoginListener{}
	eventpkg.SetupListeners(listeners)
}
