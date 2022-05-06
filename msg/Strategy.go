package msg

type IStrategy interface {
	// GetCategory returns the category of the strategy.
	getType() Category
	// 消息处理
	doHandle(msgList []*MsgRecord)
}

type InitStrategy struct {
	IStrategy
}

func (i *InitStrategy) execute(msgList []*MsgRecord) {
	impl := i.IStrategy

	if strategy, ok := Factory[impl.getType()]; ok {

		strategy.doHandle(msgList)

	}

}

func (i *InitStrategy) doHandle(msgList []*MsgRecord, predicate Predicate) {
	if msgList == nil {
		return
	}

	for _, msg := range msgList {
		// 前置操作
		if msg == nil {
			continue
		}

		if !predicate(msg) {

		}

		// 后置操作

	}
}
