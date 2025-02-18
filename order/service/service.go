package service

import (
	"github.com/gin-gonic/gin"
	"math/rand"
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

func (receiver OrderService) OrderCreate(req *entity.OrderCreateReq) (map[string]any, error) {
	orderID := rand.Int63()
	err := receiver.orderDao.OrderCreate(orderID, req.UserID, req.OrderItems)
	if err != nil {
		return nil, err
	}
	return map[string]any{
		"order_id": orderID,
	}, err
}

func (receiver OrderService) OrderEnsure(req *entity.OrderEnsureReq) (map[string]any, error) {
	order, err := receiver.orderDao.OrderEnsure(req.OrderID, req.UserID, req.DiscountAmount)
	if err != nil {
		return nil, err
	}

	return map[string]any{
		"id":           order.ID,
		"order_id":     order.OrderID,
		"user_id":      order.UserID,
		"order_status": order.OrderStatus,
		"total_amount": order.TotalAmount,
		"created_time": order.CreatedTime,
		"updated_time": order.UpdatedTime,
	}, err
}
