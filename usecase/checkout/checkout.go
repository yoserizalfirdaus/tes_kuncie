package checkout

import (
	"fmt"
	"log"
	"math"

	"github.com/yoserizalfirdaus/tes_kuncie/entity"
)

type CheckoutUsecase struct {
	productRepo ProductRepository
	cartRepo    CartRepository
	promoRepo   PromoRepository
}

type ProductRepository interface {
	GetBySku(sku string) (entity.Product, error)
}

type CartRepository interface {
	GetByTransactionID(trxID string) (*entity.Cart, error)
	SaveCart(cart *entity.Cart) error
	ClearCart(trxID string) error
}

type PromoRepository interface {
	GetPromoByProductSku(productSku string) (entity.Promotion, error)
}

type AddToCartInput struct {
	TransactionID string
	ProductSku    string
	Qty           int64
}

type CheckoutInput struct {
	TransactionID string
	PaymentMethod string
}

func NewCheckoutUsecase(promoRepo PromoRepository, productRepo ProductRepository,
	cartRepo CartRepository) CheckoutUsecase {
	return CheckoutUsecase{
		promoRepo:   promoRepo,
		productRepo: productRepo,
		cartRepo:    cartRepo,
	}
}

// AddToCart add item to a cart with specified transaction id.
// Evaluate aplicable promotion with items in cart.
// Returns updated cart data.
func (u CheckoutUsecase) AddToCart(input AddToCartInput) (*entity.Cart, error) {
	if err := input.IsValid(); err != nil {
		return nil, err
	}

	cart, err := u.cartRepo.GetByTransactionID(input.TransactionID)
	if err != nil { //TODO: check if err != no result
		cart = entity.NewCart(input.TransactionID)
	}

	product, err := u.productRepo.GetBySku(input.ProductSku)
	if err != nil {
		return nil, err
	}

	//add item to cart
	cart.AddToCart(entity.CartItem{
		ProductSku: product.Sku,
		Name:       product.Name,
		Price:      product.Price,
		Qty:        input.Qty,
	})

	err = u.evaluateProductPromo(cart, product.Sku)
	if err != nil {
		log.Println("error evaluate product promo", err.Error())
	}

	u.applyPromoToCart(cart)

	//save cart
	cart.CalculateAmount()
	err = u.cartRepo.SaveCart(cart)
	if err != nil {
		return nil, err
	}

	return cart, nil
}

// evaluateProductPromo get available promo for given productSku.
// Will update cart aplicable promo.
func (u *CheckoutUsecase) evaluateProductPromo(cart *entity.Cart, productSku string) error {
	promo, err := u.promoRepo.GetPromoByProductSku(productSku)
	if err != nil {
		return err
	}

	//evaluate requirement
	promoMultiply := int64(math.MaxInt64)
	for _, req := range promo.ProductRequirements {
		cartItem, ok := cart.Items[req.ProductSku]
		if !ok {
			return nil
		}

		multiply := (cartItem.Qty - cartItem.PromoReqCount) / req.MinQty
		if multiply == 0 {
			return nil
		}
		if multiply < promoMultiply {
			promoMultiply = multiply
		}
	}

	//add applicable promo
	for _, o := range promo.PromoOutcome {
		applicablePromo := entity.CartApplicablePromo{
			PromoID:    promo.ID,
			ProductSku: o.ProductSku,
			Qty:        o.Qty,
			Name:       promo.Name,
		}
		if o.PromotionType == "percentage" {
			applicablePromo.Amount = o.Amount / 100 * cart.Items[o.ProductSku].Price
		}
		cart.AddToApplicablePromo(applicablePromo)
	}

	//increment counter
	for _, req := range promo.ProductRequirements {
		ci := cart.Items[req.ProductSku]
		ci.PromoReqCount += (req.MinQty * promoMultiply)
		cart.Items[req.ProductSku] = ci
	}

	return nil
}

// applyPromoToCart apply aplicable promo to items in the cart
func (u *CheckoutUsecase) applyPromoToCart(cart *entity.Cart) {
	for _, promo := range cart.ApplicablePromo {
		ci := cart.Items[promo.ProductSku]
		eligibleQty := ci.Qty - ci.PromoAppliedCount
		if eligibleQty == 0 {
			continue
		}
		if eligibleQty > promo.Qty {
			eligibleQty = promo.Qty
		}

		cart.AddToAppliedPromo(entity.CartAppliedPromo{
			PromoID:     promo.PromoID,
			ProductSku:  promo.ProductSku,
			Name:        promo.Name,
			Qty:         eligibleQty,
			TotalAmount: float64(eligibleQty) * promo.Amount,
		})

		//increment counter
		ci.PromoAppliedCount += eligibleQty
		cart.Items[promo.ProductSku] = ci
		promo.Qty -= eligibleQty
		cart.ApplicablePromo[promo.ProductSku] = promo
	}
}

// Checkout the cart with specified payment method.
// Will clear cart that has given transaction id.
func (u CheckoutUsecase) Checkout(input CheckoutInput) error {
	if err := input.IsValid(); err != nil {
		return err
	}

	err := u.cartRepo.ClearCart(input.TransactionID)
	if err != nil {
		return err
	}
	//TODO: store cart to db
	return nil
}

func (i AddToCartInput) IsValid() error {
	if i.TransactionID == "" {
		return fmt.Errorf("transaction id is required")
	}
	if i.ProductSku == "" {
		return fmt.Errorf("product sku is required")
	}
	if i.Qty <= 0 {
		return fmt.Errorf("qty must be > 0")
	}
	return nil
}

func (i CheckoutInput) IsValid() error {
	if i.TransactionID == "" {
		return fmt.Errorf("transaction id is required")
	}
	if i.PaymentMethod == "" {
		return fmt.Errorf("payment method is required")
	}
	return nil
}
