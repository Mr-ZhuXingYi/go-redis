package src

import (
	"fmt"
	"time"
)

type empty struct{}

const (
	ATTR_EXPIRE = "expr" //过期时间
	ATTR_NX     = "nx"   // setnx
	ATTR_XX     = "xx"   // setnx
)

type OperateAttr struct {
	Name  string
	Value interface{}
}

type OperateAttrs []*OperateAttr

func (this OperateAttrs) Find(name string) *InterfaceResult {
	for _, attr := range this {
		if attr.Name == name {
			return NewInterfaceResult(attr.Value, nil)
		}
	}
	return NewInterfaceResult(nil, fmt.Errorf("OperationAttrs found error:%s", name))
}

func WithExpire(t time.Duration) *OperateAttr {
	return &OperateAttr{Name: ATTR_EXPIRE, Value: t}
}

func WithNx() *OperateAttr {
	return &OperateAttr{Name: ATTR_NX, Value: empty{}}
}
