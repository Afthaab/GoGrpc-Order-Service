package domain

import "github.com/google/uuid"

type CartProduct struct {
	Id          uint      ` json:"id" `
	Pid         string    ` json:"Pid"`
	Uid         int64     ` json:"uid"`
	Productname string    ` json:"productname"`
	Sizename    string    ` json:"sizename"`
	Price       float32   ` json:"Price"`
	Imageurls   string    ` json:"imageurl"`
	Quantity    int64     `json:"quantity"`
	Orderid     uuid.UUID `json:"orderid"`
}
