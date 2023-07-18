package repository

import (
	"github.com/google/uuid"
	"github.com/order/service/pkg/domain"
	interfaces "github.com/order/service/pkg/repository/interface"
	"gorm.io/gorm"
)

type OrderDatabase struct {
	DB *gorm.DB
}

func (r *OrderDatabase) AddToCart(cartData []domain.CartProduct) error {
	result := r.DB.Create(&cartData).Error
	return result
}

func (r *OrderDatabase) CashOnDelivery(paymentData domain.PaymentData) (int64, error) {
	result := r.DB.Create(&paymentData)
	return int64(paymentData.ID), result.Error

}

func (r *OrderDatabase) CreateOrder(orderData domain.Orders) (uuid.UUID, error) {
	resulr := r.DB.Create(&orderData)
	return orderData.Orderid, resulr.Error
}
func (r *OrderDatabase) UpdateStatus(orderData domain.Orders, orderStatus string) int64 {
	result := r.DB.Exec("update orders set orderstatus = ? where orderid = ?", orderStatus, orderData.Orderid)
	return result.RowsAffected
}

func (r *OrderDatabase) GetTheOrderDetailsbyId(orderData domain.Orders) (domain.Orders, error) {
	result := r.DB.Raw("select *,(EXTRACT(HOUR FROM created_at) || ':' || EXTRACT(MINUTE FROM created_at)) AS orderdtime, (EXTRACT(HOUR FROM created_at) || ':' || EXTRACT(MINUTE FROM updated_at)) AS updatedtime,DATE(created_at) AS orderddate FROM orders where orderid = ?", orderData.Orderid).Scan(&orderData)
	return orderData, result.Error
}

func (r *OrderDatabase) GetThePayementDetailsbyId(paymentData domain.PaymentData) (domain.PaymentData, error) {
	result := r.DB.Raw("select * from payment_data where id = ?", paymentData.ID).Scan(&paymentData)
	return paymentData, result.Error
}

func (r *OrderDatabase) UpdatePaymentStatus(paymentData domain.PaymentData, paymentStatus string) int64 {
	result := r.DB.Exec("update payment_data set paymentstatus = ? where id = ?", paymentStatus, paymentData.ID)
	return result.RowsAffected
}

func (r *OrderDatabase) GetAllOrders(orderData []domain.Orders) ([]domain.Orders, error) {
	result := r.DB.Raw("SELECT *,(EXTRACT(HOUR FROM created_at) || ':' || EXTRACT(MINUTE FROM created_at)) AS orderdtime, (EXTRACT(HOUR FROM created_at) || ':' || EXTRACT(MINUTE FROM updated_at)) AS updatedtime,DATE(created_at) AS orderddate FROM orders;").Scan(&orderData)
	return orderData, result.Error
}

func (r *OrderDatabase) GetTheOrderdItems(orderData domain.Orders) ([]domain.CartProduct, error) {
	orderDatas := []domain.CartProduct{}
	result := r.DB.Raw("select * from cart_products where orderid = ?", orderData.Orderid).Scan(&orderDatas)
	return orderDatas, result.Error
}
func NewOrderRepo(db *gorm.DB) interfaces.OrderRepo {
	return &OrderDatabase{
		DB: db,
	}
}
