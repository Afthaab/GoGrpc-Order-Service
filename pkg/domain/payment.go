package domain

import (
	"time"

	"github.com/google/uuid"
)

type PaymentData struct {
	ID            uint    `json:"id"`
	Totalamount   float32 `JSON:"totalamount"`
	Paymentmode   string  `JSON:"paymentmode"`
	Paymentstatus string  `JSON:"paymentstatus"`
	Useridno      int64   `JSON:"useridno"`
	Couponid      string  `JSON:"couponid"`
}

type Orders struct {
	Orderid     uuid.UUID `json:"orderid" gorm:"type:uuid;default:gen_random_uuid();not null;primaryKey"`
	Useridno    int64     `json:"useridno"`
	Paymentid   int64     `json:"paymentid"`
	Addid       int64     `json:"addid"`
	Orderstatus string    `json:"orderstatus"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Orderdtime  string
	Updatedtime string
	Orderddate  string
}
