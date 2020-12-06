package src

import (
	"context"
	"time"
)

type StringOperation struct {
	ctx context.Context
}

func NewStringOperation() *StringOperation {
	return &StringOperation{ctx: context.Background()}
}

func (this *StringOperation) Get(key string) *StringResult {
	return NewStringResult(Redis().Get(this.ctx, key).Result())
}

func (this *StringOperation) MGet(key ...string) *SliceResult {
	return NewSliceResult(Redis().MGet(this.ctx, key...).Result())
}

func (this *StringOperation) Set(key string, value interface{}, attr ...*OperateAttr) *InterfaceResult {
	exp := OperateAttrs(attr).Find(ATTR_EXPIRE).Unwrap_Or(time.Second * 0).(time.Duration)
	nx := OperateAttrs(attr).Find(ATTR_NX).Unwrap_Or(nil)
	if nx != nil {
		return NewInterfaceResult(Redis().SetNX(this.ctx, key, value, exp).Result())
	}

	xx := OperateAttrs(attr).Find(ATTR_XX).Unwrap_Or(nil)
	if xx != nil {
		return NewInterfaceResult(Redis().SetXX(this.ctx, key, value, exp).Result())
	}

	return NewInterfaceResult(Redis().Set(this.ctx, key, value, exp).Result())
}
