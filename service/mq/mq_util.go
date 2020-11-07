package mq

import (
	"Moments/pkg/log"
	"context"
	"fmt"
	"sync"
	"sync/atomic"
	"time"

	mq "github.com/apache/rocketmq-client-go/v2"
	//"github.com/apache/rocketmq-client-go/v2/admin"
	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/apache/rocketmq-client-go/v2/producer"
)

/*
 使用事务性消息，保证消息能稳健入队列
*/

const (
	GROUP_NAME    = "moments"
	HOST          = "127.0.0.1"
	PORT          = "8000"
	HOST_AND_PORT = HOST + ":" + PORT
)

var (
	p mq.TransactionProducer = nil
	c mq.PushConsumer        = nil
)

type consumerFunc func(context.Context, ...*primitive.MessageExt) (consumer.ConsumeResult, error)

type ExpanderListener struct {
	localTrans       *sync.Map
	transactionIndex int32
}

func NewExpanderListener() *ExpanderListener {
	return &ExpanderListener{
		localTrans: new(sync.Map),
	}
}

func (l *ExpanderListener) ExecuteLocalTransaction(msg *primitive.Message) primitive.LocalTransactionState {
	nextIndex := atomic.AddInt32(&l.transactionIndex, 1)
	status := nextIndex % 3
	l.localTrans.Store(msg.TransactionId, primitive.LocalTransactionState(status+1))
	log.Info(fmt.Sprintf("nextIndex: %v for transactionID: %v\n", nextIndex, msg.TransactionId))
	//todo: 增加数据库相关操作内容？
	log.Info("do sth here?")
	return primitive.UnknowState
}

func (l *ExpanderListener) CheckLocalTransaction(msg *primitive.MessageExt) primitive.LocalTransactionState {
	log.Info(fmt.Sprintf("%v msg transactionID : %v\n", time.Now(), msg.TransactionId))
	v, existed := l.localTrans.Load(msg.TransactionId)
	if !existed {
		log.Warn(fmt.Sprintf("unknown msg: %v, return Commit", msg))
		return primitive.CommitMessageState
	}
	state := v.(primitive.LocalTransactionState)
	switch state {
	case 1:
		log.Info(fmt.Sprintf("checkLocalTransaction unknow: %v\n", msg))
		return primitive.UnknowState
	case 2:
		log.Info(fmt.Sprintf("checkLocalTransaction COMMIT_MESSAGE: %v\n", msg))
		return primitive.CommitMessageState
	case 3:
		log.Info(fmt.Sprintf("checkLocalTransaction ROLLBACK_MESSAGE: %v\n", msg))
		return primitive.RollbackMessageState
	default:
		log.Info(fmt.Sprintf("checkLocalTransaction default COMMIT_MESSAGE: %v\n", msg))
		return primitive.CommitMessageState
	}
}

// TODO: 配置化
// InitMQ initialize Producer and Consumer of MQ
func InitMQ() {
	p, _ = mq.NewTransactionProducer(
		NewExpanderListener(),
		producer.WithNsResovler(primitive.NewPassthroughResolver([]string{HOST_AND_PORT})),
		producer.WithRetry(5),
	)
	err := p.Start()
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	log.Info("RocketMQ initialize success")

	traceCfg := &primitive.TraceConfig{
		Access:   primitive.Local,
		Resolver: primitive.NewPassthroughResolver([]string{HOST_AND_PORT}),
	}
	c, _ = mq.NewPushConsumer(
		consumer.WithGroupName(GROUP_NAME),
		consumer.WithNsResovler(primitive.NewPassthroughResolver([]string{HOST_AND_PORT})),
		consumer.WithTrace(traceCfg),
	)
}

// SendMessage send message to MQ
func SendMessage(topic string, body []byte) (err error) {
	for i := 1; i <= 2; i++ {
		res, err := p.SendMessageInTransaction(context.Background(), primitive.NewMessage(topic, body))
		if err != nil {
			log.Error(fmt.Sprintf("send mq message error for %d try: %s", i, err.Error()))
		} else {
			log.Info("send mq message success: result=", res.String())
			break
		}
	}
	return
}

// ConsumeMessage consume message from MQ by PULL
func ConsumeMessage(topic string, callback consumerFunc) (err error) {
	err = c.Subscribe(topic, consumer.MessageSelector{}, callback)
	if err != nil {
		log.Error(err.Error())
	}
	err = c.Start()
	if err != nil {
		log.Error(err.Error())
	}
	return
}

// StopMq stop MQ (unnecessary)
func StopMQ() {
	if p != nil {
		if err := p.Shutdown(); err != nil {
			log.Error(err.Error())
		}
	}

	if c != nil {
		if err := c.Shutdown(); err != nil {
			log.Error(err.Error())
		}
	}
}
