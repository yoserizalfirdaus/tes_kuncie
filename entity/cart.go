package entity

import (
	"time"
)

type Cart struct {
	ID              int64
	TransactionID   string
	TransactionTime time.Time
	Items           map[string]CartItem
	AppliedPromo    map[string]CartAppliedPromo
	ApplicablePromo map[string]CartApplicablePromo
	Subtotal        float64
	Discount        float64
	TotalAmount     float64
}

type CartItem struct {
	ProductSku        string
	Name              string
	Price             float64
	Qty               int64
	TotalPrice        float64
	PromoReqCount     int64
	PromoAppliedCount int64
}

type CartAppliedPromo struct {
	PromoID     int64
	ProductSku  string
	Name        string
	Qty         int64
	TotalAmount float64
}

type CartApplicablePromo struct {
	PromoID    int64
	ProductSku string
	Name       string
	Amount     float64
	Qty        int64
}

func NewCart(trxID string) *Cart {
	c := Cart{
		TransactionID:   trxID,
		TransactionTime: time.Now(),
		Items:           map[string]CartItem{},
		AppliedPromo:    map[string]CartAppliedPromo{},
		ApplicablePromo: map[string]CartApplicablePromo{},
	}
	return &c
}

func (c *Cart) AddToCart(item CartItem) {
	ci, ok := c.Items[item.ProductSku]
	if !ok {
		ci = item
	} else {
		ci.Qty += item.Qty
	}
	ci.TotalPrice += (float64(item.Qty) * item.Price)

	c.Items[item.ProductSku] = ci
}

func (c *Cart) AddToApplicablePromo(applicablePromo CartApplicablePromo) {
	cap, ok := c.ApplicablePromo[applicablePromo.ProductSku]
	if !ok {
		cap = applicablePromo
	} else {
		cap.Qty += applicablePromo.Qty
	}
	c.ApplicablePromo[applicablePromo.ProductSku] = cap
}

func (c *Cart) SubstractApplicablePromo(productSku string, qty int64) {
	cap, ok := c.ApplicablePromo[productSku]
	if ok {
		cap.Qty -= qty
	}

	c.ApplicablePromo[productSku] = cap
}

func (c *Cart) AddToAppliedPromo(appliedPromo CartAppliedPromo) {
	cadp, ok := c.AppliedPromo[appliedPromo.ProductSku]
	if !ok {
		cadp = appliedPromo
	} else {
		cadp.Qty += appliedPromo.Qty
		cadp.TotalAmount += appliedPromo.TotalAmount
	}

	c.AppliedPromo[appliedPromo.ProductSku] = cadp
}

func (c *Cart) CalculateAmount() {
	var subtotal, discount float64

	for _, ci := range c.Items {
		subtotal += ci.TotalPrice
	}

	for _, cp := range c.AppliedPromo {
		discount += cp.TotalAmount
	}

	c.Subtotal = subtotal
	c.Discount = discount
	c.TotalAmount = subtotal - discount
}
