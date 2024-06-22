package gin

import (
	"context"
	"fmt"
	"github.com/shuyi-tangerine/money/top"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Server struct {
	addr string // 启动ip:port，空则默认是 localhost:8080
	c    chan error
}

func NewServer(addr string) top.Server {
	return &Server{
		addr: addr,
		c:    make(chan error),
	}
}

func (m *Server) IsBlock(ctx context.Context) (isBlock bool) {
	return true
}

func (m *Server) Start(ctx context.Context) (err error) {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	return r.Run(m.addr)
}

func (m *Server) AsyncStart(ctx context.Context) {
	go func() {
		err := m.Start(ctx)
		if err != nil {
			fmt.Println("[AsyncStart] http Start panic", err)
			m.c <- err
		}
	}()
}

func (m *Server) ErrorC() (c chan error) {
	return m.c
}
