package dao

import (
	"github.com/gin-gonic/gin"
	"payment/svc"
)

type PaymentDao struct {
	c      *gin.Context
	svcCtx *svc.ServiceContext
}

func NewPaymentDao(c *gin.Context, svcCtx *svc.ServiceContext) *PaymentDao {
	return &PaymentDao{c: c, svcCtx: svcCtx}
}

func (receiver PaymentDao) PaymentCreate(paymentId, userID, orderId int64, paymentAmount float64) error {
	r, err := receiver.svcCtx.DB.Exec("insert into tb_payment_order(payment_id, user_id, order_id, payment_amount) value (?, ?, ?, ?)", paymentId, userID, orderId, paymentAmount)
	if err != nil {
		return err
	}
	affected, err := r.RowsAffected()
	if err != nil || affected != 1 {
		return err
	}
	return nil
}
