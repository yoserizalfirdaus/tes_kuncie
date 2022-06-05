package repository

import (
	"fmt"
	"sync"

	"github.com/yoserizalfirdaus/tes_kuncie/entity"
)

type CartRepository struct {
	storage map[string]*entity.Cart
	lock    *sync.RWMutex
}

func NewCartRepository() (CartRepository, error) {
	return CartRepository{
		storage: map[string]*entity.Cart{},
		lock:    &sync.RWMutex{},
	}, nil
}

func (c CartRepository) GetByTransactionID(trxID string) (*entity.Cart, error) {
	c.lock.RLock()
	defer c.lock.RUnlock()

	cart, ok := c.storage[trxID]
	if !ok {
		return nil, fmt.Errorf("cart not founc")
	}

	return cart, nil
}

func (c CartRepository) SaveCart(cart *entity.Cart) error {
	c.lock.Lock()
	defer c.lock.Unlock()

	c.storage[cart.TransactionID] = cart

	return nil
}

func (c CartRepository) ClearCart(trxID string) error {
	c.lock.Lock()
	defer c.lock.Unlock()

	delete(c.storage, trxID)
	return nil
}
