package mq

import (
	"Moments/pkg/log"
	"context"
	"fmt"

	mq "github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/apache/rocketmq-client-go/v2/producer"
)

/*
 mq wrapper
*/

const (
	GROUP_NAME    = "moments"
	HOST          = "127.0.0.1"
	PORT          = "6000"
	HOST_AND_PORT = HOST + ":" + PORT
)

var (
	p mq.TransactionProducer = nil
	c mq.PushConsumer        = nil
)

type consumerFunc func(context.Context, ...*primitive.MessageExt) (consumer.ConsumeResult, error)

// TODO: 配置化
// InitMQ initialize Producer and Consumer of MQ
func InitMQ() {
	var err error
	p, err = mq.NewTransactionProducer(
		NewTransactionListener(),
		producer.WithNsResovler(primitive.NewPassthroughResolver([]string{HOST_AND_PORT})),
		producer.WithRetry(3),
	)
	if err != nil {
		panic(err)
	}

	err = p.Start()
	if err != nil {
		panic(err)
	}
	log.Info(nil, "MQ producer initialize success")

	c, err = mq.NewPushConsumer(
		consumer.WithGroupName(GROUP_NAME),
		consumer.WithNsResovler(primitive.NewPassthroughResolver([]string{HOST_AND_PORT})),
		consumer.WithTrace(&primitive.TraceConfig{
			Access:   primitive.Local,
			Resolver: primitive.NewPassthroughResolver([]string{HOST_AND_PORT})},
		),
		consumer.WithConsumeFromWhere(consumer.ConsumeFromLastOffset),
	)
	if err != nil {
		panic(err)
		return
	}
	log.Info(nil, "MQ consumer initialize success")
}

// StopMq stop MQ (unnecessary)
func StopMQ() {
	if p != nil {
		if err := p.Shutdown(); err != nil {
			log.Error(nil, err.Error())
		}
	}
	if c != nil {
		if err := c.Shutdown(); err != nil {
			log.Error(nil, err.Error())
		}
	}
}

// SendMessage send message to MQ
func SendMessage(topic string, body []byte) (err error) {
	var res *primitive.TransactionSendResult
	for i := 0; i < 2; i++ {
		res, err = p.SendMessageInTransaction(context.Background(), primitive.NewMessage(topic, body))
		if err != nil {
			log.Warn(nil, fmt.Sprintf("send mq message error for %d try: %s", i, err.Error()))
		} else {
			log.Info(nil, "send mq message success: result=", res.String())
			break
		}
	}
	return
}

// RunMessageConsumer consume message from MQ by PULL, process with callback function
func RunMessageConsumer(topic string, callback consumerFunc) (err error) {
	err = c.Subscribe(topic, consumer.MessageSelector{}, callback)
	if err != nil {
		log.Error(nil, err.Error())
	}
	err = c.Start()
	if err != nil {
		log.Error(nil, err.Error())
	}
	return
}
