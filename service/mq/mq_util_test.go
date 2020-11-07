package mq

import (
	"Moments/pkg/log"
	"context"
	"os"
	"testing"

	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/primitive"
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
	msg := []byte("Hello world for testing")
	SendMessage(testTopic, msg)
}

func TestConsumeMessage(t *testing.T) {
	ConsumeMessage(testTopic, func(ctx context.Context, msg ...*primitive.MessageExt) (consumer.ConsumeResult, error) {
		log.Debug("message: ", msg)
		return 0, nil
	})
}
