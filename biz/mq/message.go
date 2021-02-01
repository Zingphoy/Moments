package mq

/*
 Define MQ message
*/

type Message struct {
	MsgType  uint // type of message, the most import sign
	Aid      int64
	Uid      int32
	Desc     string // description
	NeedSafe bool   // set true if you want the message to be executed safely, otherwise false
}

const (
	EXPAND_TIMELINE_ADD    = 1
	EXPAND_TIMELINE_DELETE = 2
)

const (
	TOPIC = "Moments"
)
