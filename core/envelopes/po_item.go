package envelopes

import (
	"database/sql"
	"github.com/shopspring/decimal"
	"time"
)

type RedEnvelopeItem struct {
	Id           int64           `db:"id,omitempty"`         //自增ID
	ItemNo       string          `db:"item_no,uni"`          //红包订单详情编号
	EnvelopeNo   string          `db:"envelope_no"`          //红包编号,红包唯一标识
	RecvUsername sql.NullString  `db:"recv_username"`        //红包接收者用户名称
	RecvUserId   string          `db:"recv_user_id"`         //红包接收者用户编号
	Amount       decimal.Decimal `db:"amount"`               //收到金额
	Quantity     int             `db:"quantity"`             //收到数量：对于收红包来说是1
	RemainAmount decimal.Decimal `db:"remain_amount"`        //收到后红包剩余金额
	AccountNo    string          `db:"account_no"`           //红包接收者账户ID
	PayStatus    int             `db:"pay_status"`           //支付状态：未支付，支付中，已支付，支付失败
	CreatedAt    time.Time       `db:"created_at,omitempty"` //创建时间
	UpdatedAt    time.Time       `db:"updated_at,omitempty"` //更新时间
	Desc         string          `db:"desc"`
}