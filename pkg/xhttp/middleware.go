package xhttp

import (
	"encoding/json"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"net/http/httputil"
	"time"
)

// NotAllowedMiddleware 不允许的请求处理中间件
type NotAllowedMiddleware struct{}

// NewNotAllowedMiddleware  不允许的请求处理中间件
func NewNotAllowedMiddleware() *NotAllowedMiddleware {
	return &NotAllowedMiddleware{}
}

// Handler 不允许的请求处理
func (m *NotAllowedMiddleware) Handler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		setHeader(w)
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("405 method not allowed"))
	})
}

// CorsMiddleware 跨域请求处理中间件
type CorsMiddleware struct{}

// NewCorsMiddleware 新建跨域请求处理中间件
func NewCorsMiddleware() *CorsMiddleware {
	return &CorsMiddleware{}
}

// Handle 跨域请求处理
func (m *CorsMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		setHeader(w)

		// 放行所有 OPTIONS 方法
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		// 处理请求
		next(w, r)
	}
}

// Handler 跨域请求处理器
func (m *CorsMiddleware) Handler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		setHeader(w)
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}
	})
}

// setHeader 设置响应头
func setHeader(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, X-CSRF-Token, Authorization, AccessToken, Token, X-Health-Secret, X-Admin-Secret")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS, PATCH")
	w.Header().Set("Access-Control-Expose-Headers", "Content-Length, Content-Type, Access-Control-Allow-Origin, Access-Control-Allow-Headers, X-GW-Error-Code, X-GW-Error-Message")
	w.Header().Set("Access-Control-Allow-Credentials", "true")

	w.Header().Set("Content-Security-Policy", "default-src 'self' blob:; script-src 'self' 'unsafe-inline' https://unpkg.com  'unsafe-eval'; style-src 'self' 'unsafe-inline'; frame-ancestors 'self' http://192.168.110.10:8080 http://192.168.110.61:8888 http://192.168.110.61/weresearch https://chatbot.weresearch.com.cn/")
}

// RecoverMiddleware 恐慌捕获恢复处理中间件
type RecoverMiddleware struct{}

var errMsg, _ = json.Marshal(&Response{
	Code: 50000,
	Msg:  "服务出现了一些问题panic， 请稍候重试",
	Data: struct{}{},
})

// NewRecoverMiddleware 新建恐慌捕获恢复处理中间件
func NewRecoverMiddleware() *RecoverMiddleware {
	return &RecoverMiddleware{}
}

// Handle 恐慌捕获恢复处理
func (m *RecoverMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		dump, err := httputil.DumpRequest(r, true)
		if err != nil {
			w.Header().Set(httpx.ContentType, httpx.JsonContentType)
			w.WriteHeader(http.StatusOK)
			w.Write(errMsg)
			return
		}
		defer func() {
			if causeErr := recover(); causeErr != nil {
				logx.WithContext(r.Context()).Errorf("[running]:\n%s \n%s", causeErr, string(dump))

				panicErr, ok := causeErr.(error)
				if ok {
					Error(w, r, panicErr)
				} else {
					w.Header().Set(httpx.ContentType, httpx.JsonContentType)
					w.WriteHeader(http.StatusOK)
					w.Write(errMsg)
				}
			}
		}()

		next(w, r)
	}
}

// LogMiddleware 全局日志处理中间件
type LogMiddleware struct{}

// NewLogMiddleware 新建全局日志处理中间件
func NewLogMiddleware() *LogMiddleware {
	return &LogMiddleware{}
}

// Handle 全局日志处理中间件处理
func (m *LogMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		dump, err := httputil.DumpRequest(r, true)
		logx.WithContext(r.Context()).Infof("[requestLog]: %s %s %s", time.Now().String(), string(dump), err)
		fmt.Println("[requestLog]:", time.Now().String(), string(dump))

		next(w, r)
	}
}
