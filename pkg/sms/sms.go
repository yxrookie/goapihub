package sms

import (
	"goapihub/pkg/config"
	"sync"
)

type Message struct {
	Template string
	Data map[string]string

	Content string
}

// SMS 使我们发送短信的操作类
type SMS struct {
	Driver Driver
}

// once 单例模式
var once sync.Once

// internalSMS 内部使用的 SMS 对象
var internalSMS *SMS

// NewSMS 单例模式获取
func NewSMS() *SMS {
    once.Do(func() {
        internalSMS = &SMS{
            Driver: &Aliyun{},
        }
    })

    return internalSMS
}

func (sms *SMS) Send(phone string, message Message) bool {
    return sms.Driver.Send(phone, message, config.GetStringMapString("sms.aliyun"))
}