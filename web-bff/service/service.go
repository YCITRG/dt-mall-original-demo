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

	// 扣减库存
	inventoryClient := utility.InventoryClient{}
	for _, item := range req.OrderItems {
		err = inventoryClient.UseInventory(userID, req.OrderID, item.ProductID, item.Quantity)
		if err != nil {
			return nil, err
		}
	}

	//使用优惠券
	couponClient := utility.CouponClient{}
	err = couponClient.UseCoupon(userID, req.OrderID, req.Coupon.CouponID, req.Coupon.DiscountAmount)
	if err != nil {
		return nil, err
	}

	// 验证金额
	// 从产品服务获取产品价格数据，加上优惠券信息，计算出最终价格，该价格与入参中价格比较，验证是否合法
	// 无产品服务，假定校验成功，此处省略

	// 创建订单
	orderClient := utility.OrderClient{}
	err = orderClient.CreateOrder(req.OrderID, userID, req.OrderItems, req.Coupon.CouponID, req.Coupon.DiscountAmount)
	if err != nil {
		return nil, err
	}

	// 创建支付单
	paymentClient := utility.PaymentClient{}
	paymentID, err := paymentClient.CreatePayment(userID, req.OrderID, req.TotalAmount)
	if err != nil {
		return nil, err
	}

	return map[string]any{
		"payment_id": paymentID,
	}, nil
}
