package mysql

import (
	"context"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/shuyi-tangerine/money/top"
	"strings"
	"time"
)

const (
	TableNameFinanceDetail = "finance_detail"
)

type FinanceDetailDao struct {
	DB *sqlx.DB
}

func NewFinanceDetailDao() (dao *FinanceDetailDao, err error) {
	db, err := sqlx.Connect("mysql", "root:123456@tcp(mysql.shuyi.com:3306)/shuyi?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		return
	}
	dao = &FinanceDetailDao{
		DB: db,
	}
	// 校验继承关系
	var _ top.FinanceDetailDao = dao
	return dao, nil
}

func (m *FinanceDetailDao) Select(ctx context.Context, req *top.FinanceDetailPO) (pos []*top.FinanceDetailPO, err error) {
	var params []interface{}
	where := strings.Builder{}
	where.WriteString("true")
	if req.ID != 0 {
		where.WriteString(" and id=?")
		params = append(params, req.ID)
	}
	if req.FinanceDetailID != 0 {
		where.WriteString(" and finance_detail_id=?")
		params = append(params, req.FinanceDetailID)
	}
	err = m.DB.SelectContext(ctx, &pos, fmt.Sprintf("select * from %s where %s", TableNameFinanceDetail, where.String()), params...)
	if err != nil {
		return
	}
	return
}

func (m *FinanceDetailDao) SelectOne(ctx context.Context, req *top.FinanceDetailPO) (po *top.FinanceDetailPO, err error) {
	pos, err := m.Select(ctx, req)
	if err != nil {
		return
	}

	if len(pos) > 1 {
		return nil, fmt.Errorf("selelct ont not one[%d]", len(pos))
	}

	if len(pos) == 0 {
		return
	}

	return pos[0], nil
}

func (m *FinanceDetailDao) Insert(ctx context.Context, req *top.FinanceDetailPO) (id int64, err error) {
	params := []interface{}{
		req.FinanceDetailID, req.AppID, req.Amount, req.OperatedType,
		req.OperatedAt, req.OperatedBy, req.Extra, req.CreatedBy, req.UpdatedBy,
	}
	fields := []string{
		`finance_detail_id`, `app_id`, `amount`, `operated_type`,
		`operated_at`, `operated_by`, `extra`, `created_by`, `updated_by`,
	}

	var placeholders []string
	for i := 0; i < len(fields); i++ {
		placeholders = append(placeholders, "?")
	}

	sql := fmt.Sprintf("insert into %s(%s) values(%s)", TableNameFinanceDetail, strings.Join(fields, ","), strings.Join(placeholders, ","))
	res, err := m.DB.ExecContext(ctx, sql, params...)
	if err != nil {
		return
	}
	return res.LastInsertId()
}

func (m *FinanceDetailDao) Update(ctx context.Context, req *top.FinanceDetailPO) (affectedRows int64, err error) {
	if req.ID == 0 && req.FinanceDetailID == 0 {
		return 0, fmt.Errorf("update no unique key(id or finance_detail_id)")
	}

	var params []interface{}
	var setFields []string
	if req.AppID != 0 {
		params = append(params, req.AppID)
		setFields = append(setFields, "`app_id`=?")
	}

	if req.Amount != 0 {
		params = append(params, req.Amount)
		setFields = append(setFields, "`amount`=?")
	}

	if req.OperatedType != 0 {
		params = append(params, req.OperatedType)
		setFields = append(setFields, "`operated_type`=?")
	}

	if req.OperatedAt.After(time.Time{}) {
		params = append(params, req.OperatedAt)
		setFields = append(setFields, "`operated_at`=?")
	}

	if req.OperatedBy != "" {
		params = append(params, req.OperatedBy)
		setFields = append(setFields, "`operated_by`=?")
	}

	if req.Extra != nil {
		params = append(params, req.Extra)
		setFields = append(setFields, "`extra`=?")
	}

	if req.UpdatedBy != "" {
		params = append(params, req.UpdatedBy)
		setFields = append(setFields, "`updated_by`=?")
	}

	var where []string
	if req.ID != 0 {
		where = append(where, "`id`=?")
		params = append(params, req.ID)
	}
	if req.FinanceDetailID != 0 {
		where = append(where, "`finance_detail_id`=?")
		params = append(params, req.FinanceDetailID)
	}

	sql := fmt.Sprintf("update %s set %s where %s", TableNameFinanceDetail, strings.Join(setFields, ","), strings.Join(where, "and"))
	res, err := m.DB.ExecContext(ctx, sql, params...)
	if err != nil {
		return
	}

	return res.RowsAffected()
}
