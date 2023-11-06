package errcode

import (
	"google.golang.org/grpc/status"
	"sync"
)

const (
	Success    = uint32(20000)
	SuccessMsq = "ok"

	InvalidParams = uint32(40000)
	// OnlyNotifyCode 该状态码用于客户端不需要操作的情况，例如填写的手机号不正确，客户端见到此状态码，只需要把msg展示出来即可
	OnlyNotifyCode = uint32(40001)
	ServerError    = uint32(50000)
)

// CustomErr 业务错误结构
type CustomErr struct {
	code uint32
	msg  string
}

// Code 业务状态码
func (e *CustomErr) Code() uint32 {
	return e.code
}

// Error 消息
func (e *CustomErr) Error() string {
	return e.msg
}

// NewErr 创建新的业务错误
func NewErr(code uint32, msg string) *CustomErr {
	e := &CustomErr{
		code: code,
		msg:  msg,
	}
	_ = em.Add(e, false)
	return e
}

func NewSingleErr(code uint32, msg string) *CustomErr {
	e := &CustomErr{
		code: code,
		msg:  msg,
	}
	return e
}

func NewOnlyNotifyErr(msg string) *CustomErr {
	e := &CustomErr{
		code: OnlyNotifyCode,
		msg:  msg,
	}
	return e
}

// IsCustomErr 判断是否为业务错误
func IsCustomErr(err error) bool {
	if err == nil {
		return true
	}

	_, ok := err.(*CustomErr)
	return ok
}

// ParseErr 解析业务错误
func ParseErr(err error) *CustomErr {
	if err == nil {
		return NoErr
	}

	if e, ok := err.(*CustomErr); ok {
		return e
	}
	s, _ := status.FromError(err)
	c := uint32(s.Code())
	if c == InvalidParams || c == OnlyNotifyCode {
		e := &CustomErr{
			code: c,
			msg:  s.Message(),
		}
		return e
	}
	if e, ok := em.Get(c); ok {
		return e
	}

	//return NewErr(ServerError, "服务出现了一点问题，请稍候！")
	return NewErr(ServerError, err.Error())
}

var em = newErrManager()

// errManager 业务错误管理器
type errManager struct {
	l sync.RWMutex
	m map[uint32]*CustomErr
}

// newErrManager 新建业务错误管理器
func newErrManager() *errManager {
	return &errManager{
		m: make(map[uint32]*CustomErr),
	}
}

// Get 获取业务错误
func (em *errManager) Get(code uint32) (*CustomErr, bool) {
	em.l.RLock()
	defer em.l.RUnlock()

	e, ok := em.m[code]
	return e, ok
}

// Add 添加业务错误
func (em *errManager) Add(e *CustomErr, mustAdd bool) bool {
	em.l.Lock()
	defer em.l.Unlock()

	if e == nil {
		return false
	}

	_, ok := em.m[e.Code()]
	if mustAdd || !ok {
		em.m[e.Code()] = e
		return true
	}

	return false
}

// GetMap 获取业务错误映射
func (em *errManager) GetMap() map[uint32]CustomErr {
	em.l.RLock()
	defer em.l.RUnlock()

	m := make(map[uint32]CustomErr)
	for c := range em.m {
		if e := em.m[c]; e != nil {
			m[c] = *e
		}
	}

	return m
}
