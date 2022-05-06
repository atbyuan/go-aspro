package msg

func init() {
	strategy := TextStrategy{}
	Factory[strategy.getType()] = strategy
}

type TextStrategy struct {
	AbstractStrategy
}

func (t TextStrategy) getType() Category {
	return TEXT
}

func (t TextStrategy) doHandle(msgList []*MsgRecord) {

	t.AbstractStrategy.doHandle(msgList, func(msg *MsgRecord) bool {

		return true
	})

}
