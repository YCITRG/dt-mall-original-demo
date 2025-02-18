package service

import (
	"coupon/dao"
	"coupon/entity"
	"coupon/svc"
	"github.com/gin-gonic/gin"
)

type CouponService struct {
	c            *gin.Context
	svcCtx       *svc.ServiceContext
	inventoryDao *dao.CouponDao
}

func NewCouponService(c *gin.Context, svcCtx *svc.ServiceContext) *CouponService {
	return &CouponService{c: c, svcCtx: svcCtx, inventoryDao: dao.NewCouponDao(c, svcCtx)}
}

func (receiver CouponService) CouponUse(req *entity.CouponUseReq) error {
	return receiver.inventoryDao.CouponUse(req.UserID, req.CouponID, req.DiscountAmount)
}
