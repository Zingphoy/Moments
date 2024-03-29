package mq

/*
参考
https://github.com/apache/rocketmq-client-go/blob/master/examples/producer/transaction/main.go
https://rocketmq.apache.org/docs/transaction-example/
*/

import (
	"Moments/biz/friend"
	"Moments/biz/timeline"
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

// InitExpander initialize the manager which will take care of go routine that consumes messages.
func InitExpander() {
	initManager()
}

func initManager() {
	c := make(chan int, 0)
	go Expand(c)

	// keep Expander alive
	go func(chan int) {
		for {
			select {
			case <-c:
				go Expand(c)
			}
		}
	}(c)
}

// Expand consume messages in MQ, expand data to other collections
func Expand(c chan int) {
	err := RunMessageConsumer(TOPIC, callback)
	if err != nil {
		log.Error(nil, fmt.Sprintf("consume message failed %v", err))
		c <- 1
	}
}

func callback(ctx context.Context, msgs ...*primitive.MessageExt) (consumer.ConsumeResult, error) {
	for _, msg := range msgs {
		// only one message at a time
		m := Message{}
		err := json.Unmarshal(msg.Message.Body, &m)
		if err != nil {
			log.Error(nil, "message deserialize failed,", err.Error())
			return consumer.ConsumeSuccess, err
		}

		aid, uid := m.Aid, m.Uid
		if aid <= 0 || uid <= 0 {
			log.Error(nil, fmt.Sprintf("aid=%d, uid=%d", aid, uid))
			return consumer.ConsumeSuccess, errors.New("aid or uid is not valided")
		}

		// todo 获取好友列表，然后遍历调用 AppendTimeline
		friendSrv := friend.FriendService{}
		err = friendSrv.GetFriendList(nil)
		if err != nil {
			log.Error(nil, "get frined list failed,", err.Error())
			return consumer.ConsumeRetryLater, err
		}
		for _, fuid := range friendSrv.Data.FriendList {
			tlSrv := timeline.NewTimelineService(&timeline.Timeline{}, &timeline.TimelineModelImpl{})
			switch m.MsgType {
			case EXPAND_TIMELINE_ADD:
				err = tlSrv.AppendArticleIntoTimeline(nil, fuid, aid)
				log.Info(nil, aid, fuid)
			case EXPAND_TIMELINE_DELETE:
				log.Info(nil, fuid)
				//err = model.RemoveTimeline(fuid, aid)
			}
			if err != nil {
				log.Error(nil, fmt.Sprintf("expand timeline failed, uid=%d ", uid))
				return consumer.ConsumeRetryLater, err
			}
		}

	}
	return consumer.ConsumeSuccess, nil
}
