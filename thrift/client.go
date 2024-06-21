package thrift

import (
	"context"
	"fmt"
	"github.com/shuyi-tangerine/money/gen-go/base"
	"github.com/shuyi-tangerine/money/gen-go/tangerine/money"

	"github.com/apache/thrift/lib/go/thrift"
)

var defaultCtx = context.Background()

func handleClient(client money.MoneyHandler) (err error) {
	res, err := client.ListFinanceDetail(defaultCtx, &money.ListFinanceDetailRequest{
		Base: &base.RPCRequest{}})
	if err != nil {
		return err
	}
	fmt.Println("res ==>", res)
	return
}

func runClient(transportFactory thrift.TTransportFactory, protocolFactory thrift.TProtocolFactory, addr string, secure bool, cfg *thrift.TConfiguration) error {
	var transport thrift.TTransport
	if secure {
		transport = thrift.NewTSSLSocketConf(addr, cfg)
	} else {
		transport = thrift.NewTSocketConf(addr, cfg)
	}
	transport, err := transportFactory.GetTransport(transport)
	if err != nil {
		return err
	}
	defer transport.Close()
	if err := transport.Open(); err != nil {
		return err
	}
	iprot := protocolFactory.GetProtocol(transport)
	oprot := protocolFactory.GetProtocol(transport)
	return handleClient(money.NewMoneyHandlerClient(thrift.NewTStandardClient(iprot, oprot)))
}
