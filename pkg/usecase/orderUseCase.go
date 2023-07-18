package usecase

import (
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/order/service/pkg/domain"
	interfaces "github.com/order/service/pkg/repository/interface"
	usecase "github.com/order/service/pkg/usecase/interface"
)

type OrderRepo struct {
	Repo interfaces.OrderRepo
}

func (u *OrderRepo) AddToCart(cartData []domain.CartProduct) error {
	err := u.Repo.AddToCart(cartData)
	if err != nil {
		return err
	}
	return nil
}

func (u *OrderRepo) CashOnDelivery(paymentData domain.PaymentData) (int64, error) {
	id, err := u.Repo.CashOnDelivery(paymentData)
	if err != nil {
		return id, err
	}
	return id, nil
}

func (u *OrderRepo) CreateOrder(orderData domain.Orders) (uuid.UUID, error) {
	id, err := u.Repo.CreateOrder(orderData)
	if err != nil {
		return id, err
	}
	return id, nil
}

func (u *OrderRepo) OrderDispacth(orderData domain.Orders) error {
	result := u.Repo.UpdateStatus(orderData, "Order Dispatched")
	if result == 0 {
		return errors.New("Could dot not update the order status")
	}
	return nil
}
func (u *OrderRepo) CancelOrder(orderData domain.Orders) error {
	// get the order details first
	orderData, err := u.Repo.GetTheOrderDetailsbyId(orderData)
	if err != nil {
		return err
	}

	if orderData.Orderstatus == "Order Placed" {
		// get the payment details from the payment table
		paymentData := domain.PaymentData{
			ID: uint(orderData.Paymentid),
		}

		paymentData, err = u.Repo.GetThePayementDetailsbyId(paymentData)
		if err != nil {
			return err
		}
		if paymentData.Paymentmode == "Cash on Delivery" {
			fmt.Println("Return to the wallet")

			result := u.Repo.UpdateStatus(orderData, "Order Cancelled")
			if result == 0 {
				return errors.New("Could dot not update the order status")
			}

			result = u.Repo.UpdatePaymentStatus(paymentData, "Refunded")
			if result == 0 {
				return errors.New("Could not update the payment status")
			}
		}

		return nil

	}
	return errors.New("Cannot cancel the order")
}

func (u *OrderRepo) GetTheOrderDetailsById(orderData domain.Orders) (domain.Orders, error) {
	orderData, err := u.Repo.GetTheOrderDetailsbyId(orderData)
	if err != nil {
		return orderData, err
	}
	return orderData, nil
}

func (u *OrderRepo) GetAllOrders(orderData []domain.Orders) ([]domain.Orders, error) {
	orderData, err := u.Repo.GetAllOrders(orderData)
	if err != nil {
		return orderData, err
	}

	return orderData, nil
}
func (u *OrderRepo) GetThePaymentDetailsById(orderData domain.Orders) (domain.PaymentData, error) {
	paymentData := domain.PaymentData{
		ID: uint(orderData.Paymentid),
	}
	paymentData, err := u.Repo.GetThePayementDetailsbyId(paymentData)
	if err != nil {
		return paymentData, err
	}
	return paymentData, nil
}

func (u *OrderRepo) GetTheOrderItems(orderData domain.Orders) ([]domain.CartProduct, error) {
	orderDatas, err := u.Repo.GetTheOrderdItems(orderData)
	if err != nil {
		return orderDatas, err
	}
	return orderDatas, nil
}

func NewOrderUseCase(repo interfaces.OrderRepo) usecase.OrderUseCase {
	return &OrderRepo{
		Repo: repo,
	}
}
