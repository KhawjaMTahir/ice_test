package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"interview/internal/controllers"
	mock_service "interview/internal/mocks"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

// setupControllerTest initializes the gomock controller, mock service, and router
func setupControllerTest(t *testing.T) (*gomock.Controller, *mock_service.MockCartServiceInterface, *gin.Engine) {
	ctrl := gomock.NewController(t)
	mockService := mock_service.NewMockCartServiceInterface(ctrl)

	router := gin.Default()

	return ctrl, mockService, router
}

func TestCartController_AddItemToCart(t *testing.T) {
	ctrl, mockService, router := setupControllerTest(t)
	defer ctrl.Finish()

	router.POST("/cart/add", controllers.NewCartController(mockService).AddItem)

	req, err := http.NewRequest(http.MethodPost, "/cart/add", nil)
	if err != nil {
		t.Fatalf("could not create request: %v", err)
	}

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, 302, w.Code)
}

func TestCartController_DeleteCartItem(t *testing.T) {
	ctrl, mockService, router := setupControllerTest(t)
	defer ctrl.Finish()

	router.DELETE("/cart/delete", controllers.NewCartController(mockService).DeleteCartItem)

	req, err := http.NewRequest(http.MethodDelete, "/cart/delete", nil)
	if err != nil {
		t.Fatalf("could not create request: %v", err)
	}

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, 302, w.Code)
}
