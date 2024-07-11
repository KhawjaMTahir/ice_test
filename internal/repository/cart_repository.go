package repository

import (
	"interview/pkg/entity"
)

type (
	CartRepositoryInterface interface {
		GetCartBySessionID(sessionID string) (*entity.CartEntity, error)
		CreateCart(cart *entity.CartEntity) error
		GetCartItemByCartIDAndProductName(cartID uint, productName string) (*entity.CartItem, error)
		CreateCartItem(cartItem *entity.CartItem) error
		UpdateCartItem(cartItem *entity.CartItem) error
		DeleteCartItem(cartItemID int, cartID uint) error
		GetCartItemsByCartID(cartID uint) ([]entity.CartItem, error)
	}
)

func (r *repository) GetCartBySessionID(sessionID string) (*entity.CartEntity, error) {
	var cartEntity entity.CartEntity
	result := r.db.Where("status = ? AND session_id = ?", entity.CartOpen, sessionID).First(&cartEntity)
	if result.Error != nil {
		return nil, result.Error
	}
	return &cartEntity, nil
}

func (r *repository) CreateCart(cart *entity.CartEntity) error {
	return r.db.Create(cart).Error
}

func (r *repository) GetCartItemByCartIDAndProductName(cartID uint, productName string) (*entity.CartItem, error) {
	var cartItem entity.CartItem
	result := r.db.Where("cart_id = ? AND product_name = ?", cartID, productName).First(&cartItem)
	if result.Error != nil {
		return nil, result.Error
	}
	return &cartItem, nil
}

func (r *repository) CreateCartItem(cartItem *entity.CartItem) error {
	return r.db.Create(cartItem).Error
}

func (r *repository) UpdateCartItem(cartItem *entity.CartItem) error {
	return r.db.Save(cartItem).Error
}

func (r *repository) DeleteCartItem(cartItemID int, cartID uint) error {
	return r.db.Where("id = ? AND cart_id = ?", cartItemID, cartID).Delete(&entity.CartItem{}).Error
}

func (r *repository) GetCartItemsByCartID(cartID uint) ([]entity.CartItem, error) {
	var cartItems []entity.CartItem
	result := r.db.Where("cart_id = ?", cartID).Find(&cartItems)
	if result.Error != nil {
		return nil, result.Error
	}
	return cartItems, nil
}
