package tests

import (
	"errors"
	"interview/internal/mocks"
	"interview/pkg/entity"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

// setupTest initializes the gomock controller and mock repository
func setupTest(t *testing.T) (*gomock.Controller, *mocks.MockCartRepositoryInterface) {
	ctrl := gomock.NewController(t)
	mockRepo := mocks.NewMockCartRepositoryInterface(ctrl)
	t.Cleanup(ctrl.Finish)
	return ctrl, mockRepo
}

func TestCartRepository_GetCartBySessionID_Success(t *testing.T) {
	_, mockRepo := setupTest(t)
	expectedCart := &entity.CartEntity{}
	mockRepo.EXPECT().GetCartBySessionID("mock_session_id").Return(expectedCart, nil)

	cart, err := mockRepo.GetCartBySessionID("mock_session_id")

	assert.NoError(t, err)
	assert.Equal(t, expectedCart, cart)
}

func TestCartRepository_GetCartBySessionID_Error(t *testing.T) {
	_, mockRepo := setupTest(t)
	mockRepo.EXPECT().GetCartBySessionID("invalid_session_id").Return(nil, errors.New("session not found"))

	cart, err := mockRepo.GetCartBySessionID("invalid_session_id")

	assert.Error(t, err)
	assert.Nil(t, cart)
}

func TestCartRepository_CreateCart_Success(t *testing.T) {
	_, mockRepo := setupTest(t)
	cart := &entity.CartEntity{}
	mockRepo.EXPECT().CreateCart(cart).Return(nil)

	err := mockRepo.CreateCart(cart)

	assert.NoError(t, err)
}

func TestCartRepository_CreateCart_Error(t *testing.T) {
	_, mockRepo := setupTest(t)
	cart := &entity.CartEntity{}
	mockRepo.EXPECT().CreateCart(cart).Return(errors.New("creation error"))

	err := mockRepo.CreateCart(cart)

	assert.Error(t, err)
}

func TestCartRepository_GetCartItemByCartIDAndProductName_Success(t *testing.T) {
	_, mockRepo := setupTest(t)
	expectedItem := &entity.CartItem{}
	mockRepo.EXPECT().GetCartItemByCartIDAndProductName(uint(1), "shoe").Return(expectedItem, nil)

	item, err := mockRepo.GetCartItemByCartIDAndProductName(uint(1), "shoe")

	assert.NoError(t, err)
	assert.Equal(t, expectedItem, item)
}

func TestCartRepository_GetCartItemByCartIDAndProductName_Error(t *testing.T) {
	_, mockRepo := setupTest(t)
	mockRepo.EXPECT().GetCartItemByCartIDAndProductName(uint(1), "invalid_product").Return(nil, errors.New("not found"))

	item, err := mockRepo.GetCartItemByCartIDAndProductName(uint(1), "invalid_product")

	assert.Error(t, err)
	assert.Nil(t, item)
}

func TestCartRepository_CreateCartItem_Success(t *testing.T) {
	_, mockRepo := setupTest(t)
	item := &entity.CartItem{ProductName: "shoe", Quantity: 2}
	mockRepo.EXPECT().CreateCartItem(item).Return(nil)

	err := mockRepo.CreateCartItem(item)

	assert.NoError(t, err)
}

func TestCartRepository_CreateCartItem_Error(t *testing.T) {
	_, mockRepo := setupTest(t)
	item := &entity.CartItem{ProductName: "", Quantity: 2} // Invalid item name
	mockRepo.EXPECT().CreateCartItem(item).Return(errors.New("invalid item"))

	err := mockRepo.CreateCartItem(item)

	assert.Error(t, err)
}

func TestCartRepository_UpdateCartItem_Success(t *testing.T) {
	_, mockRepo := setupTest(t)
	item := &entity.CartItem{ProductName: "shoe", Quantity: 2}
	mockRepo.EXPECT().UpdateCartItem(item).Return(nil)

	err := mockRepo.UpdateCartItem(item)

	assert.NoError(t, err)
}

func TestCartRepository_UpdateCartItem_Error(t *testing.T) {
	_, mockRepo := setupTest(t)
	item := &entity.CartItem{ProductName: "shoe", Quantity: 2}
	mockRepo.EXPECT().UpdateCartItem(item).Return(errors.New("update error"))

	err := mockRepo.UpdateCartItem(item)

	assert.Error(t, err)
}

func TestCartRepository_DeleteCartItem_Success(t *testing.T) {
	_, mockRepo := setupTest(t)
	mockRepo.EXPECT().DeleteCartItem(1, uint(1)).Return(nil)

	err := mockRepo.DeleteCartItem(1, uint(1))

	assert.NoError(t, err)
}

func TestCartRepository_DeleteCartItem_Error(t *testing.T) {
	_, mockRepo := setupTest(t)
	mockRepo.EXPECT().DeleteCartItem(1, uint(1)).Return(errors.New("delete error"))

	err := mockRepo.DeleteCartItem(1, uint(1))

	assert.Error(t, err)
}

func TestCartRepository_GetCartItemsByCartID_Success(t *testing.T) {
	_, mockRepo := setupTest(t)
	expectedItems := []entity.CartItem{}
	mockRepo.EXPECT().GetCartItemsByCartID(uint(1)).Return(expectedItems, nil)

	items, err := mockRepo.GetCartItemsByCartID(uint(1))

	assert.NoError(t, err)
	assert.Equal(t, expectedItems, items)
}

func TestCartRepository_GetCartItemsByCartID_Error(t *testing.T) {
	_, mockRepo := setupTest(t)
	mockRepo.EXPECT().GetCartItemsByCartID(uint(1)).Return(nil, errors.New("not found"))

	items, err := mockRepo.GetCartItemsByCartID(uint(1))

	assert.Error(t, err)
	assert.Nil(t, items)
}
