package listener

import (
	"encoding/json"
	"fmt"
	"github.com/gookit/event"
	"gohub/app/models/login_log"
	"gohub/pkg/logger"
)

type LoginListener struct {
}

// Handle 处理登录事件
func (l *LoginListener) Handle(e event.Event) error {
	m := e.Data()
	jsonData, err := json.Marshal(m)
	logger.LogIf(err)

	loginLog := login_log.LoginLog{}
	// 将 JSON 数据反序列化到 LoginLog 结构体中
	err = json.Unmarshal(jsonData, &loginLog)
	loginLog.Create()
	if loginLog.ID > 0 {
		fmt.Println("创建成功")
	} else {
		fmt.Println("创建失败")
	}
	return nil
}
