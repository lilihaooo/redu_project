package common

import (
	"encoding/json"
)

type SeckillStatus int

const (
	SeckillStatusNormal SeckillStatus = 1
	SeckillStatusStop   SeckillStatus = 0
)

func (t SeckillStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.String())
}

func (t SeckillStatus) String() string {
	var str string
	switch t {
	case SeckillStatusNormal:
		str = "正常"
	case SeckillStatusStop:
		str = "暂停"
	}
	return str
}
