package top

import (
	"context"
	"github.com/shuyi-tangerine/money/gen-go/tangerine/money"
	"time"
)

type FinanceDetailPO struct {
	ID              int64     `json:"id" db:"id"`
	FinanceDetailID int64     `json:"finance_detail_id" db:"finance_detail_id"`
	AppID           int64     `json:"app_id" db:"app_id"`
	Amount          int64     `json:"amount" db:"amount"`
	OperatedType    int64     `json:"operated_type" db:"operated_type"`
	OperatedAt      time.Time `json:"operated_at" db:"operated_at"`
	OperatedBy      string    `json:"operated_by" db:"operated_by"`
	Extra           *string   `json:"extra" db:"extra"`
	CreatedAt       time.Time `json:"created_at" db:"created_at"`
	CreatedBy       string    `json:"created_by" db:"created_by"`
	UpdatedAt       time.Time `json:"updated_at" db:"updated_at"`
	UpdatedBy       string    `json:"updated_by" db:"updated_by"`

	// 查询的时候当参数用
	OperatedAtTimeRange *money.TimeRange `json:"-"`
	Limit               int64            `json:"-"`
	Offset              int64            `json:"-"`
}

func (m *FinanceDetailPO) ToRpcFinanceDetailInfo(ctx context.Context) (rpcFinanceDetailInfo *money.FinanceDetailInfo) {
	rpcFinanceDetailInfo = &money.FinanceDetailInfo{
		ID:              m.ID,
		FinanceDetailID: m.FinanceDetailID,
		AppID:           m.AppID,
		Amount:          m.Amount,
		OperatedType:    m.OperatedType,
		OperatedAt:      m.OperatedAt.Unix(),
		OperatedBy:      m.OperatedBy,
		Extra:           m.Extra,
		CreatedAt:       m.CreatedAt.Unix(),
		CreatedBy:       m.CreatedBy,
		UpdatedAt:       m.UpdatedAt.Unix(),
		UpdatedBy:       m.UpdatedBy,
	}
	return
}

func NewRpcFinanceDetailInfos(ctx context.Context, pos []*FinanceDetailPO) (rpcFinanceDetailInfos []*money.FinanceDetailInfo) {
	for _, v := range pos {
		rpcFinanceDetailInfos = append(rpcFinanceDetailInfos, v.ToRpcFinanceDetailInfo(ctx))
	}
	return
}

type FinanceDetailDao interface {
	Select(ctx context.Context, req *FinanceDetailPO) (pos []*FinanceDetailPO, err error)
	SelectOne(ctx context.Context, req *FinanceDetailPO) (po *FinanceDetailPO, err error)
	Insert(ctx context.Context, req *FinanceDetailPO) (id int64, err error)
	Update(ctx context.Context, req *FinanceDetailPO) (affectedRows int64, err error)
}
