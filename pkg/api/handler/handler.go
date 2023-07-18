package handler

import (
	"context"
	"net/http"

	"github.com/google/uuid"
	"github.com/order/service/pkg/domain"
	"github.com/order/service/pkg/pb"
	interfaces "github.com/order/service/pkg/usecase/interface"
)

type OrderHandler struct {
	usecase interfaces.OrderUseCase
	pb.OrderManagementServer
}

func (h *OrderHandler) SaveCartItems(ctx context.Context, req *pb.CartItemsSaveRequest) (*pb.CartItemsSaveResponse, error) {
	var cartData []domain.CartProduct
	uuidstr, _ := uuid.Parse(req.Orderid)
	for _, item := range req.Items {
		data := domain.CartProduct{
			Uid:         item.Uid,
			Pid:         item.Pid,
			Productname: item.Productname,
			Sizename:    item.Sizename,
			Price:       item.Price,
			Imageurls:   item.Imagelink,
			Quantity:    item.Quantity,
			Orderid:     uuidstr,
		}
		cartData = append(cartData, data)
	}

	err := h.usecase.AddToCart(cartData)
	if err != nil {
		return &pb.CartItemsSaveResponse{
			Status: http.StatusBadRequest,
			Error:  "Could not add the product to the cart",
		}, err
	} else {
		return &pb.CartItemsSaveResponse{
			Status: http.StatusOK,
			Error:  "nil",
		}, nil
	}
}

func (h *OrderHandler) CashOnDelivery(ctx context.Context, req *pb.CashOnDeliveryRequest) (*pb.CashOnDeliveryResponse, error) {
	paymentData := domain.PaymentData{
		Totalamount:   req.Totalamount,
		Paymentmode:   req.Paymentmode,
		Paymentstatus: req.Paymentstatus,
		Useridno:      req.Userid,
		Couponid:      req.Coupondid,
	}
	id, err := h.usecase.CashOnDelivery(paymentData)
	if err != nil {
		return &pb.CashOnDeliveryResponse{
			Status: http.StatusUnprocessableEntity,
			Error:  "Could not initiate cash on delivery",
		}, err
	} else {
		return &pb.CashOnDeliveryResponse{
			Status:    http.StatusOK,
			Error:     "nil",
			Paymentid: id,
		}, nil
	}
}

func (h *OrderHandler) CreateOrder(ctx context.Context, req *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {
	orderData := domain.Orders{
		Useridno:    req.Userid,
		Paymentid:   req.Paymentid,
		Addid:       req.Addressid,
		Orderstatus: req.Orderstatus,
	}
	id, err := h.usecase.CreateOrder(orderData)
	if err != nil {
		return &pb.CreateOrderResponse{
			Status: http.StatusUnprocessableEntity,
			Error:  "Could not create order",
		}, err
	} else {
		return &pb.CreateOrderResponse{
			Status:  http.StatusOK,
			Error:   "nil",
			Orderid: id.String(),
		}, nil
	}
}

func (h *OrderHandler) OrderDispacth(ctx context.Context, req *pb.OrderDispacthRequest) (*pb.OrderDispacthResponse, error) {
	uuidstr, _ := uuid.Parse(req.Orderid)
	orderData := domain.Orders{
		Orderid: uuidstr,
	}
	err := h.usecase.OrderDispacth(orderData)
	if err != nil {
		return &pb.OrderDispacthResponse{
			Status: http.StatusBadRequest,
			Error:  "Could not update the order",
		}, err
	} else {
		return &pb.OrderDispacthResponse{
			Status: http.StatusOK,
			Error:  "nil",
		}, nil
	}
}

func (h *OrderHandler) OrderCancel(ctx context.Context, req *pb.OrderCancelRequest) (*pb.OrderCancelResponse, error) {
	uuidstr, _ := uuid.Parse(req.Orderid)
	orderData := domain.Orders{
		Orderid: uuidstr,
	}
	err := h.usecase.CancelOrder(orderData)
	if err != nil {
		return &pb.OrderCancelResponse{
			Status: http.StatusBadRequest,
			Error:  "Could not cancel the order",
		}, err
	} else {
		return &pb.OrderCancelResponse{
			Status: http.StatusOK,
			Error:  "nil",
		}, nil
	}
}

func (h *OrderHandler) AdminViewAllOrders(ctx context.Context, req *pb.AdminViewAllOrdersRequest) (*pb.AdminViewAllOrdersResponse, error) {
	orderData := []domain.Orders{}
	orderData, err := h.usecase.GetAllOrders(orderData)
	if err != nil {
		return &pb.AdminViewAllOrdersResponse{
			Status: http.StatusNotFound,
			Error:  "Could not get all orders details",
		}, err
	} else {
		var orderDetails []*pb.OrderViewDetails
		for _, items := range orderData {
			orderDatail := pb.OrderViewDetails{
				Orderid:     items.Orderid.String(),
				Orderstatus: items.Orderstatus,
				Orderddatae: items.Orderddate,
				Orderdtime:  items.Orderdtime,
				Updatedtime: items.Updatedtime,
			}
			orderDetails = append(orderDetails, &orderDatail)
		}
		return &pb.AdminViewAllOrdersResponse{
			Status:       http.StatusOK,
			Error:        "nil",
			Orderdetails: orderDetails,
		}, nil
	}

}

func (h *OrderHandler) ViewOrderById(ctx context.Context, req *pb.ViewOrderByIdRequest) (*pb.ViewOrderByIdResponse, error) {
	uuidstr, _ := uuid.Parse(req.Id)
	orderData := domain.Orders{
		Orderid: uuidstr,
	}
	orderData, err := h.usecase.GetTheOrderDetailsById(orderData)
	if err != nil {
		return &pb.ViewOrderByIdResponse{
			Status: http.StatusNotFound,
			Error:  "Could not get the order details",
		}, err
	}
	paymentData, err := h.usecase.GetThePaymentDetailsById(orderData)
	if err != nil {
		return &pb.ViewOrderByIdResponse{
			Status: http.StatusNotFound,
			Error:  "Could not get the payment data",
		}, err
	}
	itemsData, err := h.usecase.GetTheOrderItems(orderData)
	if err != nil {
		return &pb.ViewOrderByIdResponse{
			Status: http.StatusNotFound,
			Error:  "Could not get the orderd items",
		}, err
	}

	details := &pb.OrderViewDetails{
		Orderid:     orderData.Orderid.String(),
		Orderstatus: orderData.Orderstatus,
		Orderddatae: orderData.Orderddate,
		Orderdtime:  orderData.Orderdtime,
		Updatedtime: orderData.Updatedtime,
	}

	payment := &pb.CashOnDeliveryRequest{
		Totalamount:   paymentData.Totalamount,
		Paymentmode:   paymentData.Paymentmode,
		Paymentstatus: paymentData.Paymentstatus,
		Coupondid:     paymentData.Couponid,
	}

	var orderdItems []*pb.CartItems
	for _, items := range itemsData {
		item := &pb.CartItems{
			Pid:         items.Pid,
			Productname: items.Productname,
			Sizename:    items.Sizename,
			Price:       items.Price,
			Imagelink:   items.Imageurls,
			Quantity:    items.Quantity,
		}
		orderdItems = append(orderdItems, item)

	}

	return &pb.ViewOrderByIdResponse{
		Status:         http.StatusOK,
		Error:          "nil",
		Orderdetails:   details,
		Paymentdetails: payment,
		Items:          orderdItems,
	}, nil

}

func NewOrderHandler(Usecase interfaces.OrderUseCase) *OrderHandler {
	return &OrderHandler{
		usecase: Usecase,
	}
}
