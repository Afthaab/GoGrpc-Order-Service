package interfaces

import (
	"github.com/google/uuid"
	"github.com/order/service/pkg/domain"
)

type OrderUseCase interface {
	AddToCart(cartData []domain.CartProduct) error
	CashOnDelivery(paymentData domain.PaymentData) (int64, error)
	CreateOrder(orderData domain.Orders) (uuid.UUID, error)
	OrderDispacth(orderData domain.Orders) error
	CancelOrder(orderData domain.Orders) error
	GetAllOrders(orderData []domain.Orders) ([]domain.Orders, error)
	GetTheOrderDetailsById(orderData domain.Orders) (domain.Orders, error)
	GetThePaymentDetailsById(orderData domain.Orders) (domain.PaymentData, error)
	GetTheOrderItems(orderData domain.Orders) ([]domain.CartProduct, error)
}
