package service

import (
	"github.com/gin-gonic/gin"
	"order/dao"
	"order/entity"
	"order/svc"
)

type OrderService struct {
	c        *gin.Context
	svcCtx   *svc.ServiceContext
	orderDao *dao.OrderDao
}

func NewOrderService(c *gin.Context, svcCtx *svc.ServiceContext) *OrderService {
	return &OrderService{c: c, svcCtx: svcCtx, orderDao: dao.NewOrderDao(c, svcCtx)}
}

func (receiver OrderService) OrderCreate(req *entity.OrderCreateReq) error {
	err := receiver.orderDao.OrderCreate(req.OrderID, req.UserID, req.OrderItems, req.CouponID, req.DiscountAmount)
	return err
}
