package service

import (
	"github.com/gin-gonic/gin"
	"web-bff/entity"
	"web-bff/svc"
	"web-bff/utility"
)

type OrderService struct {
	c      *gin.Context
	svcCtx *svc.ServiceContext
}

func NewOrderService(c *gin.Context, svcCtx *svc.ServiceContext) *OrderService {
	return &OrderService{c: c, svcCtx: svcCtx}
}

func (receiver OrderService) OrderCreate(req *entity.OrderCreateReq) (map[string]any, error) {
	// 假定 userID 是 1
	var userID int64 = 1
	var err error

	// 创建订单
	orderClient := utility.OrderClient{}
	orderID, err := orderClient.CreateOrder(userID, req.OrderItems)
	if err != nil {
		return nil, err
	}

	// 扣减库存
	inventoryClient := utility.InventoryClient{}
	for _, item := range req.OrderItems {
		err = inventoryClient.UseInventory(userID, orderID, item.ProductID, item.Quantity)
		if err != nil {
			return nil, err
		}
	}

	//使用优惠券
	couponClient := utility.CouponClient{}
	err = couponClient.UseCoupon(userID, orderID, req.Coupon.CouponID, req.Coupon.DiscountAmount)
	if err != nil {
		return nil, err
	}

	// 更新订单状态
	order, err := orderClient.EnsureOrder(userID, orderID, req.Coupon.DiscountAmount)
	if err != nil {
		return nil, err
	}

	// 创建支付单

	paymentClient := utility.PaymentClient{}
	paymentID, err := paymentClient.CreatePayment(userID, orderID, order.TotalAmount)
	if err != nil {
		return nil, err
	}

	return map[string]any{
		"payment_id": paymentID,
	}, nil
}
