package biz

import (
	"context"
	"github.com/shuyi-tangerine/money/mysql"
	"github.com/shuyi-tangerine/money/top"
	"time"
)

type FinanceDetailService struct {
	FinanceDetailDao *mysql.FinanceDetailDao
}

func NewFinanceDetailService() (financeDetailService *FinanceDetailService, err error) {
	financeDetailDao, err := mysql.NewFinanceDetailDao()
	if err != nil {
		return
	}
	return &FinanceDetailService{
		FinanceDetailDao: financeDetailDao,
	}, nil
}

func (m *FinanceDetailService) GenFinanceDetailIDs(ctx context.Context, amount int64) (ids []int64, err error) {
	for i := 0; i < int(amount); i++ {
		id, err := m.GenFinanceDetailID(ctx)
		if err != nil {
			return nil, err
		}
		ids = append(ids, id)
	}
	return
}

func (m *FinanceDetailService) GenFinanceDetailID(ctx context.Context) (id int64, err error) {
	return time.Now().UnixMilli(), nil
}

func (m *FinanceDetailService) Add(ctx context.Context, req *top.FinanceDetailPO) (id, financeDetailID int64, err error) {
	if req.FinanceDetailID == 0 {
		req.FinanceDetailID, err = m.GenFinanceDetailID(ctx)
		if err != nil {
			return 0, 0, err
		}
	}

	if req.UpdatedBy == "" {
		req.CreatedBy = req.UpdatedBy
	}

	id, err = m.FinanceDetailDao.Insert(ctx, req)
	if err != nil {
		return
	}

	return id, req.FinanceDetailID, nil
}

func (m *FinanceDetailService) List(ctx context.Context, req *top.FinanceDetailPO) (pos []*top.FinanceDetailPO, err error) {
	financeDetailPOs, err := m.FinanceDetailDao.Select(ctx, req)
	if err != nil {
		return
	}

	return financeDetailPOs, nil
}
