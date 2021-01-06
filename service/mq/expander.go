package mq

/*
参考
https://github.com/apache/rocketmq-client-go/blob/master/examples/producer/transaction/main.go
https://rocketmq.apache.org/docs/transaction-example/
*/

import (
	"Moments/model"
	"Moments/pkg/log"
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/primitive"
)

/*
 Expander is a module which consumes message from MQ. According to message type, expander will expand data into other collections.
 It's a transaction to expand data to ensure the correctness of data.
*/

// fork一个子进程乱来进行消息消费，这里只负责管理子进程。
// 检查子进程的健康度，挂了要重启（子进程挂了是不是会丢消息？）

// go协程来消费消息，外围管理这个协程的状态即可，
// InitExpander initialize the manager which will take care of go routine that consumes messages.
func InitExpander() error {
	err := initManager()
	if err != nil {
		log.Fatal(err.Error())
	}
	return err
}

// todo, how to manager go routine
func initManager() error {
	return nil
}

// Expand consume messages in MQ, expand data to other collections
func Expand() error {
	err := ConsumeMessage(TOPIC, callback)
	if err != nil {
		log.Error("consume message failed,", err.Error())
	}
	return err
}

func callback(ctx context.Context, msgs ...*primitive.MessageExt) (consumer.ConsumeResult, error) {
	// todo: subscribe callback
	msg := Message{}
	log.Info("message length: ", len(msgs))

	// only one message at a time
	err := json.Unmarshal(msgs[0].Message.Body, &msg)
	if err != nil {
		log.Error("message deserialize failed,", err.Error())
		return consumer.ConsumeRetryLater, err
	}

	aid, uid := msg.Aid, msg.Uid
	if aid <= 0 || uid <= 0 {
		log.Error(fmt.Sprintf("aid: %d, uid: %d", aid, uid))
		return consumer.ConsumeSuccess, errors.New("aid or uid is not valided")
	}

	switch msg.MsgType {
	case 1:
		err = model.AppendTimeline(uid, aid)
	case 2:
		err = model.RemoveTimeline(uid, aid)
	}
	if err != nil {
		log.Error("expand timeline failed ")
		return consumer.ConsumeRetryLater, err
	}
	return consumer.ConsumeSuccess, nil
}
