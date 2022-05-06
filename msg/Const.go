package msg

type Category int

const (
	DEFAULT Category = iota
	VIDEO
	IMAGE
	AUDIO
	FILE
	TEXT
)

type MsgRecord struct {
	id uint8
	// 调用方消息ID
	msgId uint8
	// 消息内容
	msgContent string
	// 透传字段
	msgStr string
}

var Factory = map[Category]IStrategy{}
