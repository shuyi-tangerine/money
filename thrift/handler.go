package thrift

import (
	"context"
	"github.com/shuyi-tangerine/money/biz"
	"github.com/shuyi-tangerine/money/gen-go/base"
	"github.com/shuyi-tangerine/money/gen-go/tangerine/money"
	"github.com/shuyi-tangerine/money/top"
	"time"
)

type MoneyHandler struct {
	FinanceDetailService *biz.FinanceDetailService
}

func NewMoneyHandler() (handler money.MoneyHandler, err error) {
	financeDetailService, err := biz.NewFinanceDetailService()
	if err != nil {
		return
	}
	return &MoneyHandler{
		FinanceDetailService: financeDetailService,
	}, nil
}

func (m *MoneyHandler) GenFinanceDetailIDs(ctx context.Context, req *money.GenFinanceDetailIDsRequest) (resp *money.GenFinanceDetailIDsResponse, err error) {
	//TODO implement me
	panic("implement me")
}

func (m *MoneyHandler) AddFinanceDetail(ctx context.Context, req *money.AddFinanceDetailRequest) (resp *money.AddFinanceDetailResponse, err error) {
	resp = money.NewAddFinanceDetailResponse()
	resp.Base = base.NewRPCResponse()
	resp.Base.Code = 0
	resp.Base.Message = "ok"
	id, financeDetailID, err := m.FinanceDetailService.Add(ctx, &top.FinanceDetailPO{
		ID:              0,
		FinanceDetailID: req.FinanceDetailID,
		AppID:           req.AppID,
		Amount:          req.Amount,
		OperatedType:    req.OperatedType,
		OperatedAt:      time.Unix(req.OperatedAt, 0),
		OperatedBy:      req.OperatedBy,
		Extra:           req.Extra,
		CreatedAt:       time.Time{},
		CreatedBy:       req.CreatedBy,
		UpdatedAt:       time.Time{},
		UpdatedBy:       req.CreatedBy,
	})
	if err != nil {
		resp.Base.Code = -1
		resp.Base.Message = err.Error()
		return
	}

	resp.Info = money.NewFinanceDetailInfo()
	resp.Info.ID = id
	resp.Info.FinanceDetailID = financeDetailID
	return
}

func (m *MoneyHandler) ListFinanceDetail(ctx context.Context, req *money.ListFinanceDetailRequest) (resp *money.ListFinanceDetailResponse, err error) {
	resp = money.NewListFinanceDetailResponse()
	resp.Base = base.NewRPCResponse()
	resp.Base.Code = 0
	resp.Base.Message = "ok"
	financeDetailPOs, err := m.FinanceDetailService.List(ctx, &top.FinanceDetailPO{
		AppID:               req.AppID,
		OperatedAtTimeRange: req.OperatedAt,
		Limit:               req.Limit,
		Offset:              req.Offset,
	})
	if err != nil {
		resp.Base.Code = -1
		resp.Base.Message = err.Error()
		return
	}

	resp.FinanceDetails = top.NewRpcFinanceDetailInfos(ctx, financeDetailPOs)
	resp.Offset = 0 // TODO
	resp.Total = 0  // TODO
	return
}
