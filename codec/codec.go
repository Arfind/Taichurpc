package codec

import "io"

// 抽象出数据结构Header
type Header struct {
	ServiceMethod string //服务名和方法名，与Go结构体和方法映射
	Seq           uint64 //请求序号，比如请求ID
	Error         string //错误信息
}

// 抽象出codec接口
type Codec interface {
	io.Closer
	ReadHeader(*Header) error
	ReadBody(interface{}) error
	Writer(*Header, interface{}) error
}

type NewCodecFunc func(io.ReadWriteCloser) Codec

type Type string

const (
	GobType  Type = "application/gob"
	JsonType Type = "application/json"
)

var NewCodecFuncMap map[Type]NewCodecFunc

func init() {
	NewCodecFuncMap = make(map[Type]NewCodecFunc)
	NewCodecFuncMap[GobType] = NewGobCodec
}
