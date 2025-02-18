package dao

import (
	"database/sql"
	"errors"
	"github.com/gin-gonic/gin"
	"order/entity"
	"order/svc"
)

type OrderDao struct {
	c      *gin.Context
	svcCtx *svc.ServiceContext
}

func NewOrderDao(c *gin.Context, svcCtx *svc.ServiceContext) *OrderDao {
	return &OrderDao{c: c, svcCtx: svcCtx}
}

func (receiver OrderDao) OrderCreate(orderID int64, userID int64, orderItems []entity.OrderItem, couponID int64, discountAmount int64) error {
	var totalAmount float64 = -float64(discountAmount)
	var r sql.Result
	var err error
	var affected int64

	for _, item := range orderItems {
		totalAmount += float64(item.Quantity) * item.UnitPrice
		r, err = receiver.svcCtx.DB.Exec("insert into tb_order_item (order_id, product_id, quantity, unit_price) value (?, ?, ?, ?)", orderID, item.ProductID, item.Quantity, item.UnitPrice)
		if err != nil {
			return err
		}
		affected, err = r.RowsAffected()
		if err != nil {
			return err
		}
		if err != nil || affected != 1 {
			return err
		}
	}

	r, err = receiver.svcCtx.DB.Exec("insert into tb_order (order_id, user_id, order_status, total_amount) value (?, ?, ?, ?)", orderID, userID, "ENSURE", totalAmount)
	if err != nil {
		return err
	}
	affected, err = r.RowsAffected()
	if err != nil || affected != 1 {
		return errors.New("创建订单失败")
	}

	return nil
}

//func (receiver OrderDao) OrderEnsure(orderID int64, userID int64, discountAmount float64) (*entity.Order, error) {
//	var r sql.Result
//	var err error
//	var affected int64
//
//	//r, err = receiver.svcCtx.DB.Exec("insert into tb_order (order_id, user_id, total_amount) value (?, ?, ?)", orderID, userID, totalAmount)
//
//	r, err = receiver.svcCtx.DB.Exec("update tb_order set order_status = ?, total_amount = total_amount - ? where order_id = ? and user_id = ? and order_status = ?", "ENSURE", discountAmount, orderID, userID, "CREATED")
//
//	if err != nil {
//		return nil, err
//	}
//	affected, err = r.RowsAffected()
//	if err != nil || affected != 1 {
//		return nil, errors.New("确认订单失败")
//	}
//
//	rows, err := receiver.svcCtx.DB.Query("select * from tb_order where order_id = ? and user_id = ?", orderID, userID)
//	if err != nil {
//		return nil, err
//	}
//
//	if !rows.Next() {
//		return nil, errors.New(fmt.Sprintf("订单 %c 找不到", orderID))
//	}
//
//	var (
//		id          int64
//		orderId     int64
//		userId      int64
//		orderStatus string
//		totalAmount float64
//		createdTime string
//		updatedTime string
//	)
//	err = rows.Scan(&id, &orderId, &userId, &orderStatus, &totalAmount, &createdTime, &updatedTime)
//	if err != nil {
//		log.Println(err)
//		return nil, err
//	}
//
//	parsedCreatedTime, err := time.Parse(time.DateTime, createdTime)
//	if err != nil {
//		return nil, err
//	}
//
//	parsedUpdatedTime, err := time.Parse(time.DateTime, updatedTime)
//	if err != nil {
//		return nil, err
//	}
//
//	return &entity.Order{
//		ID: id, OrderID: orderId, UserID: userId, OrderStatus: orderStatus, TotalAmount: totalAmount, CreatedTime: parsedCreatedTime, UpdatedTime: parsedUpdatedTime,
//	}, nil
//}
