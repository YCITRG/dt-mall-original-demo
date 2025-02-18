package service

import (
	"github.com/gin-gonic/gin"
	"math/rand"
	"payment/dao"
	"payment/entity"
	"payment/svc"
)

type PaymentService struct {
	c          *gin.Context
	svcCtx     *svc.ServiceContext
	paymentDao *dao.PaymentDao
}

func NewPaymentService(c *gin.Context, svcCtx *svc.ServiceContext) *PaymentService {
	return &PaymentService{c: c, svcCtx: svcCtx, paymentDao: dao.NewPaymentDao(c, svcCtx)}
}

func (receiver PaymentService) PaymentCreate(req *entity.PaymentCreateReq) (map[string]interface{}, error) {
	paymentId := rand.Int63()
	err := receiver.paymentDao.PaymentCreate(paymentId, req.UserID, req.OrderID, req.PaymentAmount)
	if err != nil {
		return nil, err
	}
	return map[string]interface{}{
		"payment_id": paymentId,
	}, err
}
