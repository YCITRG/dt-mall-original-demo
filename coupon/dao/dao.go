package dao

import (
	"coupon/svc"
	"errors"
	"github.com/gin-gonic/gin"
)

type CouponDao struct {
	c      *gin.Context
	svcCtx *svc.ServiceContext
}

func NewCouponDao(c *gin.Context, svcCtx *svc.ServiceContext) *CouponDao {
	return &CouponDao{c: c, svcCtx: svcCtx}
}

func (receiver CouponDao) CouponUse(userID, couponID int64, discountAmount float64) error {
	r, err := receiver.svcCtx.DB.Exec("update tb_coupon set is_used=true where user_id=? and coupon_id = ? and is_used = false and now() <= expiration_date and discount_amount = ?", userID, couponID, discountAmount)
	if err != nil {
		return err
	}
	affected, err := r.RowsAffected()
	if err != nil {
		return err
	}
	if affected != 1 {
		return errors.New("优惠券不存在、不合法、已过期或已被消费")
	}
	return nil
}
