package service

import (
	"github.com/gin-gonic/gin"
	"inventory/dao"
	"inventory/entity"
	"inventory/svc"
)

type InventoryService struct {
	c            *gin.Context
	svcCtx       *svc.ServiceContext
	inventoryDao *dao.InventoryDao
}

func NewInventoryService(c *gin.Context, svcCtx *svc.ServiceContext) *InventoryService {
	return &InventoryService{c: c, svcCtx: svcCtx, inventoryDao: dao.NewInventoryDao(c, svcCtx)}
}

func (receiver InventoryService) InventoryUse(req *entity.InventoryUseReq) error {
	return receiver.inventoryDao.InventoryUse(req.ProductID, req.ProductCount)
}
