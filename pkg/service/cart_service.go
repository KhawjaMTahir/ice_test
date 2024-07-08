package service

import (
	"interview/pkg/entity"
	"interview/pkg/repository"
)

type CartService interface {
	AddItem(sessionID string, product string, quantity int) error
	DeleteItem(sessionID string, cartItemID int) error
	GetCart(sessionID string) ([]entity.CartItem, error)
}

type cartService struct {
	cartRepo repository.CartRepository
}

func NewCartService(cartRepo repository.CartRepository) CartService {
	return &cartService{
		cartRepo: cartRepo,
	}
}

func (s *cartService) AddItem(sessionID string, product string, quantity int) error {
	return s.cartRepo.AddItemToCart(sessionID, product, quantity)
}

func (s *cartService) DeleteItem(sessionID string, cartItemID int) error {
	return s.cartRepo.DeleteCartItem(sessionID, cartItemID)
}

func (s *cartService) GetCart(sessionID string) ([]entity.CartItem, error) {
	return s.cartRepo.GetCartData(sessionID)
}
