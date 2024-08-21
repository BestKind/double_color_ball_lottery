package server

import (
	"context"
	"double_color_ball_lottery/backend/routes"
	"fmt"
	"net/http"
	"time"
)

type HTTPServer struct {
	ctx context.Context
	srv *http.Server

	ctxCancel context.CancelFunc

	running bool
}

// NewHTTPServer 创建 HTTP server
func NewHTTPServer(port int) *HTTPServer {
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", port),
		Handler:      routes.InitRouter(),
		ReadTimeout:  120 * time.Second,
		WriteTimeout: 120 * time.Second,
	}
	ctx, cancel := context.WithCancel(context.TODO())
	return &HTTPServer{
		ctx:       ctx,
		srv:       srv,
		ctxCancel: cancel,
		running:   true,
	}
}

// Start 启动 HTTP 服务
func (h *HTTPServer) Start() {
	fmt.Printf("http server info:%s ", h.srv.Addr)
	go func(h *HTTPServer) {
		err := h.srv.ListenAndServe()
		if err != nil {
			fmt.Printf("http listenAndServe error : %v", err)
			return
		}
	}(h)
}

// Stop 关闭 HTTP 服务
func (h *HTTPServer) Stop() {
	go func() {
		time.Sleep(2 * time.Second)
		fmt.Println("shutdown")
		h.ctxCancel()
	}()
	err := h.srv.Shutdown(h.ctx)
	if err != nil {
		fmt.Printf("http server shutdown error : %v", err)
	}
}
