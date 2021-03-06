package runtime

// 帧
type Frame struct {
	next         *Frame
	localVars    LocalVars
	operateStack *OperateStack
}

func NewFrame(maxLocals, maxStack uint) *Frame {
	return &Frame{
		localVars:newLocalVars(maxLocals),
		operateStack:newOperateStack(maxStack),
	}
}

func (self *Frame) LocalVars() LocalVars {
	return self.localVars
}

func (self *Frame) OperateStack() *OperateStack {
	return self.operateStack
}
