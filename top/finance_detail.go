package top

import (
	"context"
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
}

type FinanceDetailDao interface {
	Select(ctx context.Context, req *FinanceDetailPO) (pos []*FinanceDetailPO, err error)
	SelectOne(ctx context.Context, req *FinanceDetailPO) (po *FinanceDetailPO, err error)
	Insert(ctx context.Context, req *FinanceDetailPO) (id int64, err error)
	Update(ctx context.Context, req *FinanceDetailPO) (affectedRows int64, err error)
}
