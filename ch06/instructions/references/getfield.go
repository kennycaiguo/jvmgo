package references

import (
	"jvmgo/ch06/instructions/base"
	"jvmgo/ch06/rtdata"
	"jvmgo/ch06/rtdata/heap"
)

type GET_FIELD struct{ base.Index16Instruction }

func (self *GET_FIELD) Execute(frame *rtdata.Frame) {
	cp := frame.Method().Class().ConstantPool()
	fieldRef := cp.GetConstant(self.Index).(*heap.FieldRef)
	field := fieldRef.ResolvedField()
	if field.IsStatic() {
		panic("java.lang.IncompatibleClassChangeError")
	}

	stack := frame.OperandStack()
	ref := stack.PopRef()
	if ref == nil {
		panic("java.lang.NullPointerException")
	}

	descriptor := field.Descripter()
	slots := ref.Fields()
	slotId := field.SlotId()

	switch descriptor[0] {
	case 'Z', 'B', 'C', 'S', 'I':
		stack.PushInt(slots.GetInt(slotId))
	case 'F':
		stack.PushFloat(slots.GetFloat(slotId))
	case 'D':
		stack.PushDouble(slots.GetDouble(slotId))
	case 'J':
		stack.PushLong(slots.GetLong(slotId))
	case 'L', '[':
		stack.PushRef(slots.GetRef(slotId))
	}
}
