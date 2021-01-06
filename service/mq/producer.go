package mq

import (
	"Moments/pkg/log"
	"encoding/json"
)

/*
 Producer is a service to send message into MQ
*/

// 这里是不是只需要吧消息send过去就好了？那看起来直接在原service中调用SendMessage不就行了吗？这个文件看起来没有存在的必要

// ExpandTimeline a encapsulation of sending message to MQ
func (m *Message) ExpandTimeline() error {
	var err error
	switch m.MsgType {
	case 1:
		err = SendMsg(m)
	case 2:
		err = SendMsg(m)
	}
	if err != nil {
		log.Error("send message to MQ failed,", err.Error())
		return err
	}
	return nil
}

// SendMsg a wrapper of sending message determine by Message.NeedSafe
func SendMsg(m *Message) error {
	msg, err := json.Marshal(&m)
	if err != nil {
		log.Error("serialize message failed,", err.Error())
		return err
	}

	switch m.NeedSafe {
	case true:
		err = SendMessage(TOPIC, msg)
	case false:
		err = SendMessage(TOPIC, msg)
	}
	return err
}
