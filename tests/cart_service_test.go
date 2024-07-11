package tests

import (
	"bytes"
	"errors"
	"interview/internal/mocks"
	services "interview/internal/service"
	"interview/pkg/entity"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

// setupServiceTest initializes the gomock controller, mock repository, and service
func setupServiceTest(t *testing.T) (*gomock.Controller, *mocks.MockCartRepositoryInterface, services.Service) {
	ctrl := gomock.NewController(t)
	mockRepo := mocks.NewMockCartRepositoryInterface(ctrl)
	service := services.NewService(mockRepo)
	return ctrl, mockRepo, service
}

func TestCartService_AddItemToCart_Success(t *testing.T) {
	ctrl, mockRepo, service := setupServiceTest(t)
	defer ctrl.Finish()

	// Create a new gin context for testing
	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	form := "product=shoe&quantity=2"
	c.Request = httptest.NewRequest("POST", "/add-item", bytes.NewBufferString(form))
	c.Request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	c.Request.AddCookie(&http.Cookie{Name: "ice_session_id", Value: "mock_session_id"})

	// Mock expectations
	mockRepo.EXPECT().GetCartBySessionID("mock_session_id").Return(&entity.CartEntity{Model: gorm.Model{ID: 1}}, nil)
	mockRepo.EXPECT().GetCartItemByCartIDAndProductName(uint(1), "shoe").Return(nil, gorm.ErrRecordNotFound)
	mockRepo.EXPECT().CreateCartItem(gomock.Any()).Return(nil)

	service.AddItemToCart(c)

	assert.Equal(t, http.StatusFound, c.Writer.Status())
}

func TestCartService_AddItemToCart_Error(t *testing.T) {
	ctrl, mockRepo, service := setupServiceTest(t)
	defer ctrl.Finish()

	// Create a new gin context for testing
	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	form := "product=shoe&quantity=2"
	c.Request = httptest.NewRequest("POST", "/add-item", bytes.NewBufferString(form))
	c.Request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	c.Request.AddCookie(&http.Cookie{Name: "ice_session_id", Value: "mock_session_id"})

	mockRepo.EXPECT().GetCartBySessionID("mock_session_id").Return(nil, errors.New("session not found"))

	service.AddItemToCart(c)

	assert.Equal(t, http.StatusFound, c.Writer.Status())
}

func TestCartService_DeleteCartItem_Success(t *testing.T) {
	ctrl, mockRepo, service := setupServiceTest(t)
	defer ctrl.Finish()

	// Create a new gin context for testing
	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	c.Request = httptest.NewRequest("GET", "/delete-item?cart_item_id=1", nil)
	c.Request.AddCookie(&http.Cookie{Name: "ice_session_id", Value: "mock_session_id"})

	mockRepo.EXPECT().GetCartBySessionID("mock_session_id").Return(&entity.CartEntity{Model: gorm.Model{ID: 1}}, nil)
	mockRepo.EXPECT().DeleteCartItem(1, uint(1)).Return(nil)

	service.DeleteCartItem(c)

	assert.Equal(t, http.StatusFound, c.Writer.Status())
}

func TestCartService_DeleteCartItem_Error(t *testing.T) {
	ctrl, mockRepo, service := setupServiceTest(t)
	defer ctrl.Finish()

	// Create a new gin context for testing
	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	c.Request = httptest.NewRequest("GET", "/delete-item?cart_item_id=1", nil)
	c.Request.AddCookie(&http.Cookie{Name: "ice_session_id", Value: "mock_session_id"})

	mockRepo.EXPECT().GetCartBySessionID("mock_session_id").Return(&entity.CartEntity{Model: gorm.Model{ID: 1}}, nil)
	mockRepo.EXPECT().DeleteCartItem(1, uint(1)).Return(errors.New("delete error"))

	service.DeleteCartItem(c)

	assert.Equal(t, http.StatusFound, c.Writer.Status())
}
