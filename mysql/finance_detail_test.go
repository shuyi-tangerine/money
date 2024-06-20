package mysql

import (
	"context"
	"fmt"
	"github.com/shuyi-tangerine/money/top"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestFinanceDetailDao_Select(t *testing.T) {
	dao, err := NewFinanceDetailDao()
	if !assert.Nil(t, err) {
		return
	}

	financeDetailPOs, err := dao.Select(context.Background(), &top.FinanceDetailPO{
		ID:              0,
		FinanceDetailID: 1,
	})
	if !assert.Nil(t, err) {
		return
	}
	fmt.Println(financeDetailPOs)
}

func TestFinanceDetailDao_Insert(t *testing.T) {
	dao, err := NewFinanceDetailDao()
	if !assert.Nil(t, err) {
		return
	}

	extra := "{}"

	financeDetailPOs, err := dao.Insert(context.Background(), &top.FinanceDetailPO{
		FinanceDetailID: 2,
		AppID:           3,
		Amount:          3,
		OperatedType:    5,
		OperatedAt:      time.Now(),
		OperatedBy:      "chenshuyi01",
		Extra:           &extra,
		CreatedBy:       "chenshuyi02",
		UpdatedBy:       "chenshuyi03",
	})
	if !assert.Nil(t, err) {
		return
	}
	fmt.Println(financeDetailPOs)
}

func TestFinanceDetailDao_Update(t *testing.T) {
	dao, err := NewFinanceDetailDao()
	if !assert.Nil(t, err) {
		return
	}

	extra := "{}"

	affectedRows, err := dao.Update(context.Background(), &top.FinanceDetailPO{
		FinanceDetailID: 2,
		AppID:           1,
		Amount:          2,
		OperatedType:    3,
		OperatedAt:      time.Now(),
		OperatedBy:      "chenshuyi02",
		Extra:           &extra,
		CreatedBy:       "chenshuyi03",
		UpdatedBy:       "chenshuyi01",
	})
	if !assert.Nil(t, err) {
		return
	}
	fmt.Println(affectedRows)
}
