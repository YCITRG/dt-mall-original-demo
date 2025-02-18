package dao

import (
	"github.com/gin-gonic/gin"
	"inventory/svc"
)

type InventoryDao struct {
	c      *gin.Context
	svcCtx *svc.ServiceContext
}

func NewInventoryDao(c *gin.Context, svcCtx *svc.ServiceContext) *InventoryDao {
	return &InventoryDao{c: c, svcCtx: svcCtx}
}

func (receiver InventoryDao) InventoryUse(productID, productCount int) error {
	r, err := receiver.svcCtx.DB.Exec("update tb_inventory set stock_quantity=stock_quantity-?, updated_time=now() where product_id=? and stock_quantity >= ?", productCount, productID, productCount)
	if err != nil {
		return err
	}
	affected, err := r.RowsAffected()
	if err != nil || affected != 1 {
		return err
	}
	return nil
}
