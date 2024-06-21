package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/shuyi-tangerine/money/gin"
	"github.com/shuyi-tangerine/money/thrift"
	"os"
)

func Usage() {
	fmt.Fprint(os.Stderr, "Usage of ", os.Args[0], ":\n")
	flag.PrintDefaults()
	fmt.Fprint(os.Stderr, "\n")
}

func main() {
	flag.Usage = Usage
	server := flag.Bool("server", false, "Run server")
	protocol := flag.String("P", "binary", "Specify the protocol (binary, compact, json, simplejson)")
	framed := flag.Bool("framed", false, "Use framed transport")
	buffered := flag.Bool("buffered", false, "Use buffered transport")
	addr := flag.String("addr", "localhost:9090", "Address to listen to rpc server")
	secure := flag.Bool("secure", false, "Use tls secure transport")
	webAddr := flag.String("web_addr", "localhost:8080", "Address to listen to web server")

	flag.Parse()

	thriftServer := thrift.NewServer(*server, *protocol, *buffered, *framed, *addr, *secure)

	// 不是服务则不启动，直接执行就行
	if !*server {
		err := thriftServer.Start(context.Background())
		if err != nil {
			fmt.Println("run error", err)
			return
		}
		return
	}

	thriftServer.AsyncStart(context.Background()) // 异步启动
	ginServer := gin.NewServer(*webAddr)
	ginServer.AsyncStart(context.Background()) // 异步启动
	select {
	case err := <-thriftServer.ErrorC():
		fmt.Println("thriftServer Start error", err)
		panic(err)
	case err := <-ginServer.ErrorC():
		fmt.Println("ginServer Start error", err)
		panic(err)
	}
}
