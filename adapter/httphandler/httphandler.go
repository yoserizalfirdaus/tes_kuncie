package httphandler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/yoserizalfirdaus/tes_kuncie/entity"
	"github.com/yoserizalfirdaus/tes_kuncie/usecase/checkout"
)

type HttpHandler struct {
	checkoutUsecase CheckoutUsecase
}

type CheckoutUsecase interface {
	AddToCart(input checkout.AddToCartInput) (*entity.Cart, error)
	Checkout(input checkout.CheckoutInput) error
}

func NewHttpHandler(checkoutUsecase CheckoutUsecase) HttpHandler {
	return HttpHandler{checkoutUsecase}
}

type (
	addToCartRequest struct {
		TransactionID string `json:"transaction_id"`
		ProductSku    string `json:"product_sku"`
		Qty           int64  `json:"qty"`
	}
	cartResponse struct {
		Status  string   `json:"status"`
		Message string   `json:"message"`
		Cart    cartData `json:"cart"`
	}

	cartData struct {
		TransactionID   string      `json:"transaction_id"`
		TransactionTime string      `json:"transaction_time"`
		Items           []cartItem  `json:"items"`
		Promos          []cartPromo `json:"promos"`
		Subtotal        float64     `json:"subtotal"`
		Discount        float64     `json:"discount"`
		TotalAmount     float64     `json:"total_amount"`
	}

	cartItem struct {
		Sku        string  `json:"sku"`
		Name       string  `json:"name"`
		Qty        int64   `json:"qty"`
		Price      float64 `json:"price"`
		TotalPrice float64 `json:"total_price"`
	}

	cartPromo struct {
		ID          int64   `json:"id"`
		Name        string  `json:"name"`
		TotalAmount float64 `json:"total_amount"`
	}

	checkoutRequest struct {
		TransactionID string `json:"transaction_id"`
		PaymentMethod string `json:"payment_method"`
	}
)

func (h HttpHandler) AddToCart(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.Write([]byte(`{"status": "failed", "message": "failed read body"}`))
		return
	}

	req := addToCartRequest{}
	err = json.Unmarshal(body, &req)
	if err != nil {
		w.Write([]byte(`{"status": "failed", "message": "failed unmarshall"}`))
		return
	}

	cart, err := h.checkoutUsecase.AddToCart(checkout.AddToCartInput{
		TransactionID: req.TransactionID,
		ProductSku:    req.ProductSku,
		Qty:           req.Qty,
	})
	if err != nil {
		w.Write([]byte(fmt.Sprintf(`{"status": "failed", "message": "failed add to cart: %v"}`, err)))
		return
	}

	resp := cartResponse{
		Status:  "success",
		Message: "success",
		Cart: cartData{
			TransactionID:   cart.TransactionID,
			TransactionTime: cart.TransactionTime.Format(time.RFC3339),
			Subtotal:        cart.Subtotal,
			Discount:        cart.Discount,
			TotalAmount:     cart.TotalAmount,
		},
	}
	for _, v := range cart.Items {
		resp.Cart.Items = append(resp.Cart.Items, cartItem{
			Sku:        v.ProductSku,
			Name:       v.Name,
			Qty:        v.Qty,
			Price:      v.Price,
			TotalPrice: v.TotalPrice,
		})
	}
	for _, v := range cart.AppliedPromo {
		resp.Cart.Promos = append(resp.Cart.Promos, cartPromo{
			ID:          v.PromoID,
			Name:        v.Name,
			TotalAmount: v.TotalAmount,
		})
	}

	respByte, err := json.Marshal(resp)
	if err != nil {
		w.Write([]byte(`{"status": "failed", "message": "failed marshall"}`))
		return
	}
	w.Write(respByte)
}

func (h HttpHandler) Checkout(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.Write([]byte(`{"status": "failed", "message": "failed read body"}`))
		return
	}

	req := checkoutRequest{}
	err = json.Unmarshal(body, &req)
	if err != nil {
		w.Write([]byte(`{"status": "failed", "message": "failed unmarshall"}`))
		return
	}

	err = h.checkoutUsecase.Checkout(checkout.CheckoutInput{
		TransactionID: req.TransactionID,
		PaymentMethod: req.PaymentMethod,
	})
	if err != nil {
		w.Write([]byte(fmt.Sprintf(`{"status": "failed", "message": "failed add to cart: %v"}`, err)))
		return
	}
	w.Write([]byte(`{"status": "success", "message": "success"}`))
}
