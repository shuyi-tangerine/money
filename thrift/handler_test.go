package thrift

import (
	"context"
	"fmt"
	"github.com/shuyi-tangerine/money/gen-go/tangerine/money"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestMoneyHandler_AddFinanceDetail(t *testing.T) {
	handler, err := NewMoneyHandler()
	if !assert.Nil(t, err) {
		return
	}

	resp, err := handler.AddFinanceDetail(context.Background(), &money.AddFinanceDetailRequest{
		AppID:           1,
		Amount:          2,
		OperatedType:    3,
		OperatedAt:      4,
		OperatedBy:      "chenshuyi",
		Extra:           nil,
		CreatedBy:       "chenshuyi",
		FinanceDetailID: 0,
		Base:            nil,
	})
	if !assert.Nil(t, err) {
		return
	}
	fmt.Println(resp)
}

func TestMoneyHandler_ListFinanceDetail(t *testing.T) {
	handler, err := NewMoneyHandler()
	if !assert.Nil(t, err) {
		return
	}

	resp, err := handler.ListFinanceDetail(context.Background(), &money.ListFinanceDetailRequest{
		AppID: 0,
		OperatedAt: &money.TimeRange{
			S: 0,
			E: time.Now().Unix(),
		},
		Offset: 0,
		Limit:  10,
		Base:   nil,
	})
	if !assert.Nil(t, err) {
		return
	}
	fmt.Println(resp)
}
