package msg

type IStrategy interface {

	// GetCategory returns the category of the strategy.
	getType() Category
	// 消息处理
	doHandle(msgList []*MsgRecord)
}

type AbstractStrategy struct {
	impl IStrategy
}

func (a *AbstractStrategy) execute(msgList []*MsgRecord) {
	impl := a.impl

	if strategy, ok := Factory[impl.getType()]; ok {

		strategy.doHandle(msgList)

	}

}

func (a *AbstractStrategy) doHandle(msgList []*MsgRecord, predicate Predicate) {
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
