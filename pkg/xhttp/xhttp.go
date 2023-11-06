package xhttp

import (
	"birdProtection/pkg/errcode"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"
	"mime/multipart"
	"net/http"
)

const (
	halfShowLen            = 100
	defaultMultipartMemory = 32 << 20 // 32 MB
)

// Response 业务通用响应体
type Response struct {
	//TraceId string      `json:"trace_id" example:"a1b2c3d4e5f6g7h8" extensions:"x-order=000"` // 链路追踪id
	Code uint32      `json:"code" example:"200" extensions:"x-order=001"` // 状态码
	Msg  string      `json:"msg" example:"OK" extensions:"x-order=002"`   // 消息
	Data interface{} `json:"data" extensions:"x-order=003"`               // 数据
}

var m = make(map[string]any)

// Error 错误响应返回
func Error(w http.ResponseWriter, r *http.Request, err error) {
	ctx := r.Context()
	logx.WithContext(ctx).Errorf("request_handle_err, url:%s err: %+v", r.RequestURI, err)

	e := errcode.ParseErr(err)
	httpx.WriteJson(w, http.StatusOK, &Response{
		Code: e.Code(),
		Msg:  e.Error(),
		Data: m,
	})
}

// OkJson 成功json响应返回
func OkJson(w http.ResponseWriter, r *http.Request, v interface{}) {
	httpx.WriteJson(w, http.StatusOK, &Response{
		Code: errcode.Success,
		Msg:  errcode.SuccessMsq,
		Data: v,
	})
}

// Parse 请求体解析
func Parse(r *http.Request, v interface{}) error {
	if err := httpx.Parse(r, v); err != nil {
		logx.WithContext(r.Context()).Errorf("request parse err, err: %s", err.Error())
		return errcode.NewSingleErr(errcode.InvalidParams, fmt.Sprintf("请求参数错误: %s", err.Error()))
	}
	return nil
}

// FromFile 请求表单文件获取
func FromFile(r *http.Request, name string) (*multipart.FileHeader, error) {
	if r.MultipartForm == nil {
		if err := r.ParseMultipartForm(defaultMultipartMemory); err != nil {
			return nil, err
		}
	}
	f, fh, err := r.FormFile(name)
	if err != nil {
		if err == http.ErrMissingFile {
			return nil, errcode.ErrInvalidParams
		}
		return nil, err
	}
	f.Close()
	return fh, nil
}
