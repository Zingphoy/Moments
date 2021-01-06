package mq

import (
	"Moments/model"
	"Moments/pkg/log"
	"encoding/json"
	"fmt"
	"sync"

	"github.com/apache/rocketmq-client-go/v2/primitive"
)

/*
 Producer is a module to wrap local transactions
*/

// ExpandTimeline a encapsulation of sending message to MQ
func (m *Message) ExpandTimeline() error {
	err := sendMsg(m)
	if err != nil {
		log.Error("send message to MQ failed,", err.Error())
	}
	return err
}

// sendMsg a wrapper of sending message determine by Message.NeedSafe
func sendMsg(m *Message) error {
	msg, err := json.Marshal(&m)
	if err != nil {
		log.Error("serialize message failed,", err.Error())
		return err
	}

	// todo need sending to mq safe or not
	switch m.NeedSafe {
	case true:
		err = SendMessage(TOPIC, msg)
	case false:
		err = SendMessage(TOPIC, msg)
	}
	return err
}

/*
 about local transaction below
*/
type TransactionListener struct {
	localTrans *sync.Map
}

func NewTransactionListener() *TransactionListener {
	return &TransactionListener{localTrans: new(sync.Map)}
}

/*
 perform local transaction, using map localTrans to determine the status of message processing
 localTrans 1:Commit 2:Rollback 3:Unknown
*/
func (l *TransactionListener) ExecuteLocalTransaction(msg *primitive.Message) primitive.LocalTransactionState {
	log.Info(fmt.Sprintf("begin local transaction, transactionID: %v\n", msg.TransactionId))
	var body Message
	err := json.Unmarshal(msg.Body, &body)
	if err != nil {
		log.Error(err.Error())
		l.localTrans.Store(msg.TransactionId, primitive.LocalTransactionState(2))
		return primitive.RollbackMessageState
	}

	// local transaction, first append sender's timeline, then append to friends' timeline asynchronously
	switch body.MsgType {
	case EXPAND_TIMELINE_ADD:
		err = model.AppendTimeline(body.Uid, body.Aid)
	case EXPAND_TIMELINE_DELETE:
		// todo not implement yet
		break
	}
	if err != nil {
		log.Error(err.Error())
		l.localTrans.Store(msg.TransactionId, primitive.LocalTransactionState(2))
		return primitive.RollbackMessageState
	}

	l.localTrans.Store(msg.TransactionId, primitive.LocalTransactionState(1))
	return primitive.CommitMessageState
}

// to check the 'Prepared-State' transaction using strategy of TransactionListener
func (l *TransactionListener) CheckLocalTransaction(msg *primitive.MessageExt) primitive.LocalTransactionState {
	log.Info(fmt.Sprintf("checking local transaction, transactionID : %v\n", msg.TransactionId))
	v, existed := l.localTrans.Load(msg.TransactionId)
	if !existed {
		return primitive.CommitMessageState
	}

	state := v.(primitive.LocalTransactionState)
	l.localTrans.Delete(msg.TransactionId)
	switch state {
	case 1:
		log.Info(fmt.Sprintf("checkLocalTransaction COMMITE_MESSAGE: %v\n", msg))
		return primitive.CommitMessageState
	case 2:
		log.Info(fmt.Sprintf("checkLocalTransaction ROLLBACK_MESSAGE: %v\n", msg))
		return primitive.RollbackMessageState
	case 3:
		log.Info(fmt.Sprintf("checkLocalTransaction UNKNOWN_MESSAGE: %v\n", msg))
		return primitive.UnknowState
	default:
		log.Error(fmt.Sprintf("message state not expected, checkLocalTransaction: %v\n", msg))
		return primitive.RollbackMessageState
	}
}
