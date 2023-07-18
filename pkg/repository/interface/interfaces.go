package interfaces

import (
	"github.com/google/uuid"
	"github.com/order/service/pkg/domain"
)

type OrderRepo interface {
	AddToCart(cartData []domain.CartProduct) error
	CashOnDelivery(paymentData domain.PaymentData) (int64, error)
	CreateOrder(orderData domain.Orders) (uuid.UUID, error)
	UpdateStatus(orderData domain.Orders, orderStatus string) int64
	UpdatePaymentStatus(paymentData domain.PaymentData, paymentmode string) int64
	GetTheOrderDetailsbyId(orderData domain.Orders) (domain.Orders, error)
	GetThePayementDetailsbyId(paymentData domain.PaymentData) (domain.PaymentData, error)
	GetAllOrders(orderData []domain.Orders) ([]domain.Orders, error)
	GetTheOrderdItems(orderData domain.Orders) ([]domain.CartProduct, error)
}
