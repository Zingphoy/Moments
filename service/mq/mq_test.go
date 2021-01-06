package mq

import (
	"Moments/pkg/log"
	"context"
	"encoding/json"
	"os"
	"testing"
	"time"

	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/stretchr/testify/assert"
)

var (
	testTopic = "test_moments"
)

func TestMain(m *testing.M) {
	log.InitLogger(true)
	InitMQ()
	exitCode := m.Run()
	StopMQ()
	os.Exit(exitCode)
}

func TestSendMessage(t *testing.T) {
	msg := Message{
		Aid:      1,
		Uid:      90000,
		Desc:     "unit test",
		NeedSafe: false,
	}
	b, err := json.Marshal(msg)
	assert.Nil(t, err)
	err = SendMessage(testTopic, b)
	if err != nil {
		log.Fatal(err.Error())
	}
	assert.Nil(t, err)
}

func TestConsumeMessage(t *testing.T) {
	err := RunMessageConsumer(testTopic,
		func(ctx context.Context, msg ...*primitive.MessageExt) (consumer.ConsumeResult, error) {
			for _, m := range msg {
				body := Message{}
				_ = json.Unmarshal(m.Body, &body)
				assert.Equal(t, "unit test", body.Desc)
			}
			return consumer.ConsumeSuccess, nil
		})
	time.Sleep(time.Second * 15)
	if err != nil {
		log.Fatal(err.Error())
	}
	assert.Nil(t, err)
}

/*
 todo
 1. 了解rocketmq消息消费机制，消费过的消息会不会删除，重启后又重新消费一遍？
 2. client库不在console打印rocketmq的日志信息
 3. client具体怎么用？看起来并没有按预期发送和消费消息

资料：https://dbaplus.cn/news-73-1123-1.html
*/
