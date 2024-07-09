package controllers

import (
	"errors"
	service "interview/pkg/service"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type CartControllerInterface interface {
	ShowAddItemForm(c *gin.Context)
	AddItem(c *gin.Context)
	DeleteCartItem(c *gin.Context)
}

type CartController struct {
	cartService service.CartServiceInterface
}

func NewCartController(cartService service.CartServiceInterface) CartControllerInterface {
	return &CartController{cartService: cartService}
}

func (t *CartController) ShowAddItemForm(c *gin.Context) {
	_, err := c.Request.Cookie("ice_session_id")
	if errors.Is(err, http.ErrNoCookie) {
		c.SetCookie("ice_session_id", time.Now().String(), 3600, "/", "localhost", false, true)
	}

	t.cartService.GetCartData(c)
}

func (t *CartController) AddItem(c *gin.Context) {
	cookie, err := c.Request.Cookie("ice_session_id")
	if err != nil || errors.Is(err, http.ErrNoCookie) || (cookie != nil && cookie.Value == "") {
		c.Redirect(302, "/")
		return
	}

	t.cartService.AddItemToCart(c)
}

func (t *CartController) DeleteCartItem(c *gin.Context) {
	cookie, err := c.Request.Cookie("ice_session_id")
	if err != nil || errors.Is(err, http.ErrNoCookie) || (cookie != nil && cookie.Value == "") {
		c.Redirect(302, "/")
		return
	}

	t.cartService.DeleteCartItem(c)
}
