package mq

/*
参考
https://github.com/apache/rocketmq-client-go/blob/master/examples/producer/transaction/main.go
https://rocketmq.apache.org/docs/transaction-example/
*/

import (
	"context"

	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/primitive"
)

/*
 消费消息的时候进行扩散写，这个过程要保证事务性
*/

func InitExpander() {

}

func callback(ctx context.Context, msgs ...*primitive.MessageExt) (consumer.ConsumeResult, error) {
	// todo: subscribe callback
	return consumer.ConsumeSuccess, nil
}
